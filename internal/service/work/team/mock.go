package team

import (
	"github.com/ozonmp/omp-bot/internal/model/work"
)

var teamMocks = []work.Team{
	{
		ID:          1,
		Name:        "NaVi",
		Description: "Natus Vincere - team Dota 2",
		Size:        5,
		CreatedAt:   "15-10-2010",
	},
	{
		ID:          2,
		Name:        "Team Spirit",
		Description: "Russian team Dota 2",
		Size:        5,
		CreatedAt:   "15-09-2005",
	},

	{
		ID:          3,
		Name:        "Лена Кукв",
		Description: "Команда КВН",
		Size:        3,
		CreatedAt:   "15-04-2007",
	},
	{
		ID:          4,
		Name:        "PSG",
		Description: "PSG (Paris Saint-Germain) - футбольная команда",
		Size:        11,
		CreatedAt:   "15-09-1970",
	},
}
