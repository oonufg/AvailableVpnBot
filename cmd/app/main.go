package main

import (
	"AvailableVpn/internal/bot"
	"AvailableVpn/internal/domain"
	"context"
)

func main() {
	ctx := context.Background()
	rep := domain.CreateOvpnRepository()
	bot := bot.CreateBot("", rep)
	bot.Start(ctx)
}
