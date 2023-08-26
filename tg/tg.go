package tg

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

func NewTgBot(token string, chatID int64) *TgBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	return &TgBot{
		bot:    bot,
		chatID: chatID,
	}
}

func (t *TgBot) PostShitAboutTown(text string) error {
	msg := tgbotapi.NewMessage(t.chatID, text)
	_, err := t.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
