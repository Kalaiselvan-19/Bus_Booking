package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"Name"`
	Email        string `gorm:"primaryKey" json:"Email" gorm:"unique"`
	Password     string `json:"Password"`
	Phone_Number int    `json:"Phone_Number"`
}
type Buses struct {
	gorm.Model
	Bus_Name    string `json:"Bus_Name"`
	Bus_Number  string `gorm:"primaryKey" ,json:"Bus_Number"`
	Total_seats int    `json:"Total_seats"`
	Seats       Seats  `gorm:"embedded"`
}
type Seats struct {
	Bus_Number  int  `json:"Bus_Number"`
	Seat_Number int  `gorm:"primaryKey" json:"Seat_Number"`
	Is_Booked   bool `Default:"False" json:"Is_Booked"`
	Price       int  `json:"Price"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
