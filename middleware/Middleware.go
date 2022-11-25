package middleware

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Middleware struct {
	DB           *gorm.DB
	SECRET_TOKEN string
}

func NewMiddleware() *Middleware {

	// Cargar archivo .env
	err := godotenv.Load(".env")

	if err != nil {
		return nil
	}

	//Cargar variables:

	server := os.Getenv("SERVER")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	databaseName := os.Getenv("DATABASE")
	secretToken := os.Getenv("SECRET_TOKEN")

	dsn := "sqlserver://" + user + ":" + password + "@" + server + "?database=" + databaseName

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		return nil
	}

	return &Middleware{
		DB:           db,
		SECRET_TOKEN: secretToken}
}
