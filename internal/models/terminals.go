package models

import "time"

type (
	Terminal struct {
		Id         int        `json:"id"`
		Name       string     `json:"name"`
		LocationId int        `json:"location_id"`
		CreatedAt  time.Time  `json:"created_at"`
		UpdatedAt  *time.Time `json:"updated_at"`
	}

	Transaction struct {
		Id             int        `json:"id"`
		UserId         string     `json:"user_id"`
		FromTerminalId int        `json:"from_terminal_id"`
		ToTerminalId   *int       `json:"to_terminal_id"`
		IsSuccess      bool       `json:"is_success"`
		CreatedAt      time.Time  `json:"created_at"`
		UpdatedAt      *time.Time `json:"updated_at"`
	}

	UpdateTransaction struct {
		ToTerminalId int `json:"to_terminal_id"`
	}
)
