package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/babyninee/demoapi/course-go-for-developer/flight"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDatabase() *sql.DB {
	db, err := sql.Open(viper.GetString("db.driver"), viper.GetString("db.url"))

	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	fmt.Println("Connect to database success")
	return db
}

func main() {

	initConfig()

	db := initDatabase()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection database successfully")

	r := gin.Default()
	getHandler := flight.NewflightHandler(db)

	r.POST("/flight/create/:id", getHandler.CreateFlightHandler)
	r.POST("/flight/update/:id", getHandler.UpdateFlightHandler)
	r.GET("/flight/:id", getHandler.GetByIDHandler)
	r.GET("/flights", getHandler.GetAllHandler)
	r.DELETE("/flight/:id", getHandler.DeleteFlightByIDHandler)
	r.Run()
}
