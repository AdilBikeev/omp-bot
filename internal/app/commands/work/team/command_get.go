package team

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkTeamCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	teamId, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf(`Invalid argument "%v"`, args))
		c.bot.Send(msg)

		return
	}

	team, err := c.service.Describe(teamId)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)

		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, team.String())
	c.bot.Send(msg)
}