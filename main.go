package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Naokiiiiiii/BlogApiPractice/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {

	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", os.Getenv("USERNAME"), os.Getenv("USERPASS"), os.Getenv("DATABASE"))

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIANT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIANT_SECRET"),
		RedirectURL:  "http://localhost:5173",
		Endpoint:     google.Endpoint,
		Scopes:       []string{"profile", "email"},
	}

	r := api.NewRouter(db, config)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
