package initial

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres Dialect for Gorm
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	viper.AutomaticEnv()
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

//GetDB return db instance
func GetDB() *gorm.DB {

	if db != nil {
		return db
	}
	Init()
	return db
}

// Init returns db connection instance
func Init() *gorm.DB {
	var connection string

	if os.Getenv("ENV") == "TEST" {
		dbHost := os.Getenv("ASK_SQUARE_TEST_DB_HOST")
		port := os.Getenv(`ASK_SQUARE_TEST_DB_PORT`)
		dbPort, _ := strconv.Atoi(port)
		dbUser := os.Getenv(`ASK_SQUARE_TEST_DB_USER`)
		dbPass := os.Getenv(`ASK_SQUARE_TEST_DB_PASSWORD`)
		dbName := os.Getenv(`ASK_SQUARE_TEST_DB_NAME`)
		connection = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPass, dbName)
	} else {
		dbHost := viper.GetString("ASK_SQUARE_DB_HOST")
		dbPort := viper.GetInt(`ASK_SQUARE_DB_PORT`)
		dbUser := viper.GetString(`ASK_SQUARE_DB_USER`)
		dbPass := viper.GetString(`ASK_SQUARE_DB_PASSWORD`)
		dbName := viper.GetString(`ASK_SQUARE_DB_NAME`)
		connection = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPass, dbName)
	}

	db, err = gorm.Open("postgres", connection)
	if err != nil {
		log.Println(err)
		panic("Connection Failed to Open due to following error => ")

	} else {
		log.Println("Connection Established")
	}
	if os.Getenv("GO_ENV") != "TEST" {
		db.LogMode(true)
	}
	return db
}
