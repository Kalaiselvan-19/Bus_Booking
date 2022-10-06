package controllers

import (
	"Bus_Booking/initializers"
	"Bus_Booking/models"

	"github.com/gin-gonic/gin"
)

func Createbus(c *gin.Context) {
	var body struct {
		Bus_Name    string
		Bus_Number  string
		Total_seats int
		Seats       models.Seats
	}

	c.Bind(&body)
	var seat struct {
		Seat_Number int
		Is_Booked   bool
		Price       int
	}
	// Create a post
	newbus := models.Buses{Bus_Name: body.Bus_Name, Bus_Number: body.Bus_Number, Total_seats: body.Total_seats, Seats: models.Seats{Seat_Number: seat.Seat_Number, Is_Booked: seat.Is_Booked, Price: seat.Price}}
	result := initializers.DB.Create(&newbus)
	if result.Error != nil {
		c.Status(400)
	}
	// return it
	c.JSON(200, gin.H{
		"message": newbus,
	})
}
func One_bus(c *gin.Context) {
	id := c.Param("id")
	// Get the User
	var bus models.Buses
	initializers.DB.Find(&bus, id)
	//Response
	c.JSON(200, gin.H{
		"message": bus,
	})

}
func Add_seat(c *gin.Context) {
	var seat struct {
		Seat_Number int
		Bus_Number  string
		Is_Booked   bool
		Price       int
	}
	c.Bind(&seat)
	addseat := models.Seats{Seat_Number: seat.Seat_Number, Is_Booked: seat.Is_Booked, Price: seat.Price, Bus_Number: seat.Seat_Number}
	result := initializers.DB.Create(&addseat)
	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"message": addseat,
	})

}
func Update_Bus(c *gin.Context) {
	Bus_Number := c.Param("id")
	var body struct {
		Bus_Name    string
		Bus_Number  string
		Total_seats int
	}
	c.Bind(&body)
	var bus models.Buses
	initializers.DB.First(&bus, Bus_Number)
	initializers.DB.Model(&bus).Updates(models.Buses{
		Bus_Name:    body.Bus_Name,
		Bus_Number:  body.Bus_Number,
		Total_seats: body.Total_seats,
	})
	c.JSON(200, gin.H{
		"message": bus,
	})

}
func AllBus(c *gin.Context) {
	// Get the User
	var buses []models.Buses
	initializers.DB.Find(&buses)
	//Response
	c.JSON(200, gin.H{
		"message": buses,
	})

}
func DeleteBus(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Buses{}, id)
	c.JSON(200, gin.H{
		"message": "Deleted Successfully",
	})

}
func DeleteSeat(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Seats{}, id)
	c.JSON(200, gin.H{
		"message": "Deleted Successfully",
	})

}
