package mantenimiento

import (
	"net/http"
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	database "github.com/frank1995alfredo/api/database"
	token "github.com/frank1995alfredo/api/functions"
	mantenimiento "github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/gin-gonic/gin"
)

/**************METODOS PARA PROVINCIA*******************/

// ValidarEntradaInput ... VALIDAMOS LOS JSON
func ValidarEntradaInput() {
	c := gin.Context{}
	var inp token.InputGlobal

	if err := c.ShouldBindJSON(&inp); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

//ObtenerProvincia ... metodo para obtener todas las provincias
func ObtenerProvincia(c *gin.Context) {

	var provincia []mantenimiento.Provincia

	token.ValidarToken()

	err2 := database.DB.Order("provincia_id").Find(&provincia).Error
	if err2 != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err2.Error})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))

	db := database.DB.Where("estado=?", true).Find(&provincia)

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"provincia_id asc"},
		ShowSQL: false,
	}, &provincia)
	c.SecureJSON(http.StatusOK, gin.H{"data": paginator})

}

//CrearProvincia ... metodo para crear una provincia
func CrearProvincia(c *gin.Context) {

	var inp token.InputGlobal
	var provinc mantenimiento.Provincia

	//token.ValidarToken()

	if err := c.ShouldBindJSON(&inp); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("descripcion=?", inp.Descripcion).First(&provinc).Error; err == nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Ya existe esta provincia, ingrese otra."})
		return
	}

	provincia := mantenimiento.Provincia{Descripcion: inp.Descripcion, Estado: inp.Estado}

	tx := database.DB.Begin()
	err2 := tx.Create(&provincia).Error //si no hay un error, se guarda el articulo
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusOK, gin.H{"data": provincia})
}

//BuscarProvincia ...
func BuscarProvincia(c *gin.Context) {

	var provinc mantenimiento.Provincia

	token.ValidarToken()

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&provinc).Error; err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "No existe esta provincia."})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{"data": provinc})
}

//ActualizarProvincia ...
func ActualizarProvincia(c *gin.Context) {

	var input ProvinciaInput
	var provinc mantenimiento.Provincia

	token.ValidarToken()

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.SecureJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("provincia_id=?", c.Param("id")).First(&provinc).Error; err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Provincia no encontrada."})
		return
	}

	provin := mantenimiento.Provincia{Descripcion: input.Descripcion, Estado: input.Estado}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Model(&provinc).Where("provincia_id=?", c.Param("id")).Update(provin).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": provinc})
}

//EliminarProvincia ...
func EliminarProvincia(c *gin.Context) {

	var provincia mantenimiento.Provincia

	token.ValidarToken()

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Model(&provincia).Where("provincia_id=?", c.Param("id")).Update("estado", false).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}
