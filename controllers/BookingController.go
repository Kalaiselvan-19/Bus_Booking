package controllers

import (
	"Bus_Booking/initializers"
	"Bus_Booking/models"

	"github.com/gin-gonic/gin"
)

func BusBooking(c *gin.Context) {
	var body struct {
		Email       string
		Bus_Number  string
		Seat_Number int
	}

	c.Bind(&body)
	booking := models.Booking{Email: body.Email, Bus_Number: body.Bus_Number, Seat_Number: body.Seat_Number}
	result := initializers.DB.Create(&booking)
	if result.Error != nil {
		c.Status(400)
	}
	mail(body.Email, "Subject: Ticket Confirmed\n", "U have Successfully Booked the Ticket  ")
	// return it
	c.JSON(200, gin.H{
		"message": booking,
	})
}
func Update_Bookedseat(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Email       string
		Bus_Number  string
		Seat_Number int
	}
	c.Bind(&body)
	var seat models.Booking
	initializers.DB.First(&seat, id)
	initializers.DB.Model(&seat).Updates(models.Booking{

		Email:       body.Email,
		Bus_Number:  body.Bus_Number,
		Seat_Number: body.Seat_Number,
	})
	mail(body.Email, "Subject: Ticket Updated\n", "U have Successfully Updated the Booked  Ticket  ")
	c.JSON(200, gin.H{
		"message": seat,
	})

}
func DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Booking{}, id)
	c.JSON(200, gin.H{
		"message": "Deleted Successfully",
	})

}
