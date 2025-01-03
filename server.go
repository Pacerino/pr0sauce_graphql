package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/joho/godotenv"
	"github.com/pacerino/pr0sauce_graphql/graph"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func envLoad() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		panic(err)
	}
}

func main() {

	envLoad()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/query", srv)

	log.Printf("Started Service under Port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
