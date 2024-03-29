package parser

import (
	"context"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	tradeengineproto "swallowtail/s.trade-engine/proto"
)

func TestDMAParser(t *testing.T) {
	tests := []struct {
		name          string
		content       string
		username      string
		currentValue  float64
		expectedTrade *tradeengineproto.TradeStrategy
		withErr       bool
	}{
		{
			name: "rego_full_trade_wwg_with_three_tp",
			content: `Hey guys I'm LONG BTC here.

			ENTRY: 50000
			STOP: 49000

			TP1: 52000
			TP2: 54000
			TP3: 58000

			This should give us an 4.5RR and 15.7% increase
			`,
			username:     "rego",
			currentValue: 50000,
			expectedTrade: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "REGO",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DMA_MARKET,
				Asset:              "BTC",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				Entries:            []float32{50000},
				StopLoss:           49000,
				CurrentPrice:       50000,
				TakeProfits: []float32{
					52000, 54000, 58000,
				},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
				},
			},
		},
		{
			name:         "bluntz_example_second_entry",
			username:     "bluntz",
			currentValue: 170,
			content: `Going to enter that second sol entry here as i think it just got frontrun by 0.3%

			entry 2: now 165
			stop 135.61
			target 259.7

			57% 3.25RR
			`,
			expectedTrade: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "BLUNTZ",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DMA_MARKET,
				Asset:              "SOL",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				CurrentPrice:       170,
				Entries:            []float32{165},
				StopLoss:           135.61,
				TakeProfits: []float32{
					259.7,
				},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
					tradeengineproto.VENUE_FTX,
				},
			},
		},
		{
			name:         "astekz_example_1_aave_no_take_profit",
			username:     "astekz",
			currentValue: 344,
			content: `
			aave 
			spot or low lev long 343
			stop 323
			@​everyone
			[Attachments]
			https://cdn.discordapp.com/attachments/869596440777883749/885529381479518219/unknown.png
			`,
			expectedTrade: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "ASTEKZ",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DMA_MARKET,
				Asset:              "AAVE",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_SPOT,
				CurrentPrice:       344,
				Entries:            []float32{343},
				StopLoss:           323,
				TakeProfits:        []float32{},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
				},
			},
		},
		{
			name:         "eli_example_1_limit_srm",
			username:     "eli",
			currentValue: 10.9,
			content:      `SRM LIMIT LONG 9.80 stop 8.90 tp 13 18 @​everyone`,
			expectedTrade: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "ELI",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DMA_LIMIT,
				Asset:              "SRM",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				CurrentPrice:       10.9,
				Entries:            []float32{9.80},
				StopLoss:           8.90,
				TakeProfits:        []float32{13, 18},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
				},
			},
		},
		{
			name:         "cryptogodjohnny_example_1_market_buy_srm",
			username:     "cryptogodjohnny",
			currentValue: 0.041,
			content: `
			RSR $0.0402

			SL $0.0374
			`,
			expectedTrade: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "CRYPTOGODJOHNNY",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DMA_MARKET,
				Asset:              "RSR",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				CurrentPrice:       0.041,
				Entries:            []float32{0.0402},
				StopLoss:           0.0374,
				TakeProfits:        []float32{},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
				},
			},
		},
		{
			name:         "cryptogodjohnny_example_2_market_btc_short",
			username:     "cryptogodjohnny",
			currentValue: 46500,
			content: `
			Btc short $46650

			SL 47801 

			Tp 45800 44540 43680 42112
			@​Scalps High risk
			`,
			expectedTrade: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "CRYPTOGODJOHNNY",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DMA_MARKET,
				Asset:              "BTC",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_SHORT,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				CurrentPrice:       46500,
				Entries:            []float32{46650},
				StopLoss:           47801,
				TakeProfits: []float32{
					45800,
					44540,
					43680,
					42112,
				},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
				},
			},
		},
		{
			name:          "ticker_but_no_valid_information_example_ftt",
			content:       `if i ever get ftt at 50 again im gonna put entire portfolio there like jeliaz said`,
			expectedTrade: nil,
			withErr:       true,
		},
		{
			name:         "cryptogodjohnny_example_3_market_xtz_long",
			username:     "cryptogodjohnny",
			currentValue: 6.28,
			content: `XTZ 
			6.28 6.02 
			`,
			expectedTrade: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "CRYPTOGODJOHNNY",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DMA_MARKET,
				Asset:              "XTZ",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				CurrentPrice:       6.28,
				Entries:            []float32{6.28},
				StopLoss:           6.02,
				TakeProfits:        []float32{},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
				},
			},
		},
		{
			name:         "rego_avax_trade_long",
			username:     "rego",
			currentValue: 46.0,
			content: `
			rego: AVAXUSDT - SCALP LONG

			Entry:  46.64
			Stop: 44.00 (5.00%) 
			TP: 50, 52 , 55 
			`,
			expectedTrade: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "REGO",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DMA_MARKET,
				Asset:              "AVAX",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				CurrentPrice:       46.0,
				Entries:            []float32{46.64},
				StopLoss:           44.0,
				TakeProfits:        []float32{50, 52, 55},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
				},
			},
		},
		{
			name:         "multi_venue",
			username:     "rego",
			currentValue: 46.0,
			content: `
			rego: STEP - SCALP LONG

			Entry:  46.64
			Stop: 44.00 (5.00%) 
			TP: 50, 52 , 55 
			`,
			expectedTrade: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "REGO",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DMA_MARKET,
				Asset:              "STEP",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				CurrentPrice:       46.0,
				Entries:            []float32{46.64},
				StopLoss:           44.0,
				TakeProfits:        []float32{50, 52, 55},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
					tradeengineproto.VENUE_FTX,
				},
			},
		},
		{
			name:         "single_venue_entry_above_current_price_long",
			username:     "rego",
			currentValue: 91.41,
			content: `
			SOL 99 77
			`,
			expectedTrade: &tradeengineproto.TradeStrategy{
				HumanizedActorName: "REGO",
				ActorType:          tradeengineproto.ACTOR_TYPE_EXTERNAL,
				ExecutionStrategy:  tradeengineproto.EXECUTION_STRATEGY_DMA_MARKET,
				Asset:              "SOL",
				Pair:               tradeengineproto.TRADE_PAIR_USDT,
				TradeSide:          tradeengineproto.TRADE_SIDE_LONG,
				InstrumentType:     tradeengineproto.INSTRUMENT_TYPE_FUTURE_PERPETUAL,
				CurrentPrice:       91.41,
				Entries:            []float32{99},
				StopLoss:           77,
				TakeProfits:        []float32{},
				TradeableVenues: []tradeengineproto.VENUE{
					tradeengineproto.VENUE_BINANCE,
					tradeengineproto.VENUE_FTX,
				},
			},
		},
	}

	originalBinanceAssetPairs := binanceInstruments
	binanceInstruments = map[string]bool{
		"btc":  true,
		"sol":  true,
		"aave": true,
		"srm":  true,
		"rsr":  true,
		"xtz":  true,
		"ftt":  true,
		"avax": true,
		"step": true,
	}

	originalFTXAssetPairs := ftxInstruments
	ftxInstruments = map[string]bool{
		"sol-perp":  true,
		"step-perp": true,
	}

	originalFetcher := fetchLatestPrice
	t.Cleanup(func() {
		binanceInstruments = originalBinanceAssetPairs
		ftxInstruments = originalFTXAssetPairs
		fetchLatestPrice = originalFetcher
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fetchLatestPrice = func(_ context.Context, _ string) (float64, error) {
				return tt.currentValue, nil
			}

			trade, err := (&DMAParser{}).Parse(context.Background(), tt.content, &discordgo.MessageCreate{
				Message: &discordgo.Message{
					Author: &discordgo.User{
						Username: tt.username,
					},
				},
			}, tradeengineproto.ACTOR_TYPE_EXTERNAL)

			switch {
			case !tt.withErr:
				require.NoError(t, err)
				assert.Equal(t, tt.expectedTrade, trade)
			default:
				require.Error(t, err)
			}
		})
	}
}
