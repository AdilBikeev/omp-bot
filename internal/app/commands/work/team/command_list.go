package team

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkTeamCommander) List(inputMessage *tgbotapi.Message) {
	outputMessage := "All the teams: \n\n"
	var msg tgbotapi.MessageConfig

	teams, err := c.service.List(0, c.service.Count())
	if err == nil {
		for i, t := range teams {
			outputMessage += fmt.Sprintf("#%d", i+1)
			outputMessage += t.String()
			outputMessage += "\n\n"
		}

		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)
	} else {
		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
	}

	c.bot.Send(msg)
}
