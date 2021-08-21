package commands

import (
	"context"
	"fmt"
	"swallowtail/libraries/gerrors"

	"github.com/bwmarrin/discordgo"
)

const (
	pingCommandID = "ping"
)

func init() {
	register(pingCommandID, &Command{
		ID:      pingCommandID,
		Usage:   `!ping`,
		Handler: pingCommand,
	})
}

func pingCommand(ctx context.Context, tokens []string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	if _, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":wave: Hello <@%s>, what can I do for you?", m.Author.ID)); err != nil {
		return gerrors.Augment(err, "failed_to_ping_discord", nil)
	}

	return nil
}
