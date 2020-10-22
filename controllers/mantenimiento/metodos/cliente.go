package metodos

import (
	"net/http"

	inputsMantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento/inputsMantenimiento"
	database "github.com/frank1995alfredo/api/database"
	mantenimiento "github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/gin-gonic/gin"
)

/************MODELO DE CLIENTE********************/

//ObtenerCliente ...
func ObtenerCliente(c *gin.Context) {

	var cliente []mantenimiento.Cliente

	err := database.DB.Order("cliente_id").Find(&cliente).Error
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": 1})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"data": cliente})
	}
}

//CrearCliente ...
func CrearCliente(c *gin.Context) {

	var input inputsMantenimiento.ClienteInput
	var clien mantenimiento.Cliente

	//validams los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//valido que los camplos obligatorios no esten vacios
	if input.ValidarEntrada() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Por favor, ingrese los campos que son obligatorios."})
		return
	}

	//pregunto si el cliente existe en la base de datos
	if err := database.DB.Where("num_cedula=?", input.NumCedula).First(&clien).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe este cliente con el mismo número de cédula, ingrese otro."})
		return
	}

	cliente := mantenimiento.Cliente{DiscID: input.DiscID, CiuID: input.CiuID,
		PriNombre: input.PriNombre, SegNombre: input.SegNombre,
		PriApellido: input.PriApellido, SegApellido: input.SegApellido, FechaNac: input.FechaNac,
		NumCedula: input.NumCedula, CodigoCli: input.CodigoCli, Direccion: input.Direccion,
		Email: input.Email, Telefono: input.Telefono, Genero: input.Genero,
		NivelDis: input.NivelDis}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Create(&cliente).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": cliente})
}

//BuscarCliente ...
func BuscarCliente(c *gin.Context) {

	var cliente mantenimiento.Cliente

	if err := database.DB.Where("num_cedula=?", c.Param("numcedula")).First(&cliente).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe este cliente."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": cliente})
}

//ActualizarCliente ...
func ActualizarCliente(c *gin.Context) {

	var input inputsMantenimiento.ClienteInput
	var clien mantenimiento.Cliente

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	//valido que los camplos obligatorios no esten vacios
	if input.CiuID == 0 ||
		input.PriNombre == "" || input.SegNombre == "" ||
		input.PriApellido == "" || input.SegApellido == "" ||
		input.NumCedula == "" || input.CodigoCli == "" ||
		input.Direccion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Por favor, ingrese los campos que son obligatorios."})
		return
	}

	if err := database.DB.Where("cliente_id=?", c.Param("id")).First(&clien).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cliente no encontrado."})
		return
	}

	cliente := mantenimiento.Cliente{DiscID: input.DiscID, CiuID: input.CiuID,
		PriNombre: input.PriNombre, SegNombre: input.SegNombre,
		PriApellido: input.PriApellido, SegApellido: input.SegApellido, FechaNac: input.FechaNac,
		NumCedula: input.NumCedula, CodigoCli: input.CodigoCli, Direccion: input.Direccion,
		Email: input.Email, Telefono: input.Telefono, Genero: input.Genero,
		NivelDis: input.NivelDis}

	//inicio de la transaccions
	tx := database.DB.Begin()
	err := tx.Model(&clien).Where("cliente_id=?", c.Param("id")).Update(cliente).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": cliente})
}

//EliminarCliente ...
func EliminarCliente(c *gin.Context) {

	var cliente mantenimiento.Cliente

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Where("cliente_id=?", c.Param("id")).Delete(&cliente).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}
