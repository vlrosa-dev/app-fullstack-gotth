package main

import (
	"app-fullstack-gotth/internal/handlers/authhandler"
	"app-fullstack-gotth/share"
	"database/sql"
	"fmt"
	"log"
	"os"

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
	e.HTTPErrorHandler = share.CustomHTTPErrorHandler

	// routes
	e.GET("/login", authhandler.NewGetLoginHandler().Serve)
	e.GET("/register", authhandler.NewGetRegisterHandler().Serve)

	// Start Server
	e.Logger.Fatal(e.Start(":8082"))
}
