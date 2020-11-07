package main

import (
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/XoronEdge/asksquare/graph/generated"
	graph "github.com/XoronEdge/asksquare/graph/resolvers"
	"github.com/XoronEdge/asksquare/initial"
	userRepo "github.com/XoronEdge/asksquare/internal/user/repo/postgres"
	userUsecase "github.com/XoronEdge/asksquare/internal/user/usecase"
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
	dbConn := initial.GetDB()
	e := echo.New()
	ur := userRepo.NewUserRepo(dbConn)
	uc := userUsecase.NewUserUsecase(ur, time.Minute)

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Uc: uc}}))
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
