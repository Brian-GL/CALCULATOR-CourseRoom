package rpcserver

import (
	"bytes"
	"calculator-courseroom/async"
	"calculator-courseroom/entities"
	"calculator-courseroom/infrastructure"
	"calculator-courseroom/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type RpcServer struct {
	DB     *gorm.DB
	BRIDGE *string
}

func NewRpcServer() *RpcServer {

	//godotenv.Load(".env")

	//Cargar variables:

	server := os.Getenv("SERVER")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	databaseName := os.Getenv("DATABASE")
	BRIDGE := os.Getenv("BRIDGE")

	dsn := "sqlserver://" + user + ":" + password + "@" + server + "?database=" + databaseName

	db, _ := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	return &RpcServer{DB: db, BRIDGE: &BRIDGE}
}

func (server *RpcServer) Calificacion(model *models.TareaCalificacionInputModel, reply *any) error {

	future := async.Exec(func() interface{} {
		return infrastructure.InformacionDesempenoUsuarioGetAsync(server.DB, model.IdUsuario)
	})

	//Obtener las estadisticas iniciales del usuario:
	estadisticasUsuario := future.Await().([]entities.CalculatorInformacionDesempenoObtenerEntity)

	if len(estadisticasUsuario) > 0 {

		var array_calificaciones_x []float32
		var array_calificaciones_y []float32

		for _, value := range estadisticasUsuario {

			array_calificaciones_x = append(array_calificaciones_x, float32(value.Indice))
			array_calificaciones_y = append(array_calificaciones_y, value.Resultado)
		}

		if len(estadisticasUsuario) == 1 {

			desempeno := estadisticasUsuario[0]
			prediccionCalificacion := desempeno.Resultado
			rumboCalificacionCurso := fmt.Sprintf("%f", desempeno.Resultado) + "x"

			//Si la respuesta es correcta actualizar el desempeño por lo que nos regresa el algoritmo inteligente:
			modelDesempenoActualizar := models.UsuarioDesempenoActualizarInputModel{
				IdDesempeno:                 model.IdDesempeno,
				PrediccionCalificacionCurso: &prediccionCalificacion,
				RumboCalificacionCurso:      &rumboCalificacionCurso,
			}

			future = async.Exec(func() interface{} {
				return infrastructure.UsuarioDesempenoActualizarPutAsync(server.DB, &modelDesempenoActualizar)
			})

			future.Await()

			*reply = "OK"
		} else {

			//Llamar script de matlab:

			modelBridge := models.BridgeModel{
				X: array_calificaciones_x,
				Y: array_calificaciones_y,
			}

			jsonValue, _ := json.Marshal(modelBridge)

			//Llamar al bridge:
			resp, err := http.Post(*server.BRIDGE+"RegresionPolinomial", "application/json", bytes.NewBuffer(jsonValue))

			if err == nil {

				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)

				if err == nil {

					//Obtener respuesta del bride como json:
					var bridgeResponse entities.BridgeEntity
					err = json.Unmarshal(body, &bridgeResponse)

					if err == nil {

						if bridgeResponse.Codigo > 0 {

							//Si la respuesta es correcta actualizar el desempeño por lo que nos regresa el algoritmo inteligente:
							modelDesempenoActualizar := models.UsuarioDesempenoActualizarInputModel{
								IdDesempeno:                 model.IdDesempeno,
								PrediccionCalificacionCurso: &bridgeResponse.Resultado,
								RumboCalificacionCurso:      &bridgeResponse.Mensaje,
							}

							future = async.Exec(func() interface{} {
								return infrastructure.UsuarioDesempenoActualizarPutAsync(server.DB, &modelDesempenoActualizar)
							})

							future.Await()

							*reply = bridgeResponse.Mensaje

						} else {
							return errors.New(bridgeResponse.Mensaje)
						}
					} else {
						return err
					}
				} else {
					return err
				}
			} else {
				return err
			}
		}
	} else {
		return errors.New("no se encontraron registros")
	}

	return nil
}

func (server *RpcServer) InformacionDesempeno(args *models.UsuarioDesempenoObtenerInputModel, reply *entities.ResponseInfrastructure) error {

	if args != nil {

		future := async.Exec(func() interface{} {
			return infrastructure.InformacionDesempenoUsuarioGetAsync(server.DB, args.IdUsuario)
		})

		//Obtener las estadisticas iniciales del usuario:
		estadisticasUsuario := future.Await().([]entities.CalculatorInformacionDesempenoObtenerEntity)

		if len(estadisticasUsuario) > 0 {
			*reply = entities.ResponseInfrastructure{Status: entities.SUCCESS, Data: estadisticasUsuario}
		} else {
			*reply = entities.ResponseInfrastructure{Status: entities.ALERT, Data: "No se encontraron registros"}
		}

	} else {
		*reply = entities.ResponseInfrastructure{Status: entities.ALERT, Data: "El parámetro de entrada se encuentra nulo"}
	}

	return nil
}
