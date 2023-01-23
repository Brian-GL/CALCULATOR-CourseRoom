package models

type TareaCalificacionInputModel struct {
	IdUsuario    int
	IdTarea      int
	IdCurso      int
	Calificacion float32
	Puntualidad  float32
	SECRET_TOKEN string
}
