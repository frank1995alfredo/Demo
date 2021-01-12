package clientes

import (
	"net/http"
	"strconv"

	clientes "github.com/frank1995alfredo/api/class/cliente"
	input "github.com/frank1995alfredo/api/functions"
	"github.com/gin-gonic/gin"
)

/************MODELO DE CLIENTE********************/

//ObtenerCliente ...
func ObtenerCliente(c *gin.Context) {

	cliente := new(clientes.Cliente)

	//token.ValidarToken()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	cli := cliente.ObtenerClientes(page, limit)

	c.SecureJSON(http.StatusOK, gin.H{"data": cli})

}

//CrearCliente ...
func CrearCliente(c *gin.Context) {

	var Inp input.InputGlobal

	cliente := new(clientes.Cliente)

	//token.ValidarToken()

	//validamos los inputs
	if err := c.ShouldBindJSON(&Inp); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cli := cliente.CrearCliente(Inp.DiscID, Inp.CiuID, Inp.PriNombre, Inp.SegNombre,
		Inp.PriApellido, Inp.SegApellido, Inp.FechaNac, Inp.NumCedula, Inp.Direccion, Inp.Email,
		Inp.Telefono, Inp.Genero, Inp.Estado, Inp.NivelDis, Inp.CargoEmpID, Inp.CodigoCli)

	c.SecureJSON(http.StatusOK, gin.H{"data": cli})
}

//BuscarCliente ...
func BuscarCliente(c *gin.Context) {

	cliente := new(clientes.Cliente)

	//token.ValidarToken()

	cli, err := cliente.BuscarCliente(c.Param("valor"))

	if err != "" {
		c.SecureJSON(http.StatusOK, gin.H{"error": err})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{"data": cli})
}

//ActualizarCliente ...
func ActualizarCliente(c *gin.Context) {

	var Inp input.InputGlobal

	cliente := new(clientes.Cliente)

	//token.ValidarToken()

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&Inp); err != nil {
		c.SecureJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	clien := cliente.ActualizarCliente(c.Param("id"), Inp.DiscID, Inp.CiuID, Inp.PriNombre,
		Inp.SegNombre, Inp.PriApellido, Inp.SegApellido, Inp.FechaNac, Inp.NumCedula,
		Inp.Direccion, Inp.Email, Inp.Telefono, Inp.Genero, Inp.Estado, Inp.NivelDis,
		Inp.CodigoCli)

	c.SecureJSON(http.StatusOK, gin.H{"data": clien})
}

//EliminarCliente ...
func EliminarCliente(c *gin.Context) {

	//token.ValidarToken()

	cliente := new(clientes.Cliente)

	cli := cliente.EliminarCliente(c.Param("id"))

	c.SecureJSON(http.StatusOK, gin.H{"data": cli})

}
