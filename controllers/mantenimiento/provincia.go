package mantenimiento

import (
	"net/http"
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	token "github.com/frank1995alfredo/api/controllers/usuarios"
	database "github.com/frank1995alfredo/api/database"
	mantenimiento "github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/gin-gonic/gin"
)

/**************METODOS PARA PROVINCIA*******************/

//ObtenerProvincia ... metodo para obtener todas las provincias
func ObtenerProvincia(c *gin.Context) {

	var provincia []mantenimiento.Provincia

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

	err2 := database.DB.Order("provincia_id").Find(&provincia).Error
	if err2 != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

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

	var input ProvinciaInput
	var provinc mantenimiento.Provincia

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}
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

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&provinc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe esta provincia."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": provinc})
}

//ActualizarProvincia ...
func ActualizarProvincia(c *gin.Context) {

	var input ProvinciaInput
	var provinc mantenimiento.Provincia

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

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
	err2 := tx.Model(&provinc).Where("provincia_id=?", c.Param("id")).Update(provin).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusCreated, gin.H{"data": provinc})
}

//EliminarProvincia ...
func EliminarProvincia(c *gin.Context) {

	var provincia mantenimiento.Provincia

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

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
