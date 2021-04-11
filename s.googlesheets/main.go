package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	coingecko "swallowtail/s.coingecko/clients"
	discord "swallowtail/s.discord/clients"
	"swallowtail/s.googlesheets/clients"
	"swallowtail/s.googlesheets/owner"
	"swallowtail/s.googlesheets/spreadsheet"
	"swallowtail/s.googlesheets/sync"
	"syscall"
	"time"

	"github.com/monzo/slog"
)

var (
	// Move
	defaultAlexGoogleSpreadsheetID = "1AYtRsdEcoEjmh-OtribxJ9et7qvCf6Z_UkkYNnKqqZY"
	defaultBenGoogleSpreadsheetID  = "1Krg7O8h-ItK42dTC-ey9HOh6v1w8T2SCHcUsVJdUnmI"

	defaultSyncInterval = time.Duration(1 * time.Minute)
)

// TODO: move
type exchangeClient struct {
	exchangeID string
	c          *coingecko.CoinGeckoClient
}

func (ex exchangeClient) GetPrice(ctx context.Context, symbol, assetPair string) (float64, error) {
	switch ex.exchangeID {
	case coingecko.CoingeckoClientID:
		return ex.c.GetCurrentPriceFromSymbol(ctx, symbol, assetPair)
	default:
		return ex.c.GetCurrentPriceFromSymbol(ctx, symbol, assetPair)
	}
}

func (ex exchangeClient) Ping(ctx context.Context) bool {
	return ex.c.Ping(ctx)
}

func main() {
	ctx := context.Background()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	defer slog.Warn(ctx, "Received shutdown signal....")

	c, err := clients.New(ctx)
	if err != nil {
		log.Fatalf("Failed to init googlesheets spreadsheet: %v", err)
	}
	// MessageClient TODO: fetch from elsewhere as to not have two discord clients
	mc := discord.New()
	// ExchangeClient
	ex := exchangeClient{
		c: coingecko.New(ctx),
	}
	done := make(chan struct{}, 1)

	// Parse yaml? will be web-based one day.
	owners := []*owner.GooglesheetOwner{
		owner.New(
			defaultAlexGoogleSpreadsheetID, "Alex", "805513165428883487",
			[]string{"TestSpots", "ParentsSpots"},
			mc, true,
		),
		// owner.New(defaultBenGoogleSpreadsheetID, "Ben", "814142503393558558", sheetIDs, mc, true),
	}

	for _, owner := range owners {
		ss := spreadsheet.New(owner.SpreadsheetID, owner.SheetsID, c, owner)
		syncer := sync.NewGoogleSheetsPorfolioSyncer(
			ss, ex, defaultSyncInterval, done, true,
		)
		go syncer.Start(ctx)
	}

	slog.Info(ctx, "Starting googlesheets syncer...")
	if err != nil {
		panic(fmt.Sprintf("%s", err.Error()))
	}
	select {
	case <-sc:
	}
}
