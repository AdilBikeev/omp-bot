package team

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkTeamCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		`/help__work__team - print list of commands
		 /get__work__team - get a team by id
		 /list__work__team - get a list of teams
		 /delete__work__team - delete an existing team by id
		 /new__work__team - create a new team.
		 /edit__work__team - edit a team by id`,
	)

	c.bot.Send(msg)
}
