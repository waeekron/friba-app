package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) createScorecardHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CourseID int    `json:"courseID"`
		Player   string `json:"player"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	id, err := s.scorecardManager.Create(input.Player, input.CourseID)

	if err != nil {
		fmt.Fprintln(w, err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)
}

func (s *Server) joinScorecardHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Player string `json:"player"`
		GameID GameID `json:"gameID"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	s.scorecardManager.Join(input.GameID, input.Player)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}

func (s *Server) testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testHandler()")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.scorecardManager.Test()["FR1B4"])
}

type ScorecardStateUpdate struct {
	ScorecardState Score  `json:"scorecardState"`
	Originator     string `json:"Originator"`
	UpdateType     string `json:"updateType"` // player joined/left, score update, game ended
	GameID         GameID `json:"gameID"`
}

func (s *Server) updateScorecardHandler(w http.ResponseWriter, r *http.Request) {
	var input ScorecardStateUpdate
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	fmt.Println("IN update HANDLER", input)
	//s.scorecardManager.Update(input.GameID, input.ScorecardState)
	ok := s.scorecardManager.Update(input)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Something went wrong!")
	} else {
		// broadcast the update to all connected clients
		j, _ := json.Marshal(s.scorecardManager.scorecards[input.GameID].Scores)
		s.scorecardManager.scorecards[input.GameID].broker.Notifier <- j
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("State updated")
	}

}
