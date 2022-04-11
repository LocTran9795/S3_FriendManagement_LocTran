package configs

import (
	"database/sql"
	"fmt"
	"myapp/scripts/logger"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToDatabase() {
	connStr := "postgresql://postgres:Aa@123456@db:5432/mydb?sslmode=disable"
	// Connect to database
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.ErrorLogger.Println("Mysql config error", err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		logger.ErrorLogger.Println("Ping Error ", pingErr)
	}
	fmt.Println("Connected")
}

func GetDbAccess() *sql.DB {
	return db
}
