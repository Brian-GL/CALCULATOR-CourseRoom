package infrastructure

import (
	"calculator-courseroom/entities"
	"calculator-courseroom/models"

	"gorm.io/gorm"
)

func TareaInformacionCalificacionGetAsync(db *gorm.DB, model *models.TareaCalificacionInputModel) *entities.CalculatorInformacionTareaObtenerEntity {

	if db != nil {

		var resultado entities.CalculatorInformacionTareaObtenerEntity

		exec := "EXEC dbo.CalculatorInformacion_Obtener @IdUsuario = ?, @IdCurso = ?, @Calificacion = ?, @Puntualidad = ?"

		db.Raw(exec, model.IdUsuario, model.IdCurso, model.Calificacion, model.Puntualidad).Scan(&resultado)

		return &resultado
	}

	return nil
}

func CalculatorRespuestaRegistrarPostAsync(db *gorm.DB, codigo *int, mensaje *string) entities.AccionEntity {
	var resultCalculator entities.AccionEntity
	exec := "EXEC dbo.CalculatorRespuesta_Registrar @Codigo = ?, @Mensaje = ?"
	db.Raw(exec, codigo, mensaje).Scan(&resultCalculator)
	return resultCalculator
}
