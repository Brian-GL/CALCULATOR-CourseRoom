package infrastructure

import (
	"calculator-courseroom/entities"
	"calculator-courseroom/models"

	"gorm.io/gorm"
)

func InformacionDesempenoUsuarioGetAsync(db *gorm.DB, model *models.TareaCalificacionInputModel) []entities.CalculatorInformacionDesempenoObtenerEntity {

	if db != nil {

		var resultado []entities.CalculatorInformacionDesempenoObtenerEntity

		exec := "EXEC dbo.CalculatorInformacionDesempeo_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)

		return resultado
	}

	return nil
}

func CalculatorRespuestaRegistrarPostAsync(db *gorm.DB, codigo *int, mensaje *string) *entities.AccionEntity {
	if db != nil {
		var resultCalculator *entities.AccionEntity
		exec := "EXEC dbo.CalculatorRespuesta_Registrar @Codigo = ?, @Mensaje = ?"
		db.Raw(exec, codigo, mensaje).Scan(&resultCalculator)
		return resultCalculator
	}

	return nil
}

func UsuarioDesempenoActualizarPutAsync(db *gorm.DB, model *models.UsuarioDesempenoActualizarInputModel) *entities.AccionEntity {
	if db != nil {
		var resultCalculator *entities.AccionEntity
		exec := "EXEC dbo.UsuarioDesempeno_Actualizar @IdDesempeno = ?, @PrediccionCalificacionCurso = ?, @RumboCalificacionCurso = ?, @PrediccionPuntualidadCurso = ?, @RumboPuntualidadCurso = ?"
		db.Raw(exec, model.IdDesempeno, model.PrediccionCalificacionCurso, model.RumboCalificacionCurso, model.PrediccionPuntualidadCurso, model.RumboPuntualidadCurso).Scan(&resultCalculator)
		return resultCalculator
	}

	return nil
}
