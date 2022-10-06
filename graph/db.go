// db.go

package graph

import (
	"fmt"

	"github.com/go-pg/pg"
)

func Connected() *pg.DB {
	connStr := "host=localhost user=postgres password=welcome dbname=Bus_Booking port=5432 sslmode=disable "
	db, err := pg.ParseURL(connStr)
	DB := pg.Connect(db)
	if err == nil {
		fmt.Print("error")
	}

	return DB
}
