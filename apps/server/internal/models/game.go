package models

import "time"

type GameStatus string

const (
	GameStatusInProgress GameStatus = "in_progress"
	GameStatusCompleted  GameStatus = "completed"
	GameStatusAbandoned  GameStatus = "abandoned"
)

type Game struct {
	ID        string     `json:"id" db:"id"`
	PlayerID  string     `json:"player_id" db:"player_id"`
	Status    GameStatus `json:"status" db:"status"`
	FENState  string     `json:"fen_state" db:"fen_state"`
	Money     int        `json:"money" db:"money"`
	Level     int        `json:"level" db:"level"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}