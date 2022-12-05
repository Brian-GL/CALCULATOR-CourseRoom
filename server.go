package main

import (
	"calculator-courseroom/async"
	"calculator-courseroom/entities"
	"calculator-courseroom/infraestructure"
	"calculator-courseroom/models"
	"errors"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func NewServer() *Server {

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

	dsn := "sqlserver://" + user + ":" + password + "@" + server + "?database=" + databaseName

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		return nil
	}
	return &Server{DB: db}
}

func (server *Server) Calificacion(model *models.CalificacionInputModel, reply *string) error {

	// Obtener imagenes:

	future := async.Exec(func() interface{} {
		return infraestructure.TareaArchivosAdjuntosObtenerGetAsync(server.DB, &model.IdTarea, &model.IdUsuario)
	})

	imagenes := future.Await().([]entities.TareaImagenesEntregadasObtenerEntity)

	if imagenes != nil {

		*reply = "Se encontraron imágenes"
		return nil

	}

	*reply = "No se encontraron imágenes"
	return errors.New("No se encontraron imágenes")
}
