package main

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	port             int
	scorecardManager ScorecardManager
}

func NewServer() *http.Server {
	//port, _ := strconv.Atoi(os.Getenv("PORT"))

	NewServer := &Server{
		port:             8080,
		scorecardManager: ScorecardManager{map[GameID]*Scorecard{}},
	}

	server := &http.Server{
		Addr:        fmt.Sprintf(":%d", NewServer.port),
		Handler:     NewServer.RegisterRoutes(),
		IdleTimeout: time.Minute,
		//ReadTimeout:  10 * time.Second,
		//WriteTimeout: 30 * time.Second,
	}
	return server
}
func main() {
	s := NewServer()

	fmt.Println("Starting server on port :8080")
	err := s.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
