package main

import (
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
	s.scorecards["FR1B4"] = &Scorecard{Leader: player, Scores: []Score{{player, make([]int, 0)}}, courseName: "Laajalahti", broker: Newserver()}
	return "FR1B4", nil
}

func (s *ScorecardManager) Test() map[GameID]*Scorecard {

	fmt.Println("Test()")
	fmt.Println(s.scorecards["FR1B4"].Scores[0])
	return s.scorecards
}
func (s *ScorecardManager) Join(id GameID, player string) {
	s.addPlayer(id, player)
}

func (s *ScorecardManager) Delete() bool {
	return true
}

func (s *ScorecardManager) Update(state ScorecardStateUpdate) bool {

	fmt.Println("scorecard manager update")
	fmt.Println(state.ScorecardState)
	ok := s.scorecards[state.GameID].Update(state)
	if ok {
		return true
	}
	return false
}

func (s *ScorecardManager) addPlayer(id GameID, player string) bool {
	if s.scorecards[id] == nil {
		return false
	}
	s.scorecards[id].Scores = append(s.scorecards[id].Scores, Score{player, []int{}})
	fmt.Println(s.scorecards[id])
	return true
}

func (s *ScorecardManager) RemovePlayer() bool {
	return true
}

type Scorecard struct {
	Leader     string
	courseName string
	Scores     []Score `json:"scores"`
	mu         sync.Mutex
	broker     *Broker
	// also needs a broker
}

func (s *Scorecard) Update(state ScorecardStateUpdate) bool {
	//fmt.Printf("scorecard update, %v \n", score)
	fmt.Printf(", %v", s.Scores)
	for i, v := range s.Scores {
		if v.Player == state.ScorecardState.Player {
			fmt.Println("updating", v.Player, "value", state)
			//v.Score = score.Score
			s.Scores[i].Score = state.ScorecardState.Score
			return true
		}
	}
	return false
}
