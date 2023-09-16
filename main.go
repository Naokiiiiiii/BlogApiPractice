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
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", os.Getenv("USERNAME"), os.Getenv("USERPASS"), os.Getenv("DATABASE"))

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
