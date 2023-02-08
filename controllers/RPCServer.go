package controllers

import (
	"bytes"
	"calculator-courseroom/async"
	"calculator-courseroom/entities"
	"calculator-courseroom/infrastructure"
	"calculator-courseroom/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type RPCServer struct {
	DB           *gorm.DB
	BRIDGE       *string
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
	BRIDGE := os.Getenv("BRIDGE")

	dsn := "sqlserver://" + user + ":" + password + "@" + server + "?database=" + databaseName

	db, _ := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	return &RPCServer{DB: db, BRIDGE: &BRIDGE, SECRET_TOKEN: SECRET_TOKEN}
}

func (server *RPCServer) Calificacion(model *models.TareaCalificacionInputModel, reply *any) error {

	// Validar que el token sea el correcto:

	if server.SECRET_TOKEN == model.SECRET_TOKEN {
		future := async.Exec(func() interface{} {
			return infrastructure.InformacionDesempenoUsuarioGetAsync(server.DB, model)
		})

		//Obtener las estadisticas iniciales del usuario:
		estadisticasUsuario := future.Await().([]entities.CalculatorInformacionDesempenoObtenerEntity)

		if estadisticasUsuario != nil {
			if len(estadisticasUsuario) > 0 {

				var array_calificaciones_x []float32
				var array_calificaciones_y []float32
				var array_puntualidades_x []float32
				var array_puntualidades_y []float32

				for index, value := range estadisticasUsuario {

					array_calificaciones_x = append(array_calificaciones_x, (float32)(index))
					array_calificaciones_y = append(array_calificaciones_y, value.ResultadoCalificacionCurso)
					array_puntualidades_x = append(array_puntualidades_x, (float32)(index))
					array_puntualidades_y = append(array_puntualidades_y, value.ResultadoPuntualidadCurso)
				}

				//Llamar script de matlab

				modelBridge := models.BridgeModel{
					X: array_calificaciones_x,
					Y: array_calificaciones_y,
				}

				jsonValue, _ := json.Marshal(modelBridge)

				resp, err := http.Post(*server.BRIDGE+"RegresionPolinomial", "application/json", bytes.NewBuffer(jsonValue))
				if err != nil {
					future = async.Exec(func() interface{} {
						return infrastructure.CalculatorRespuestaRegistrarPostAsync(server.DB, 500, "Se presentó un error al conectar al bridge de CourseRoom Calculator")
					})
				} else {
					defer resp.Body.Close()

					body, err := ioutil.ReadAll(resp.Body)

					if err != nil {
						future = async.Exec(func() interface{} {
							return infrastructure.CalculatorRespuestaRegistrarPostAsync(server.DB, 500, "Se presentó un error al leer la respuesta del bridge de CourseRoom Calculator")
						})
					} else {
						var bridgeResponse entities.BridgeEntity
						err = json.Unmarshal(body, &bridgeResponse)
						future = async.Exec(func() interface{} {
							return infrastructure.CalculatorRespuestaRegistrarPostAsync(server.DB, bridgeResponse.Codigo, "OK")
						})
					}
				}

			}
		}

	}

	return nil
}
