package infrastructure

import (
	"calculator-courseroom/entities"
	"calculator-courseroom/models"

	"gorm.io/gorm"
)

func UsuarioDesempenoRegistrarPostAsync(db *gorm.DB, model *models.UsuarioDesempenoRegistrarInputModel) *entities.AccionEntity {

	var resultado *entities.AccionEntity

	exec := "EXEC dbo.UsuarioDesempeno_Registrar @IdUsuario = ?, @IdTarea = ?, @Calificacion = ?, @Puntualidad = ?, @PromedioCurso = ?, @PrediccionPromedioCurso = ?, @RumboPromedioCurso = ?, @PromedioGeneral = ?, @PrediccionPromedioGeneral = ?, @RumboPromedioGeneral = ?, @PuntualidadCurso = ?, @PrediccionPuntualidadCurso = ?, @RumboPuntualidadCurso = ?, @PuntualidadGeneral = ?, @PrediccionPuntualidadGeneral = ?, @RumboPuntualidadGeneral = ?"

	db.Raw(exec, model.IdUsuario, model.IdTarea, model.Calificacion, model.Puntualidad, model.PromedioCurso, model.PrediccionPromedioCurso, model.RumboPromedioCurso, model.PromedioGeneral, model.PrediccionPromedioGeneral, model.RumboPromedioGeneral, model.PuntualidadCurso, model.PrediccionPuntualidadCurso, model.RumboPuntualidadCurso, model.PuntualidadGeneral, model.PrediccionPuntualidadGeneral, model.RumboPuntualidadGeneral).Scan(&resultado)

	return resultado

}
