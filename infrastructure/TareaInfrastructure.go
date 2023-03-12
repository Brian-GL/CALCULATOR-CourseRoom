package infrastructure

import (
	"calculator-courseroom/entities"
	"calculator-courseroom/models"

	"gorm.io/gorm"
)

func InformacionDesempenoUsuarioGetAsync(db *gorm.DB, IdUsuario int) []entities.CalculatorInformacionDesempenoObtenerEntity {

	var resultado []entities.CalculatorInformacionDesempenoObtenerEntity

	if db != nil {

		exec := "EXEC dbo.CalculatorInformacionDesempeno_Obtener @IdUsuario = ?"

		db.Raw(exec, IdUsuario).Scan(&resultado)
	}

	return resultado
}

func UsuarioDesempenoActualizarPutAsync(db *gorm.DB, model *models.UsuarioDesempenoActualizarInputModel) entities.AccionEntity {

	var resultado entities.AccionEntity
	if db != nil {

		exec := "EXEC dbo.UsuarioDesempeno_Actualizar @IdDesempeno = ?, @PrediccionCalificacionCurso = ?, @RumboCalificacionCurso = ?"
		db.Raw(exec, model.IdDesempeno, model.PrediccionCalificacionCurso, model.RumboCalificacionCurso).Scan(&resultado)

	} else {
		resultado = entities.AccionEntity{
			Codigo:  -1,
			Mensaje: "La base de datos no es accesible por el momento",
		}
	}

	return resultado
}
