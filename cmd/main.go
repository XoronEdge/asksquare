package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	wira "github.com/XoronEdge/asksquare/cmd/wire"
	"github.com/XoronEdge/asksquare/graph/generated"
	graph "github.com/XoronEdge/asksquare/graph/resolvers"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres Dialect for Gorm
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	os.Setenv("ENV", "DEVELOP")
	e := echo.New()
	di := wira.InitializeDi()

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Di: di}}))
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	//e.Start(":" + viper.GetString("ASK_SQUARE_ADDRESS"))
	log.Fatal(e.Start(viper.GetString("ASK_SQUARE_ADDRESS")))

}
