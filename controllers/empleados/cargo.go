package empleados

import (
	"net/http"
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	database "github.com/frank1995alfredo/api/database"
	token "github.com/frank1995alfredo/api/functions"
	cargo "github.com/frank1995alfredo/api/models/empleados"
	"github.com/gin-gonic/gin"
)

/***************METODO PARA CARGO EMPLEADO********************/

//ObtenerCargo ...
func ObtenerCargo(c *gin.Context) {
	var cargo []cargo.CargoEmp

	token.ValidarToken()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	db := database.DB.Find(&cargo)

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"cargo_emp_id asc"},
		ShowSQL: false,
	}, &cargo)
	c.SecureJSON(http.StatusOK, gin.H{"data": paginator})

}

//CrearCargo ...
func CrearCargo(c *gin.Context) {

	var input CargoInput
	var carg cargo.CargoEmp

	token.ValidarToken()

	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//valido que la descripcion no este vacia
	if input.Descripcion == "" {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción."})
		return
	}

	//pregunto si este cargo existe en la base de datos
	if err := database.DB.Where("descripcion=?", input.Descripcion).First(&carg).Error; err == nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Ya existe este cargo, ingrese otro."})
		return
	}

	cargo := cargo.CargoEmp{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Create(&cargo).Error //si no hay un error, se guarda el articulo
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": cargo})
}

//BuscarCargo ...
func BuscarCargo(c *gin.Context) {

	var cargo cargo.CargoEmp

	token.ValidarToken()

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&cargo).Error; err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "No existe este cargo."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": cargo})
}

//ActualizarCargo ...
func ActualizarCargo(c *gin.Context) {

	var input CargoInput
	var carg cargo.CargoEmp

	token.ValidarToken()

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.SecureJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if input.Descripcion == "" {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción."})
		return
	}

	if err := database.DB.Where("cargo_emp_id=?", c.Param("id")).First(&carg).Error; err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Cargo no encontrada."})
		return
	}

	cargo := cargo.CargoEmp{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Model(&carg).Where("cargo_emp_id=?", c.Param("id")).Update(cargo).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": cargo})
}

//EliminarCargo ...
func EliminarCargo(c *gin.Context) {

	var cargo cargo.CargoEmp

	token.ValidarToken()
	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Where("cargo_emp_id=?", c.Param("id")).Delete(&cargo).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}
