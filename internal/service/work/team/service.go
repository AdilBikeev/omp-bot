package team

import (
	"fmt"
	"log"

	"github.com/ozonmp/omp-bot/internal/model/work"
	// 	"errors"
	// "fmt"
)

type WorkTeamService struct {
	entities map[uint64]work.Team
}

func (w *WorkTeamService) Describe(teamID uint64) (*work.Team, error) {
	log.Printf("Work.TeamService: getting team with id %v", teamID)

	entity, ok := w.entities[teamID]
	if !ok {
		return nil, fmt.Errorf(entityNotFound, teamID)
	}

	return &entity, nil
}

func (w *WorkTeamService) List(cursor uint64, limit uint64) ([]work.Team, error) {
	log.Printf("Work.TeamService: listing teams in range from %v to %v", cursor, cursor+limit)

	count := uint64(len(w.entities))
	if cursor > count {
		return []work.Team{}, fmt.Errorf("IndexOutOfRange, cursor=%d", cursor)
	} else if count == 0 {
		return []work.Team{}, nil
	}

	maxIndex := cursor + limit
	if maxIndex > count {
		maxIndex = count
	}

	lst := make([]work.Team, 0, limit)
	for _, val := range w.entities {
		lst = append(lst, val)
	}

	return lst, nil
}

func (w *WorkTeamService) Create(newTeam work.Team) (uint64, error) {
	log.Printf("Work.TeamService: creating team %#v", newTeam)

	w.entities[newTeam.ID] = newTeam

	return uint64(len(w.entities)), nil
}

func (w *WorkTeamService) Update(teamID uint64, team work.Team) error {
	log.Printf("Work.TeamService: updating team %v with value: %#v", teamID, team)

	entity, ok := w.entities[teamID]

	if !ok {
		return fmt.Errorf(entityNotFound, teamID)
	}

	entity.Name = team.Name
	entity.Description = team.Description
	entity.Size = team.Size
	entity.CreatedAt = team.CreatedAt

	return nil
}

func (w *WorkTeamService) Remove(teamID uint64) (bool, error) {
	log.Printf("Work.TeamService: deleting team with id %v", teamID)

	if _, ok := w.entities[teamID]; !ok {
		return false, fmt.Errorf(entityNotFound, teamID)
	}

	delete(w.entities, teamID)

	return true, nil
}

func (w *WorkTeamService) Count() uint64 {
	return uint64(len(w.entities))
}
