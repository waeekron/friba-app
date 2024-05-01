package main

import (
	"fmt"
	"log"
	"net/http"
)

// source https://gist.github.com/Ananto30/8af841f250e89c07e122e2a838698246
type Broker struct {
	Notifier chan []byte

	NewClients     chan chan []byte
	ClosingClients chan chan []byte
	Clients        map[chan []byte]bool
}

func Newserver() (broker *Broker) {
	broker = &Broker{
		Notifier:       make(chan []byte, 1),
		NewClients:     make(chan chan []byte),
		ClosingClients: make(chan chan []byte),
		Clients:        make(map[chan []byte]bool),
	}

	// set it running - listening and broadcasting events
	go broker.listen()
	return
}

type Message struct {
	MessageType string `json:"name"`
	Message     string `json:"msg"`
}

func (broker *Broker) Stream(w http.ResponseWriter, r *http.Request) {
	// check that w implements http.Flusher interface
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

//	func (broker *Broker) BroadcastMessage(w http.ResponseWriter, r *http.Request) {
//		var update ScorecardStateUpdate
//		_ = json.NewDecoder(r.Body).Decode(&update)
//		TODO validation
//		j, _ := json.Marshal(update)
//
// broker.Notifier <- j
// json.NewEncoder(w).Encode(update)
// broker.scorecardManager.Update("FR1B4", Score{"Miisu", []int{3}})
// update the in-memory state
// broker.scorecardManager.Update("FR1B4", update.ScorecardState)
// broker.scorecardManager.Update(input.GameID, input.Score)
// }
func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.NewClients:
			// A new client has connected, register their message channel
			broker.Clients[s] = true
			log.Printf("Client added. %d registered clients", len(broker.Clients))

		case s := <-broker.ClosingClients:
			// when a client disconnects we want to stop sending messages to them
			delete(broker.Clients, s)
			log.Printf("Removed client. %d registered clients", len(broker.Clients))

		case event := <-broker.Notifier:
			// we have a message which we want to send to all connected clients
			for clientMessageChan := range broker.Clients {
				clientMessageChan <- event
			}
		}
	}
}
