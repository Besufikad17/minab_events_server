package web

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Besufikad17/minab_events/internal"
	"github.com/joho/godotenv"
)

type Application struct {
	Ctx           context.Context
	GraphQLClient internal.GraphQLClient
}

func Serve() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	ctx := context.Background()

	hasuraGraphQLURL := os.Getenv("HASURA_GRAPHQL_URL")
	if hasuraGraphQLURL == "" {
		log.Fatal("HASURA_GRAPHQL_URL not found")
	}

	hasuraAdminSecret := os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")
	if hasuraAdminSecret == "" {
		log.Fatal("HASURA_GRAPHQL_ADMIN_SECRET not found")
	}

	app := &Application{
		Ctx: ctx,
		GraphQLClient: *internal.NewGraphQLClient(
			map[string]string{
				"content-type":          "application/json",
				"x-hasura-admin-secret": hasuraAdminSecret,
			},
			hasuraGraphQLURL,
		),
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if port[0] != ':' {
		port = ":" + port
	}

	log.Println("Starting server on " + port)
	err = http.ListenAndServe(port, app.routes())
	if err != nil {
		log.Fatal("Error running server on "+port, err.Error())
	}
}
