package main

import (
	"Bus_Booking/initializers"
	"Bus_Booking/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()

}
func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Buses{})
	initializers.DB.AutoMigrate(&models.Seats{})
	initializers.DB.AutoMigrate(&models.Booking{})

}
