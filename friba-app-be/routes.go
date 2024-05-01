package main

import (
	"fmt"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("POST /scorecard/create", s.createScorecardHandler)
	mux.HandleFunc("POST /scorecard/join", s.joinScorecardHandler)
	mux.HandleFunc("POST /scorecard/update", s.updateScorecardHandler)

	mux.HandleFunc("GET /scorecard-updates/{gameID}", s.subscribe)
	mux.HandleFunc("GET /scorecard/test", s.testHandler)
	return s.logRequest(mux)
}

func (s *Server) subscribe(w http.ResponseWriter, r *http.Request) {
	gameID := GameID(r.PathValue("gameID"))
	broker := s.scorecardManager.scorecards[gameID].broker
	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// each connection registers its own message channel with the Broker's connections registery
	messageChan := make(chan []byte)

	// Signal the broker that we have a new connection
	broker.NewClients <- messageChan

	// Remove this client from the map of connected clients when this function exists
	defer func() {
		broker.ClosingClients <- messageChan
	}()

	// Listen to connection close and un-register messageChan
	ctx := r.Context()
	notify := ctx.Done()
	go func() {
		<-notify
		broker.ClosingClients <- messageChan
	}()

	for {
		// Write to the ReponseWriter sse
		fmt.Fprintf(w, "data: %s\n\n", <-messageChan)

		// Flush the data immediately instead of buffering it for later
		flusher.Flush()
	}
}

/*
	func (broker *Broker) BroadcastMessage(w http.ResponseWriter, r *http.Request) {
		var update ScorecardStateUpdate
		_ = json.NewDecoder(r.Body).Decode(&update)
		// TODO validation
		j, _ := json.Marshal(update)

		broker.Notifier <- j
		json.NewEncoder(w).Encode(update)
		//broker.scorecardManager.Update("FR1B4", Score{"Miisu", []int{3}})
		// update the in-memory state
		broker.scorecardManager.Update("FR1B4", update.ScorecardState)
		//broker.scorecardManager.Update(input.GameID, input.Score)
	}
*/
func (s *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
