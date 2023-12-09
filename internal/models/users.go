package models

import (
	"time"
)

type (
	User struct {
		Id        string     `json:"id"`
		Email     string     `json:"email"`
		Password  string     `json:"password"`
		Balance   int        `json:"balance"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	UserLogin struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
)
