package bot

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// /VpnList country protocol
func (tgb *TgBot) GetVpnsList(ctx context.Context, upd *models.Update) ([]models.InputMedia, error) {
	parsedQuery := strings.Split(upd.Message.Text, " ")
	ovpns := tgb.ovpnRepository.GetOvpnsByParam(parsedQuery[1], parsedQuery[2])

	ovpnsList := make([]models.InputMedia, 0)
	for _, val := range ovpns {

		file, err := os.Open(fmt.Sprintf("./ovpn/%s", val.GetFilename()))
		if err != nil {
			log.Println("Не удалось получить файл")
			return nil, errors.New("Не удалось получить файл")
		}
		msg := &models.InputMediaDocument{
			Media:           fmt.Sprintf("attach://%s", val.GetFilename()),
			Caption:         val.GetFilename(),
			MediaAttachment: file,
		}
		ovpnsList = append(ovpnsList, msg)
	}

	return ovpnsList, nil
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
