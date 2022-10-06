package main

import (
	"Bus_Booking/controllers"
	"Bus_Booking/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func Main() {
	r := gin.Default()
	r.POST("/newuser", controllers.SignUp)                        //createuser
	r.POST("/newbus", controllers.Createbus)                      //createbus
	r.GET("/all-user", controllers.AllUser)                       //alluser
	r.GET("/all-bus", controllers.AllBus)                         //allbus
	r.GET("/one-user/:id", controllers.One_User)                  //oneuser
	r.GET("/one-bus/:id", controllers.One_bus)                    //onebus
	r.POST("/signin", controllers.SignIn)                         //get token
	r.GET("/addseat", controllers.Add_seat)                       //add seat
	r.POST("/Bookseat", controllers.BusBooking)                   //bookseat
	r.PUT("/update-user/:email", controllers.Update_user)         //updateuser
	r.PUT("/update-bus/:id", controllers.Update_Bus)              //updatebus
	r.PUT("/update-seat/:id", controllers.Update_Bookedseat)      //updateseat
	r.DELETE("/delete/user/:id", controllers.DeleteUser)          //deleteuser
	r.DELETE("/delete/Bus/:id", controllers.DeleteBus)            //deletebus
	r.DELETE("/delete/Bookedseat/:id", controllers.DeleteBooking) //deletebookedseat
	r.Run()

}
