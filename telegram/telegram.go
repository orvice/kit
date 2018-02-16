package telegram

import (
	"fmt"

	"github.com/orvice/kit/mod"
	"gopkg.in/telegram-bot-api.v4"
)

var (
	bots        map[string]*tgbotapi.BotAPI
	BotNotFound = fmt.Errorf("telegram bot not found")
)

func Init(cfgs map[string]mod.Telegram) error {
	if bots == nil {
		bots = make(map[string]*tgbotapi.BotAPI)
	}
	for k, v := range cfgs {
		bot, err := tgbotapi.NewBotAPI(v.Token)
		if err != nil {
			return err
		}
		bots[k] = bot
	}
	return nil
}

func GetBot(k string) (*tgbotapi.BotAPI, error) {
	b, ok := bots[k]
	if !ok {
		return nil, BotNotFound
	}
	return b, nil
}
