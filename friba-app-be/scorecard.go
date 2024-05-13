package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (s *Server) createScorecardHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CourseID int    `json:"courseID"`
		Player   string `json:"player"`
	}

	err := s.readJSON(w, r, &input)
	if err != nil {
		fmt.Println(w, err)
	}

	id, err := s.scorecardManager.Create(input.Player, input.CourseID)

	if err != nil {
		err = fmt.Errorf("couldn't create a new scorecard, error: %v", err)
		s.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/scorecard-updates/%s", id))

	err = s.writeJSON(w, http.StatusCreated, envelope{"id": id}, headers)
	if err != nil {
		s.serverErrorResponse(w, r, err)
	}
}

func (s *Server) joinScorecardHandler(w http.ResponseWriter, r *http.Request) {
	var input ScorecardStateUpdate
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		err = fmt.Errorf("invalid reqest body, error: %v", err)
		s.badRequestResponse(w, r, err)
		return
	}
	err = s.scorecardManager.Join(input.GameID, input.Originator)
	if err != nil {
		err = fmt.Errorf("joining the scorecard failed, error: %v", err)
		s.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	err = s.writeJSON(w, http.StatusCreated, nil, nil)
	if err != nil {
		s.serverErrorResponse(w, r, err)
	}
}

func (s *Server) testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testHandler()")
	s.writeJSON(w, http.StatusOK, envelope{"data": s.scorecardManager.Test()["FR1B4"]}, nil)
}

type ScorecardStateUpdate struct {
	ScorecardState Score  `json:"scorecardState"`
	Originator     string `json:"originator"`
	UpdateType     string `json:"updateType"` // player joined/left, score update, game ended
	GameID         GameID `json:"gameID"`
}

func (s *Server) updateScorecardHandler(w http.ResponseWriter, r *http.Request) {
	var input ScorecardStateUpdate
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		fmt.Println(err.Error())
		s.serverErrorResponse(w, r, errors.New("something wrong with the request"))
	}
	err = s.scorecardManager.Update(input)
	if err != nil {
		s.serverErrorResponse(w, r, err)
	} else {
		// TODO: move this logic to scorecard file?
		// broadcast the update to all connected clients
		//j, _ := json.Marshal(s.scorecardManager.scorecards[input.GameID].Scores)
		//s.scorecardManager.scorecards[input.GameID].broker.Notifier <- j
		err = s.writeJSON(w, http.StatusOK, nil, nil)
		if err != nil {
			s.serverErrorResponse(w, r, err)
		}
	}
}
