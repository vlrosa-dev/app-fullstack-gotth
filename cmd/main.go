package main

import (
	"app-fullstack-gotth/database/db"
	"database/sql"
	"fmt"
	"log"
	"os"

	utils "app-fullstack-gotth/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	// loaf environment variables
	godotenv.Load()

	// connect to database
	postgresURI := os.Getenv("DATABASE_URL")
	dbConnection, err := sql.Open("postgres", postgresURI)
	if err != nil {
		log.Panic(err)
	}

	// check db is up
	err = dbConnection.Ping()
	if err != nil {
		dbConnection.Close()
		log.Panic(err)
	}
	fmt.Println("Connected to database")

	// files static
	e.Static("/static", "static")
	e.HTTPErrorHandler = utils.CustomHTTPErrorHandler

	// queries db
	queries := db.New(dbConnection)
	fmt.Println(queries)

	// Start Server
	e.Logger.Fatal(e.Start(":8082"))
}
