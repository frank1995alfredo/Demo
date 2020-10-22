package metodos

import (
	"net/http"

	inputsmantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento/inputsmantenimiento"
	database "github.com/frank1995alfredo/api/database"
	mantenimiento "github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/gin-gonic/gin"
)

/**************METODOS PARA PROVINCIA*******************/

//ObtenerProvincia ... metodo para obtener todas las provincias
func ObtenerProvincia(c *gin.Context) {
	var provincia []mantenimiento.Provincia

	err := database.DB.Order("provincia_id").Find(&provincia).Error
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"data": provincia})
	}
}

//CrearProvincia ... metodo para crear una provincia
func CrearProvincia(c *gin.Context) {

	var input inputsmantenimiento.ProvinciaInput
	var provinc mantenimiento.Provincia
	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ValidarEntrada() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción"})
		return
	}

	if err := database.DB.Where("descripcion=?", input.Descripcion).First(&provinc).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe esta provincia, ingrese otra."})
		return
	}

	provincia := mantenimiento.Provincia{Descripcion: input.Descripcion, Estado: input.Estado}

	tx := database.DB.Begin()
	err := tx.Create(&provincia).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusOK, gin.H{"data": provincia})
}

//BuscarProvincia ...
func BuscarProvincia(c *gin.Context) {

	var provinc mantenimiento.Provincia

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&provinc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe esta provincia."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": provinc})
}

//ActualizarProvincia ...
func ActualizarProvincia(c *gin.Context) {

	var input inputsmantenimiento.ProvinciaInput
	var provinc mantenimiento.Provincia

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("provincia_id=?", c.Param("id")).First(&provinc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provincia no encontrada."})
		return
	}

	provin := mantenimiento.Provincia{Descripcion: input.Descripcion, Estado: input.Estado}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&provinc).Where("provincia_id=?", c.Param("id")).Update(provin).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": provinc})
}

//EliminarProvincia ...
func EliminarProvincia(c *gin.Context) {

	var provincia mantenimiento.Provincia

	if err := database.DB.Where("provincia_id=?", c.Param("id")).First(&provincia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provincia no existe"})
		return
	}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Where("provincia_id=?", c.Param("id")).Delete(&provincia).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}
