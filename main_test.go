package main

import (
	"Bus_Booking/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	return r
}
func TestSignUpHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/newuser")

	user := models.User{
		Name:         "Demo name",
		Email:        "demo@gmail.com",
		Password:     "demo",
		Phone_Number: 1234567890,
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/newuser", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetUserHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/all-user")
	req, _ := http.NewRequest("GET", "/all-user", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var user []*models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)

}
func TestUpdateuserHandler(t *testing.T) {
	r := SetUpRouter()
	r.PUT("/update-user/:email")
	user := models.User{
		Name:         "Updated demo name",
		Email:        "demo@gmail.com",
		Password:     "Updateddemo ",
		Phone_Number: 0,
	}
	jsonValue, _ := json.Marshal(user)
	reqFound, _ := http.NewRequest("PUT", "/update-user/"+user.Email, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("PUT", "/update-user/demooo@gmail.com", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

//bus
func TestCreatebusHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/newbus")

	user := models.Buses{
		Bus_Name:    "Demo bus",
		Bus_Number:  "123",
		Total_seats: 10,
		Seats: models.Seats{
			Seat_Number: 1,
			Price:       100,
			Is_Booked:   false,
		},
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/newbus", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdatebusHandler(t *testing.T) {
	r := SetUpRouter()
	r.PUT("/update-bus/:id")
	bus := models.Buses{
		Bus_Name:    "updated bus name",
		Bus_Number:  "123",
		Total_seats: 20,
		Seats:       models.Seats{},
	}

	jsonValue, _ := json.Marshal(bus)
	reqFound, _ := http.NewRequest("PUT", "/update-bus/"+bus.Bus_Number, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("PUT", "/update-bus/123", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
func TestGetbusHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/all-bus")
	req, _ := http.NewRequest("GET", "/all-bus", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var user []*models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, http.StatusOK, w.Code)

}
