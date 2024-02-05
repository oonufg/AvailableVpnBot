package main

import (
	"AvailableVpn/internal/bot"
	"AvailableVpn/internal/domain"
	"context"
)

func main() {
	ctx := context.Background()
	domain.DownloadAllOvpnFiles()
	repo := domain.CreateOvpnRepository()
	b := bot.CreateBot("Token", repo)
	b.Start(ctx)
}
