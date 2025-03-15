package routing

import "time"

type PlayingState struct {
	IsPaused bool
}

type GameLog struct {
	CurrentTime time.Time
	Message     string
	Username    string
}

type MoveData struct {
	Username    string `json:"username"`
	Destination string `json:"destination"`
	UnitIDs     []int  `json:"unit_ids"`
}
