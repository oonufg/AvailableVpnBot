package bot

import (
	"AvailableVpn/internal/domain"
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type TgBot struct {
	ovpnRepository *domain.OvpnRepo
	apiKey         string
	bot            *bot.Bot
	TgAPI
	EndPoints
}

type EndPoints interface {
	GetVpnsList(ctx context.Context, upd *models.Update) (*bot.SendMessageParams, error)
	GetAvailableCountries(ctx context.Context, upd *models.Update) *bot.SendMessageParams
}

type TgAPI interface {
	Start()
	HandleUpdate()
	SendFiles()
}

func CreateBot(apiKey string, repo *domain.OvpnRepo) *TgBot {

	return &TgBot{
		apiKey:         apiKey,
		ovpnRepository: repo,
	}
}

func (tgBot *TgBot) Start(ctx context.Context) {
	bot, _ := bot.New(tgBot.apiKey, bot.WithDefaultHandler(tgBot.HandleUpdate))
	tgBot.bot = bot
	tgBot.bot.Start(ctx)
}

func (tgBot *TgBot) HandleUpdate(ctx context.Context, b *bot.Bot, update *models.Update) {
	command := update.Message.Text
	parsedString := strings.Split(command, " ")

	switch parsedString[0] {
	case "/VpnList":
		if len(parsedString) != 3 {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Комманда введена не верно \nПример /VpnList Russia tcp",
			})
		}

		vList, err := tgBot.GetVpnsList(ctx, update)

		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Комманда введена не верно \nПример /VpnList Russia tcp",
			})
		}
		b.SendMediaGroup(ctx, &bot.SendMediaGroupParams{
			ChatID: update.Message.Chat.ID,
			Media:  vList,
		})

	case "/AvailableCountry":
		b.SendMessage(ctx, tgBot.GetAvailableCountries(ctx, update))

	default:
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Команда не распознана",
		})
	}
}
