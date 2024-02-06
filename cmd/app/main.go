package main

import (
	cfg "AvailableVpn/config"
	"AvailableVpn/internal/bot"
	"AvailableVpn/internal/domain"
	"context"
)

func main() {
	ctx := context.Background()
	cfg := cfg.LoadConfig()
	domain.DownloadAllOvpnFiles()
	repo := domain.CreateOvpnRepository()

	b := bot.CreateBot(cfg.TG_API_KEY, repo)
	b.Start(ctx)
}
