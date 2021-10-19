package work

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/work/team"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type WorkCommander struct {
	bot           *tgbotapi.BotAPI
	teamCommander Commander
}

func NewWorkCommander(
	bot *tgbotapi.BotAPI,
) *WorkCommander {
	return &WorkCommander{
		bot: bot,
		// teamCommander
		teamCommander: team.NewWorkTeamCommander(bot),
	}
}
