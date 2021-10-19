package team

import (
	"github.com/ozonmp/omp-bot/internal/model/work"
)

type TeamService interface {
	Describe(teamID uint64) (*work.Team, error)
	List(cursor uint64, limit uint64) ([]work.Team, error)
	Create(work.Team) (uint64, error)
	Update(teamID uint64, team work.Team) error
	Remove(teamID uint64) (bool, error)
	Count() uint64
}

func NewWorkTeamService() *WorkTeamService {
	return &WorkTeamService{}
}
