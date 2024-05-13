package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

type Score struct {
	Player string `json:"player"`
	Score  []int  `json:"score"`
}

type ScorecardManager struct {
	scorecards map[GameID]*Scorecard
}

type GameID string

func (s *ScorecardManager) Create(player string, courseID int) (GameID, error) {
	// TODO: create init function for scorecard struct
	// TODO: create a function which generates gameIDs
	var gameID GameID = "FR1B4"
	_, exists := s.scorecards[gameID]
	if exists {
		return gameID, errors.New(fmt.Sprintf("There already exists a game for id %s", gameID))
	}
	s.scorecards[gameID] = &Scorecard{Leader: player, Scores: []Score{{player, make([]int, 0)}}, courseName: "Laajalahti", broker: Newserver()}
	return "FR1B4", nil
}

func (s *ScorecardManager) Test() map[GameID]*Scorecard {
	fmt.Println("Test()")
	fmt.Println(s.scorecards["FR1B4"].Scores[0])
	return s.scorecards
}
func (s *ScorecardManager) Join(id GameID, player string) error {
	var err error
	card, ok := s.scorecards[id]
	if ok {
		//err = card.addPlayer(player)
		err = card.update(ScorecardStateUpdate{
			Originator:     player,
			UpdateType:     "player-joined",
			GameID:         id,
			ScorecardState: Score{player, make([]int, 0)},
		})
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *ScorecardManager) Delete() error {
	return errors.New("not implemented")
}

func (s *ScorecardManager) Update(state ScorecardStateUpdate) error {
	err := s.scorecards[state.GameID].update(state)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func (s *ScorecardManager) addPlayer(id GameID, player string) error {
	if s.scorecards[id] == nil {
		return errors.New(fmt.Sprintf("There is no active games for id: %v", id))
	}
	s.scorecards[id].Scores = append(s.scorecards[id].Scores, Score{player, []int{}})
	return nil
}

func (s *ScorecardManager) RemovePlayer() error {
	return errors.New("not implemented")
}

type Scorecard struct {
	Leader     string
	courseName string
	Scores     []Score `json:"scores"`
	mu         sync.Mutex
	broker     *Broker
}

func (s *Scorecard) addPlayer(player string) error {
	if s.Scores == nil {
		return errors.New("scorecard has no scores field")
	}
	s.Scores = append(s.Scores, Score{player, make([]int, 18)})
	return nil
}

func (s *Scorecard) update(state ScorecardStateUpdate) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf(", %v", s.Scores)
	switch state.UpdateType {
	case "score-update":
		for i, v := range s.Scores {
			if v.Player == state.ScorecardState.Player {
				fmt.Println("updating", v.Player, "value", state)
				s.Scores[i].Score = state.ScorecardState.Score
				notify(state, s.broker)
				return nil
			}
		}
	case "player-joined":
		s.Scores = append(s.Scores, Score{state.Originator, make([]int, 18)})
		notify(state, s.broker)
		return nil
	}

	return errors.New(fmt.Sprintf("Scorecard state update failed. State which failed\n %v", state))
}

func notify(state any, broker *Broker) {
	j, _ := json.Marshal(state)
	broker.Notifier <- j
}
