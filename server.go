package main

import (
	"Bus_Booking/directives"
	"Bus_Booking/graph"
	"Bus_Booking/graph/generated"
	"Bus_Booking/initializers"
	migration "Bus_Booking/migration"
	"Bus_Booking/service"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	migration.MigrateTable()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := initializers.DB
	Sqldb, _ := db.DB()
	defer Sqldb.Close()
	router := mux.NewRouter()
	router.Use(service.AuthMiddleware)
	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
