package metodos

import (
	"net/http"

	inputsMantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento/inputsMantenimiento"
	database "github.com/frank1995alfredo/api/database"
	mantenimiento "github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/gin-gonic/gin"
)

/***************METODO PARA CARGO EMPLEADO********************/

//ObtenerCargo ...
func ObtenerCargo(c *gin.Context) {
	var cargo []mantenimiento.CargoEmp

	err := database.DB.Order("cargo_emp_id").Find(&cargo).Error
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"data": cargo})
	}
}

//CrearCargo ...
func CrearCargo(c *gin.Context) {

	var input inputsMantenimiento.CargoInput
	var carg mantenimiento.CargoEmp

	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//valido que la descripcion no este vacia
	if input.Descripcion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción."})
		return
	}

	//pregunto si este cargo existe en la base de datos
	if err := database.DB.Where("descripcion=?", input.Descripcion).First(&carg).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe este cargo, ingrese otro."})
		return
	}

	cargo := mantenimiento.CargoEmp{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Create(&cargo).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": cargo})
}

//BuscarCargo ...
func BuscarCargo(c *gin.Context) {

	var cargo mantenimiento.CargoEmp

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&cargo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe este cargo."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": cargo})
}

//ActualizarCargo ...
func ActualizarCargo(c *gin.Context) {

	var input inputsMantenimiento.CargoInput
	var carg mantenimiento.CargoEmp

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if input.Descripcion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción."})
		return
	}

	if err := database.DB.Where("cargo_emp_id=?", c.Param("id")).First(&carg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cargo no encontrada."})
		return
	}

	cargo := mantenimiento.CargoEmp{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&carg).Where("cargo_emp_id=?", c.Param("id")).Update(cargo).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": cargo})
}

//EliminarCargo ...
func EliminarCargo(c *gin.Context) {

	var cargo mantenimiento.CargoEmp

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Where("cargo_emp_id=?", c.Param("id")).Delete(&cargo).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}
