package main

import (
	"fmt"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/XoronEdge/quizzifire/graph"
	"github.com/XoronEdge/quizzifire/graph/generated"
	userRepo "github.com/XoronEdge/quizzifire/internal/user/repo/postgres"
	userUsecase "github.com/XoronEdge/quizzifire/internal/user/usecase"
	"github.com/jinzhu/gorm"
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
	dbHost := viper.GetString("DBHOST")
	dbPort := viper.GetInt(`DBPORT`)
	dbUser := viper.GetString(`DBUSER`)
	dbPass := viper.GetString(`DBPASS`)
	dbName := viper.GetString(`DBNAME`)
	connection := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	fmt.Println(connection)
	dbConn, err := gorm.Open("postgres", connection)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	dbConn.LogMode(true)
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

	log.Fatal(e.Start(viper.GetString("ADDRESS")))

}
