package database

import (
	"fmt"
	"assigment3/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "wheater"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	// db.AutoMigrate(&model.Employee{}) biar debug nya gak muncul
	db.Debug().AutoMigrate(&models.Environtment{})
}

func GetDB() *gorm.DB {
	return db
}
