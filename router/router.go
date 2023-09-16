package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"usergraphql/graph"
)

func GraphQLHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func PlayGroundHandler() gin.HandlerFunc {
	p := playground.Handler("GraphQL playground", "/graphql")
	return func(c *gin.Context) {
		p.ServeHTTP(c.Writer, c.Request)
	}
}
