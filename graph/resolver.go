package graph

import (
	"Bus_Booking/graph/model"

	"github.com/go-pg/pg"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateUser map[string]model.User
	DB         *pg.DB
}
