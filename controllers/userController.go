package controllers

import (
	"Bus_Booking/initializers"
	"Bus_Booking/models"
	"fmt"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func SignUp(c *gin.Context) {

	//Get data off req body
	var body struct {
		Name         string
		Email        string
		Password     string
		Phone_Number int
	}
	c.Bind(&body)
	// Create a post
	newuser := models.User{Name: body.Name, Email: body.Email, Password: body.Password, Phone_Number: body.Phone_Number}
	result := initializers.DB.Create(&newuser)
	mail(newuser.Email, "Subject: Created Account fo Booking Ticket\n", "U have Successfully Created an account ")
	if result.Error != nil {
		c.Status(400)
	}

	// return it
	c.JSON(200, gin.H{
		"message": newuser,
	})
}
func AllUser(c *gin.Context) {
	// Get the User
	var users []models.User
	initializers.DB.Find(&users)
	//Response
	c.JSON(200, gin.H{
		"message": users,
	})

}
func One_User(c *gin.Context) {
	if !IsAuthorized(c.GetHeader("token"), false) {
		c.JSON(200, gin.H{
			"message": "You are not Authorized",
		})
		return
	}
	id := c.Param("id")
	// Get the User
	var users models.User
	initializers.DB.Find(&users, id)
	//Response
	c.JSON(200, gin.H{
		"message": users,
	})

}
func Update_user(c *gin.Context) {
	if !IsAuthorized(c.GetHeader("token"), true) {
		c.JSON(200, gin.H{
			"message": "You not an admin",
		})
		return
	}
	email := c.Param("email")
	var body struct {
		Name         string
		Email        string
		Password     string
		Phone_Number int
	}
	c.Bind(&body)
	var user models.User
	initializers.DB.First(&user, email)
	initializers.DB.Model(&user).Updates(models.User{
		Name:         body.Name,
		Email:        body.Email,
		Password:     body.Password,
		Phone_Number: body.Phone_Number,
	})
	mail(user.Email, "Subject: Updated Account \n", "U have Successfully Update Your account ")
	c.JSON(200, gin.H{
		"message": user,
	})

}
func DeleteUser(c *gin.Context) {
	if !IsAuthorized(c.GetHeader("token"), true) {
		c.JSON(200, gin.H{
			"message": "You not an admin",
		})
		return
	}
	id := c.Param("id")
	initializers.DB.Delete(&models.User{}, id)
	c.JSON(200, gin.H{
		"message": "Deleted Successfully",
	})

}
func GenerateJWT(email string) (string, error) {

	var mySigningKey = []byte("my_secret_key")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = "user"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
	}

	return tokenString, nil

}

func SignIn(c *gin.Context) {

	var authdetails models.Authentication

	c.Bind(&authdetails)

	var authUser models.User
	initializers.DB.Where("email = ?", authdetails.Email).First(&authUser)
	if authUser.Email == "" {
		fmt.Println("error in getting email from user")
		return
	}

	check := authUser.Password == authdetails.Password

	if check == false {
		fmt.Println("error in check")
	}

	validToken, err := GenerateJWT(authUser.Email)

	if err != nil {
		fmt.Println("error in token generating")
	}

	var token models.Token

	token.Email = authUser.Email
	token.TokenString = validToken

	c.JSON(200, gin.H{
		"token": token,
	})

}

func IsAuthorized(usertoken string, check bool) bool {
	var signinkey = []byte("my_secret_key")

	token, err := jwt.Parse(usertoken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return signinkey, nil
	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if !check && claims["role"] == "user" {
			println("Valid User")
			return true
		}
		if check && claims["role"] == "admin" {
			println("Valid User")
			return true
		}
	}
	return false
}

func mail(EmailAddress, Subject, Body string) {

	from := "kalaiselvan1360@gmail.com"
	password := "pxhloscpfdlmyrmq"

	toEmailAddress := EmailAddress
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := Subject
	body := Body
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}
}
