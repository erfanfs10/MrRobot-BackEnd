package db

import (
	"fmt"
	"github.com/erfanfs10/MrRobot-BackEnd/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func ConnectToDB() {
	dbUser := utils.GetEnv("DB_USER")
	dbPassword := utils.GetEnv("DB_PASSWORD")
	dbHost := utils.GetEnv("DB_HOST")
	dbPort := utils.GetEnv("DB_PORT")
	dbName := utils.GetEnv("DB_NAME")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		fmt.Println("Can not connect to DB")
		log.Fatalln(err.Error())
	}
	DB = db
	fmt.Println("Database Connected")

}
