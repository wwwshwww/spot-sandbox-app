package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/wwwwshwww/spot-sandbox/external"
	"github.com/wwwwshwww/spot-sandbox/graph"
	"github.com/wwwwshwww/spot-sandbox/graph/generated"
	"github.com/wwwwshwww/spot-sandbox/loader"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := external.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db.Debug()}}))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	dataloaderMw := loader.DataLoaderMiddleware(db)

	r.Get("/", playground.Handler("GraphQL playground", "/query"))
	r.With(dataloaderMw).Post("/query", srv.ServeHTTP)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
