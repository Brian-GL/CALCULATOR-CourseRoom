package entities

type CalculatorInformacionTareaObtenerEntity struct {
	PromedioGeneral         float32
	PuntualidadGeneral      float32
	PromedioCurso           float32
	PuntualidadCurso        float32
	VarianzaCalificacion    *float32
	VarianzaPuntualidad     *float32
	NumeroTareasCalificadas *int
}
