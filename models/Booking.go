package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	Email       string `json:"Email"`
	Bus_Number  string `json:"Bus_Number"`
	Seat_Number int    `json:"Seat_Number"`
}
