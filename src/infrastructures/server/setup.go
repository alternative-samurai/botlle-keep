package server

import (
	"example/src/controllers"
	"example/src/infrastructures/generated"

	"github.com/99designs/gqlgen/graphql/handler"
)

func SetupServer() *handler.Server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &controllers.Resolver{}}))
	return srv
}

func aaa() {

}
