package controllers

import (
	"calculator-courseroom/async"
	"calculator-courseroom/entities"
	"calculator-courseroom/infrastructure"
	"calculator-courseroom/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type RPCServer struct {
	DB           *gorm.DB
	SECRET_TOKEN string
}

func NewRPCServer() *RPCServer {

	// Cargar archivo .env
	_ = godotenv.Load(".env")

	//Cargar variables:

	server := os.Getenv("SERVER")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	databaseName := os.Getenv("DATABASE")
	SECRET_TOKEN := os.Getenv("SECRET_TOKEN")

	dsn := "sqlserver://" + user + ":" + password + "@" + server + "?database=" + databaseName

	db, _ := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	return &RPCServer{DB: db, SECRET_TOKEN: SECRET_TOKEN}
}

func (server *RPCServer) Calificacion(model *models.TareaCalificacionInputModel, reply *any) error {

	// Validar que el token sea el correcto:

	if server.SECRET_TOKEN == model.SECRET_TOKEN {
		future := async.Exec(func() interface{} {
			return infrastructure.TareaInformacionCalificacionGetAsync(server.DB, model)
		})

		//Obtener las estadisticas iniciales del usuario:
		estadisticasUsuario := future.Await().(entities.CalculatorInformacionTareaObtenerEntity)

		fmt.Println(estadisticasUsuario.NumeroTareasCalificadas)

		//Llamar script de matlab
	}

	return nil
}
