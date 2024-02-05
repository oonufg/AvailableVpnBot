package main

import (
	"AvailableVpn/internal/bot"
	"AvailableVpn/internal/domain"
	"context"
)

func main() {
	ctx := context.Background()
	repo := domain.CreateOvpnRepository()
	b := bot.CreateBot("6865414091:AAFXg2YL-f_W5zJFLKbvzVTRT6gzIYstCz4", repo)
	b.Start(ctx)
}
