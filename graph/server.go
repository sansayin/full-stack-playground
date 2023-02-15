package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
  "go-rest/db"
	"github.com/labstack/echo/v4"
)

type Graph struct {
}

func NewGraph(e *echo.Echo,da *db.Adapter) {
	graphqlHandler := handler.NewDefaultServer(
		NewExecutableSchema(
			Config{Resolvers: &Resolver{DB: da.Get()}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
