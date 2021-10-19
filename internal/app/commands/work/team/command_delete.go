package team

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkTeamCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	var msg tgbotapi.MessageConfig
	teamId, err := strconv.ParseUint(args, 10, 64)

	if err != nil {
		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf(`Invalid argument "%v"`, args))
	} else {
		_, err = c.service.Remove(teamId)
		if err != nil {
			msg = tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		} else {
			msg = tgbotapi.NewMessage(inputMessage.Chat.ID, "Team was deleted successfully.")
		}
	}

	c.bot.Send(msg)
}
