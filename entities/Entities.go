package entities

type ResponseStatus int

const (
	SUCCESS = 200
	ALERT   = 404
	ERROR   = 500
)

type ResponseInfrastructure struct {
	Status ResponseStatus `json:"status"`
	Data   any            `json:"data"`
}

type AccionEntity struct {
	Codigo  int
	Mensaje string
}

type BridgeEntity struct {
	Codigo    int     `json:"codigo"`
	Mensaje   string  `json:"mensaje"`
	Resultado float32 `json:"resultado"`
}

type CalculatorInformacionDesempenoObtenerEntity struct {
	Indice     int
	Resultado  float32
	Prediccion float32
}
