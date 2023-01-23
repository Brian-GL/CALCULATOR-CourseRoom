package models

type UsuarioDesempenoRegistrarInputModel struct {
	IdUsuario                    int
	IdTarea                      int
	Calificacion                 float64
	Puntualidad                  float64
	PromedioCurso                float64
	PrediccionPromedioCurso      *float64
	RumboPromedioCurso           *string
	PromedioGeneral              float64
	PrediccionPromedioGeneral    *float64
	RumboPromedioGeneral         *string
	PuntualidadCurso             float64
	PrediccionPuntualidadCurso   *float64
	RumboPuntualidadCurso        *string
	PuntualidadGeneral           float64
	PrediccionPuntualidadGeneral *float64
	RumboPuntualidadGeneral      *string
}
