package team

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	workTeamService "github.com/ozonmp/omp-bot/internal/service/work/team"
	//"github.com/ozonmp/omp-bot/internal/app/commands/work/team"
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

func (c *WorkTeamCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "team":
		c.HandleCallback(callback, callbackPath)
	default:
		log.Printf("WorkCommander.HandleCallback: unknown team - %s", callbackPath.Subdomain)
	}
}

func (c *WorkTeamCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "team":
		c.HandleCommand(msg, commandPath)
	default:
		log.Printf("WorkCommander.HandleCommand: unknown team - %s", commandPath.Subdomain)
	}
}

func NewWorkTeamCommander(bot *tgbotapi.BotAPI) *WorkTeamCommander {
	service := workTeamService.NewWorkTeamService()

	return &WorkTeamCommander{
		bot:     bot,
		service: service,
	}
}
