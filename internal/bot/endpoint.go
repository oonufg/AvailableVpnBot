package bot

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// country protocol
func (tgb *TgBot) GetVpnsList(ctx context.Context, upd *models.Update) *bot.SendMessageParams {

	return &bot.SendMessageParams{
		ChatID: upd.Message.Chat.ID,
		Text:   "List",
	}
}

func (tgb *TgBot) GetVpnFile(ctx context.Context, upd *models.Update) *bot.SendMessageParams {
	return &bot.SendMessageParams{
		ChatID: upd.Message.Chat.ID,
		Text:   "File",
	}
}
