## How to run
1. Install Go  version 1.22
2. run `go run .`

## Endpoints

#### Scorecard Service

|HTTP Method|URL|Description|
|---|---|---|
|`POST`|/scorecard/create | Create a new scorecard |
|`POST`| /scorecard/join | Join a scorecard |
|`POST`| /scorecard/update | Update a scorecard, updates are streamed to connected clients|

##### SSE Endpoint `scorecard-updates/{gameID}`

| Field         | Description                                              |
|---------------|----------------------------------------------------------|
| Event Stream  | The event stream containing real-time updates.           |
| Event Types   |player-join, score-update (for now)|
| Connection    | Persistent connection over HTTP.                         |
| Content-Type  | `text/event-stream`                                      |



## Project structure
```
.
├── broker.go # server-sent-event logic
├── errors.go
├── go.mod
├── helpers.go  
├── main.go # entry point
├── routes.go 
├── scorecard.go # scorecard route handlers
└── scorecardManager.go # business logic related to scorecards
```

### TODO

- add a db, tests, env file and validation.
