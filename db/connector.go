package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type PostgresConnector struct {
}

func (p *PostgresConnector) GetConnection() (db *gorm.DB, err error) {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	username := "*"
	password := "*"
	dbHost := "vehicles-demo.*.us-east-1.rds.amazonaws.com"
	dbURI := fmt.Sprintf("host=%s user=%s password=%s", dbHost, username, password)
	return gorm.Open("postgres", dbURI)
}
