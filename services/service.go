package services

import (
	"database/sql"

	"golang.org/x/oauth2"
)

type MyAppService struct {
	db     *sql.DB
	config oauth2.Config
}

func NewMyAppService(db *sql.DB, config oauth2.Config) *MyAppService {
	return &MyAppService{db: db, config: config}
}
