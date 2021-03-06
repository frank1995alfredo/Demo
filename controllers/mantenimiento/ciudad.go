package mantenimiento

import (
	"net/http"

	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	_ "github.com/biezhi/gorm-paginator/pagination" //paginador

	database "github.com/frank1995alfredo/api/database"
	token "github.com/frank1995alfredo/api/functions"
	mantenimiento "github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/gin-gonic/gin"
)

/**************METODOS PARA CIUDAD*******************/

//ObtenerCiudad ...
func ObtenerCiudad(c *gin.Context) {

	var ciudad []mantenimiento.Ciudad

	token.ValidarToken()

	err2 := database.DB.Order("ciudad_id").Find(&ciudad).Error

	if err2 != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err2.Error})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))

	db := database.DB.Where("estado=?", true).Find(&ciudad)

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"ciudad_id asc"},
		ShowSQL: false,
	}, &ciudad)
	c.SecureJSON(http.StatusOK, gin.H{"data": paginator})

}

//CrearCiudad ... metodo para crear una provincia
func CrearCiudad(c *gin.Context) {
	//var ciudadId uint64
	var input CiudadInput
	var ciu mantenimiento.Ciudad

	token.ValidarToken()
	//validamos los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.VerificarDescripcion() {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Debe ingregar una descripción y una provincia."})
		return
	}

	if err := database.DB.Where("descripcion=?", input.Descripcion).First(&ciu).Error; err == nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Ya existe esta ciudad, ingrese otra."})
		return
	}

	ciudad := mantenimiento.Ciudad{ProID: input.ProID, Descripcion: input.Descripcion, Estado: true}

	tx := database.DB.Begin()
	err2 := tx.Create(&ciudad).Error //si no hay un error, se guarda el articulo
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusOK, gin.H{"data": ciudad})
}

//BuscarCiudad ...
func BuscarCiudad(c *gin.Context) {

	var ciu mantenimiento.Ciudad

	token.ValidarToken()

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&ciu).Error; err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "No existe esta provincia."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": ciu})
}

//ActualizarCiudad ...
func ActualizarCiudad(c *gin.Context) {

	var input CiudadInput
	var ciu mantenimiento.Ciudad

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.SecureJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	/*	_, err := token.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.SecureJSON(http.StatusUnauthorized, "No tiene permisos para acceder a esta opción.")
			return
		}*/

	ciudad := mantenimiento.Ciudad{ProID: input.ProID, Descripcion: input.Descripcion, Estado: input.Estado}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Model(&ciu).Where("ciudad_id=?", c.Param("id")).Update(ciudad).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro actualizado correctamente."})
}

//EliminarCiudad ...
func EliminarCiudad(c *gin.Context) {

	var ciudad mantenimiento.Ciudad

	token.ValidarToken()

	if err := database.DB.Where("ciudad_id=?", c.Param("id")).First(&ciudad).Error; err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Ciudad no existe."})
		return
	}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Model(&ciudad).Where("ciudad_id=?", c.Param("id")).Update("estado", false).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}
