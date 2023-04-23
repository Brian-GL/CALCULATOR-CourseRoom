package main

import (
	"calculator-courseroom/rpcserver"
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func main() {

	//Cargar variables de entorno:
	//godotenv.Load(".env")

	jsonIter := jsoniter.ConfigCompatibleWithStandardLibrary
	secretToken := os.Getenv("SECRET_TOKEN")

	rpcServer := rpcserver.NewRpcServer()
	rpc.Register(rpcServer)

	//Home page:
	http.HandleFunc("/rpc", Index)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		defer req.Body.Close()

		// Cabecera de respuesta:
		w.Header().Add("Content-Type", "application/json")

		// Obtener token
		token := req.Header.Get("Authorization")

		// Validar que el token no se encuentre vacío:
		if token == "" {

			jsonBytes, err := jsonIter.Marshal("El token es necesario para acceder a este recurso")

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(jsonBytes)
			}

			return
		}

		// Validar que el token sea el correcto:

		if token == secretToken {

			switch req.Method {

			case "POST":
				{
					rpcServerRequest := rpcserver.NewRpcRequest(req.Body)
					res := rpcServerRequest.Call()
					io.Copy(w, res)
				}

			default:
				{
					jsonBytes, err := jsonIter.Marshal("Ruta inválida")

					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte(err.Error()))
					} else {
						w.WriteHeader(http.StatusNotImplemented)
						w.Write(jsonBytes)
					}
				}
			}

		} else {

			jsonBytes, err := jsonIter.Marshal("Token inválido")

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(jsonBytes)
			}
		}
	})

	fmt.Println("\nCourseRoom Calculator Opened At " + time.Now().Format("2006-01-02 15:04:05 Monday"))

	err := http.ListenAndServe(":1414", nil)
	if err != nil {
		panic(err)
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
