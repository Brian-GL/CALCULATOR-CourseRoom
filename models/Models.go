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

type AccionEntity struct {
	Codigo  int
	Mensaje string
}

type BridgeEntity struct {
	Codigo    int     `json:"codigo"`
	Mensaje   string  `json:"mensaje"`
	Resultado float64 `json:"resultado"`
}

type CalculatorInformacionDesempenoObtenerEntity struct {
	Indice    int
	Resultado float32
}
