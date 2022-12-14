package service

import (
	"Bus_Booking/graph/model"
	"Bus_Booking/initializers"
	"context"
	"net/smtp"
)

func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	input.Password = HashPassword(input.Password)
	user := model.User{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UserGetByEmail(ctx context.Context, email string) (*model.User, error) {

	var user model.User

	if err := initializers.DB.Model(&user).Where("Email LIKE ?", email).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
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
