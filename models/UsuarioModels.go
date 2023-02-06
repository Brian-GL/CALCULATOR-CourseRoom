package models

type UsuarioDesempenoActualizarInputModel struct {
	IdDesempeno                 int
	PrediccionCalificacionCurso *float32
	RumboCalificacionCurso      *string
	PrediccionPuntualidadCurso  *float32
	RumboPuntualidadCurso       *string
}
