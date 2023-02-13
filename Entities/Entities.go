package entities

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
