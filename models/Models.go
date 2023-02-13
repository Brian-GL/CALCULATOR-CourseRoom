package models

type TareaCalificacionInputModel struct {
	IdUsuario    int
	IdDesempeno  int
	SECRET_TOKEN string
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
