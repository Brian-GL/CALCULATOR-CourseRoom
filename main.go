package main

import (
	"bytes"
	"calculator-courseroom/async"
	"calculator-courseroom/infrastructure"
	"calculator-courseroom/models"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"

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
	//godotenv.Load(".env")

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
		estadisticasUsuario := future.Await().([]models.CalculatorInformacionDesempenoObtenerEntity)

		if estadisticasUsuario != nil {
			if len(estadisticasUsuario) > 0 {

				var array_calificaciones_x []float32
				var array_calificaciones_y []float32

				for _, value := range estadisticasUsuario {

					array_calificaciones_x = append(array_calificaciones_x, float32(value.Indice))
					array_calificaciones_y = append(array_calificaciones_y, value.Resultado)
				}

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
						var bridgeResponse models.BridgeEntity
						err = json.Unmarshal(body, &bridgeResponse)

						if err == nil {

							if bridgeResponse.Codigo > 0 {

								//Si la respuesta es correcta actualizar el desempeño por lo que nos regresa el algoritmo inteligente:
								modelDesempenoActualizar := models.UsuarioDesempenoActualizarInputModel{
									IdDesempeno:                 model.IdDesempeno,
									PrediccionCalificacionCurso: &bridgeResponse.Resultado,
									RumboCalificacionCurso:      &bridgeResponse.Mensaje,
								}

								infrastructure.UsuarioDesempenoActualizarPutAsync(server.DB, &modelDesempenoActualizar)
							}
						}
					}
				}
			}
		}
	}

	return nil
}

func main() {
	fmt.Print("\033[H\033[2J")
	go Servering()
	fmt.Scanln()
	fmt.Print("\033[H\033[2J")
}

func Servering() {

	rpcServer := NewRPCServer()
	rpc.Register(rpcServer)

	server_rpc, error := net.Listen("tcp", ":2023")
	if error != nil {
		fmt.Println("Found Error : ", error.Error())
		server_rpc.Close()
		return
	}

	fmt.Println("\nCourseRoom Calculator Opened At " + time.Now().Format("Monday 2006-01-02 15:04"))

	//Home page:
	//http.HandleFunc("/rpc", Index)

	defer server_rpc.Close()

	for {
		fmt.Println("Listening...")
		client, err_or := server_rpc.Accept()
		if err_or != nil {
			fmt.Println("Found Error: ", err_or)
			continue
		}
		fmt.Println("Connected Client With Address:", client.RemoteAddr())
		go rpc.ServeConn(client)
	}
}

func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprint(
		res,
		LoadHtml("./public/index.html"),
	)
}

func LoadHtml(filename string) string {
	html, _ := os.ReadFile(filename)
	return string(html)
}
