package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// _articleHttpDelivery "github.com/XoronEdge/quizzifire/internal/article/delivery/http"
	// _articleHttpDeliveryMiddleware "github.com/XoronEdge/quizzifire/internal/article/delivery/http/middleware"
	// _articleRepo "github.com/XoronEdge/quizzifire/internal/article/repository/mysql"
	// _articleUcase "github.com/XoronEdge/quizzifire/internal/article/usecase"
	// _authorRepo "github.com/XoronEdge/quizzifire/internal/author/repository/mysql"
)

func init() {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
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
	_, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// middL := _articleHttpDeliveryMiddleware.InitMiddleware()
	// e.Use(middL.CORS)
	// authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	// timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	// ar := _articleRepo.NewMysqlArticleRepository(dbConn)
	// au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	// _articleHttpDelivery.NewArticleHandler(e, au)
	log.Fatal(e.Start(viper.GetString("ADDRESS")))

}
