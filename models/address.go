package models

import (
	"encoding/json"
	"time"
)

type Address struct {
	ID      int       `json:"id" db:"id"`
	UserID  int       `json:"user_id" db:"user_id"`
	Address string    `json:"address" db:"address"`
	Title   string    `json:"title" db:"title"`
	Created time.Time `json:"created" db:"created"`
}

type AddressUser struct {
	Addresses json.RawMessage `json:"addresses" db:"addresses"`
}

type AddressCreate struct {
	Address string `json:"address" db:"address" validate:"required"`
	Title   string `json:"title" db:"title" validate:"required"`
}

type AddressUpdate struct {
	Address string `json:"address" db:"address" validate:"required"`
	Title   string `json:"title" db:"title" validate:"required"`
}
