package work

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/work/team"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type workCommander struct {
	bot           *tgbotapi.BotAPI
	teamCommander Commander
}

func NewWorkCommander(
	bot *tgbotapi.BotAPI,
) *workCommander {
	return &workCommander{
		bot:           bot,
		teamCommander: team.NewWorkTeamCommander(bot),
	}
}

func (c *workCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "team":
		c.teamCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("workCommander.HandleCallback: unknown team - %s", callbackPath.Subdomain)
	}
}

func (c *workCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "team":
		c.teamCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("workCommander.HandleCommand: unknown team - %s", commandPath.Subdomain)
	}
}
