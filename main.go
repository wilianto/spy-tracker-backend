package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/wilianto/spy-tracker-backend/user"
)

func main() {
	viper.AutomaticEnv()

	//get server config
	appPort := fmt.Sprintf(":%s", viper.GetString("PORT"))

	//get database config
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPass := viper.GetString("DB_PASS")
	dbName := fmt.Sprintf("%s_%s", viper.GetString("DB_NAME"), viper.GetString("ENV"))
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", dbName, dbUser, dbPass, dbHost, dbPort)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Fatalf("Error Connect to DB: %s", err.Error())
	}

	err = dbConn.Ping()
	if err != nil {
		logrus.Fatalf("Error PING to DB: %s", err.Error())
	}

	//handler router
	r := mux.NewRouter()

	//user endpoint
	userValidator := user.NewValidator()
	userRepository := user.NewPsqlRepository(dbConn)
	userService := user.NewService(userRepository, userValidator)
	user.NewHTTPHandler(r, userService)

	//start http server
	logrus.Infof("Starting server on :%s", appPort)
	logrus.Fatal(http.ListenAndServe(appPort, r))
}
