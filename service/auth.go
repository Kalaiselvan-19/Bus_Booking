package service

import (
	"Bus_Booking/graph/model"
	"context"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

func UserRegister(ctx context.Context, input model.NewUser) (interface{}, error) {
	// Check Email

	createdUser, err := UserCreate(ctx, input)
	if err != nil {
		return nil, err
	}

	token, err := JwtGenerate(ctx, createdUser.Email)
	if err != nil {
		return nil, err
	}
	mail(input.Email, "Subject: Created Account fo Booking Ticket\n", "U have Successfully Created an account ")

	return map[string]interface{}{
		"userdetails": &createdUser,
		"token":       token,
	}, nil
}

func UserLogin(ctx context.Context, Email string, Password string) (interface{}, error) {
	getUser, err := UserGetByEmail(ctx, Email)
	if err != nil {
		// if user not found
		if err == gorm.ErrRecordNotFound {
			return nil, &gqlerror.Error{
				Message: "Email not found",
			}
		}
		return nil, err
	}

	if err := ComparePassword(getUser.Password, Password); err != nil {
		return nil, err
	}
	token, err := JwtGenerate(ctx, getUser.Email)
	if err != nil {
		return nil, err
	}
	mail(Email, "Subject: Created Account fo Booking Ticket\n", "Login u r account successfully")
	return map[string]interface{}{
		"token": token,
	}, nil

}
