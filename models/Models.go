package models

type TareaCalificacionInputModel struct {
	IdUsuario   int
	IdDesempeno int
}

type BridgeModel struct {
	X []float32 `json:"x"`
	Y []float32 `json:"y"`
}

type UsuarioDesempenoActualizarInputModel struct {
	IdDesempeno                 int
	PrediccionCalificacionCurso *float64
	RumboCalificacionCurso      *string
}

type UsuarioDesempenoObtenerInputModel struct {
	IdUsuario int
}
