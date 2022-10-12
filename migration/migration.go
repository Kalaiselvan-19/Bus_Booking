package migration

import (
	"Bus_Booking/graph/model"
	"Bus_Booking/initializers"
	"Bus_Booking/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()

}
func MigrateTable() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Buses{})
	initializers.DB.AutoMigrate(&models.Seats{})
	initializers.DB.AutoMigrate(&models.Booking{})
	initializers.DB.AutoMigrate(&model.User{})
	initializers.DB.AutoMigrate(&model.Buses{})

}
