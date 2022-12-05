package infraestructure

import (
	"calculator-courseroom/entities"

	"gorm.io/gorm"
)

func TareaArchivosAdjuntosObtenerGetAsync(db *gorm.DB, IdTarea *int, IdUsuario *int) *[]entities.TareaImagenesEntregadasObtenerEntity {

	if db != nil {

		var resultado []entities.TareaImagenesEntregadasObtenerEntity

		exec := "EXEC dbo.TareaImagenesEntregadas_Obtener @IdTarea = ?, @IdUsuario = ?"

		db.Raw(exec, *IdTarea, *IdUsuario).Scan(&resultado)

		return &resultado

	}

	return nil

}
