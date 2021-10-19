package team

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	workTeamService "github.com/ozonmp/omp-bot/internal/service/work/team"
)

type TeamCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type WorkTeamCommander struct {
	bot     *tgbotapi.BotAPI
	service workTeamService.TeamService
}

func NewWorkTeamCommander(bot *tgbotapi.BotAPI) *WorkTeamCommander {
	service := workTeamService.NewWorkTeamService()

	return &WorkTeamCommander{
		bot:     bot,
		service: service,
	}
}

func (c *WorkTeamCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("WorkTeamCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *WorkTeamCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		log.Printf("WorkTeamCommander.HandleCommand: unknown CommandName: %s", commandPath.CommandName)
	}
}
