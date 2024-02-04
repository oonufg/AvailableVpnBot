package bot

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// /VpnList country protocol
func (tgb *TgBot) GetVpnsList(ctx context.Context, upd *models.Update) *bot.SendMessageParams {

	return &bot.SendMessageParams{
		ChatID: upd.Message.Chat.ID,
		Text:   "List",
	}
}

func (tgb *TgBot) GetAvailableCountries(ctx context.Context, upd *models.Update) *bot.SendMessageParams {
	countries := tgb.ovpnRepository.GetAvailableCountries()
	stringBuilder := strings.Builder{}
	for _, val := range countries {
		stringBuilder.WriteString(val + "\n")
	}
	return &bot.SendMessageParams{
		ChatID: upd.Message.Chat.ID,
		Text:   stringBuilder.String(),
	}
}
