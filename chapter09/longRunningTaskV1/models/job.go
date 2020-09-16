package models

import (
	"github.com/google/uuid"
	"time"
)

type Job struct {
	ID uuid.UUID `json:"uuid"`
	Type string `json:"type"`
	ExtraData interface{} `json:"extra_data"`
}

type Log struct {
	ClientTime time.Time `json:"client_time"`
}

type CallBack struct {
	CallbackURL string `json:"callback_url"`
}

type Mail struct {
	EmailAddress string `json:"email_address"`
}