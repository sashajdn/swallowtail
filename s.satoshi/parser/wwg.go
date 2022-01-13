package parser

import (
	discordproto "swallowtail/s.discord/proto"
)

func init() {
	register(discordproto.DiscordMoonModMessagesChannel, []TradeParser{
		&DCAParser{},
		&DMAParser{},
	})
}
