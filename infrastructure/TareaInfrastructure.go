package infrastructure

import (
	"calculator-courseroom/models"

	"gorm.io/gorm"
)

func InformacionDesempenoUsuarioGetAsync(db *gorm.DB, model *models.TareaCalificacionInputModel) []models.CalculatorInformacionDesempenoObtenerEntity {

	var resultado []models.CalculatorInformacionDesempenoObtenerEntity

	if db != nil {

		exec := "EXEC dbo.CalculatorInformacionDesempeno_Obtener @IdUsuario = ?"

		db.Raw(exec, model.IdUsuario).Scan(&resultado)
	}

	return resultado
}

func UsuarioDesempenoActualizarPutAsync(db *gorm.DB, model *models.UsuarioDesempenoActualizarInputModel) models.AccionEntity {

	var resultado models.AccionEntity
	if db != nil {

		exec := "EXEC dbo.UsuarioDesempeno_Actualizar @IdDesempeno = ?, @PrediccionCalificacionCurso = ?, @RumboCalificacionCurso = ?"
		db.Raw(exec, model.IdDesempeno, model.PrediccionCalificacionCurso, model.RumboCalificacionCurso).Scan(&resultado)

	} else {
		resultado = models.AccionEntity{
			Codigo:  -1,
			Mensaje: "La base de datos no es accesible por el momento",
		}
	}

	return resultado
}
