package metodos

import (
	"net/http"

	inputsMantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento/inputsMantenimiento"
	database "github.com/frank1995alfredo/api/database"
	mantenimiento "github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/gin-gonic/gin"
)

/**************METODOS PARA CIUDAD*******************/

//ObtenerCiudad ...
func ObtenerCiudad(c *gin.Context) {
	var ciudad []mantenimiento.Ciudad

	err := database.DB.Order("ciudad_id").Find(&ciudad).Error
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"data": ciudad})
	}
}

//CrearCiudad ... metodo para crear una provincia
func CrearCiudad(c *gin.Context) {

	var input inputsMantenimiento.CiudadInput
	var ciu mantenimiento.Ciudad
	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.VerificarDescripcion() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe ingregar una descripción y una provincia."})
		return
	}

	if err := database.DB.Where("descripcion=?", input.Descripcion).First(&ciu).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe esta ciudad, ingrese otra."})
		return
	}

	ciudad := mantenimiento.Ciudad{ProID: input.ProID, Descripcion: input.Descripcion, Estado: input.Estado}

	tx := database.DB.Begin()
	err := tx.Create(&ciudad).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusOK, gin.H{"data": ciudad})
}

//BuscarCiudad ...
func BuscarCiudad(c *gin.Context) {

	var ciu mantenimiento.Ciudad

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&ciu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe esta provincia."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": ciu})
}

//ActualizarCiudad ...
func ActualizarCiudad(c *gin.Context) {

	var input inputsMantenimiento.CiudadInput
	var ciu mantenimiento.Ciudad

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("ciudad_id=?", c.Param("id")).First(&ciu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provincia no encontrada."})
		return
	}

	ciudad := mantenimiento.Ciudad{ProID: input.ProID, Descripcion: input.Descripcion, Estado: input.Estado}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&ciu).Where("ciudad_id=?", c.Param("id")).Update(ciudad).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": ciudad})
}

//EliminarCiudad ...
func EliminarCiudad(c *gin.Context) {

	var ciudad mantenimiento.Ciudad

	if err := database.DB.Where("ciudad_id=?", c.Param("id")).First(&ciudad).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ciudad no existe."})
		return
	}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Where("ciudad_id=?", c.Param("id")).Delete(&ciudad).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}
