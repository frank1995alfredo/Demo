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

/**************METODOS PARA DISCAPACIDAD*******************/

//ObtenerDiscapacidad ...
func ObtenerDiscapacidad(c *gin.Context) {
	var discapacidad []mantenimiento.Discapacidad

	token.ValidarToken()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	db := database.DB.Find(&discapacidad)

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"discapacidad_id asc"},
		ShowSQL: false,
	}, &discapacidad)
	c.SecureJSON(http.StatusOK, gin.H{"data": paginator})

}

//CrearDiscapacidad ...
func CrearDiscapacidad(c *gin.Context) {

	var input DiscapacidadInput
	var discap mantenimiento.Discapacidad

	token.ValidarToken()

	//validamos los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//valido que la descripcion no este vacia
	if input.ValidarEntrada() {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción."})
		return
	}

	//pregunto si esa discapacidad existe en la base de datos
	if err := database.DB.Where("descripcion=?", input.Descripcion).First(&discap).Error; err == nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Ya existe esta discapacidad, ingrese otra."})
		return
	}

	discapacidad := mantenimiento.Discapacidad{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Create(&discapacidad).Error //si no hay un error, se guarda el articulo
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": discapacidad})
}

//BuscarDiscapacidad ...
func BuscarDiscapacidad(c *gin.Context) {

	var discap mantenimiento.Discapacidad

	token.ValidarToken()

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&discap).Error; err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "No existe esta discapacidad."})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{"data": discap})
}

//ActualizarDiscapacidad ...
func ActualizarDiscapacidad(c *gin.Context) {

	var input DiscapacidadInput
	var disc mantenimiento.Discapacidad

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

	if err := database.DB.Where("discapacidad_id=?", c.Param("id")).First(&disc).Error; err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Discapacidad no encontrada."})
		return
	}

	discapacidad := mantenimiento.Discapacidad{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Model(&disc).Where("discapacidad_id=?", c.Param("id")).Update(discapacidad).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusOK, gin.H{"data": discapacidad})
}

//EliminarDiscapacidad ...
func EliminarDiscapacidad(c *gin.Context) {

	var discapacidad mantenimiento.Discapacidad

	token.ValidarToken()

	if err := database.DB.Where("discapacidad_id=?", c.Param("id")).First(&discapacidad).Error; err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Esta discapacidad no existe"})
		return
	}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Where("discapacidad_id=?", c.Param("id")).Delete(&discapacidad).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Resgistro eliminado."})

}
