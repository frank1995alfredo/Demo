package clientes

import (
	"net/http"
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	database "github.com/frank1995alfredo/api/database"
	token "github.com/frank1995alfredo/api/functions"
	cliente "github.com/frank1995alfredo/api/models/clientes"
	"github.com/gin-gonic/gin"
)

/************MODELO DE CLIENTE********************/

//ObtenerCliente ...
func ObtenerCliente(c *gin.Context) {

	var cliente []cliente.Cliente

	token.ValidarToken()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	db := database.DB.Find(&cliente)

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"cliente_id asc"},
		ShowSQL: false,
	}, &cliente)
	c.SecureJSON(http.StatusOK, gin.H{"data": paginator})

}

//CrearCliente ...
func CrearCliente(c *gin.Context) {

	var input ClienteInput
	var clien cliente.Cliente

	token.ValidarToken()

	//validamos los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//valido que los camplos obligatorios no esten vacios
	if input.ValidarEntrada() {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Por favor, ingrese los campos que son obligatorios."})
		return
	}

	//pregunto si el cliente existe en la base de datos
	if err := database.DB.Where("num_cedula=?", input.NumCedula).First(&clien).Error; err == nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Ya existe este cliente con el mismo número de cédula, ingrese otro."})
		return
	}

	cliente := cliente.Cliente{DiscID: input.DiscID, CiuID: input.CiuID,
		PriNombre: input.PriNombre, SegNombre: input.SegNombre,
		PriApellido: input.PriApellido, SegApellido: input.SegApellido, FechaNac: input.FechaNac,
		NumCedula: input.NumCedula, CodigoCli: input.CodigoCli, Direccion: input.Direccion,
		Email: input.Email, Telefono: input.Telefono, Genero: input.Genero, Estado: input.Estado,
		NivelDis: input.NivelDis}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Create(&cliente).Error //si no hay un error, se guarda el cliente
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": cliente})
}

//BuscarCliente ...
func BuscarCliente(c *gin.Context) {

	var cliente cliente.Cliente

	token.ValidarToken()

	if err := database.DB.Where("num_cedula=?", c.Param("numcedula")).First(&cliente).Error; err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "No existe este cliente."})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{"data": cliente})
}

//ActualizarCliente ...
func ActualizarCliente(c *gin.Context) {

	var input ClienteInput
	var clien cliente.Cliente

	token.ValidarToken()

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.SecureJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	//valido que los campos obligatorios no esten vacios
	if input.ValidarEntrada() {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Por favor, ingrese los campos que son obligatorios."})
		return
	}

	if err := database.DB.Where("cliente_id=?", c.Param("id")).First(&clien).Error; err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Cliente no encontrado."})
		return
	}

	cliente := cliente.Cliente{DiscID: input.DiscID, CiuID: input.CiuID,
		PriNombre: input.PriNombre, SegNombre: input.SegNombre,
		PriApellido: input.PriApellido, SegApellido: input.SegApellido, FechaNac: input.FechaNac,
		NumCedula: input.NumCedula, CodigoCli: input.CodigoCli, Direccion: input.Direccion,
		Email: input.Email, Telefono: input.Telefono, Genero: input.Genero, Estado: input.Estado,
		NivelDis: input.NivelDis}

	//inicio de la transaccions
	tx := database.DB.Begin()
	err2 := tx.Model(&clien).Where("cliente_id=?", c.Param("id")).Update(cliente).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusOK, gin.H{"data": cliente})
}

//EliminarCliente ...
func EliminarCliente(c *gin.Context) {

	var cliente cliente.Cliente

	token.ValidarToken()

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Where("cliente_id=?", c.Param("id")).Delete(&cliente).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}
