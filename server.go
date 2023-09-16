package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	database "usergraphql/db"
	"usergraphql/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("err load env file: " + err.Error())
	}

	log.Println(os.Getenv("SECRET_KEY"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.InitDB()
	defer database.CloseDB()

	chi := chi.NewRouter()

	//chi.Use(middleware.AuthMiddleware())
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	chi.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	chi.Handle("/graphql", srv)

	fmt.Printf("connect to http://localhost:%s/ for GraphQL playground\n", port)

	log.Fatal(http.ListenAndServe(":"+port, chi))
}
