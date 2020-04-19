package model

import "time"

type Users struct {
	ID          int64
	UserName    string
	DisplayName string
	Password    string
	Email       string
	Bio         string
	Image       string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
