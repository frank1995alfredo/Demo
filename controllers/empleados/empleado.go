package empleados

import (
	"net/http"
	"strconv"

	empleados "github.com/frank1995alfredo/api/class/empleado"
	input "github.com/frank1995alfredo/api/functions"
	token "github.com/frank1995alfredo/api/functions"
	"github.com/gin-gonic/gin"
)

var inp input.InputGlobal

/**************METODO PARA EMPLEADO******************/

//ObtenerEmpleado ...
func ObtenerEmpleado(c *gin.Context) {

	empleado := new(empleados.Empleado)

	token.ValidarToken()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	emp := empleado.ObtenerEmpleados(page, limit)

	c.SecureJSON(http.StatusOK, gin.H{"data": emp})

}

//CrearEmpleado ...
func CrearEmpleado(c *gin.Context) {

	empleado := new(empleados.Empleado)

	//validaops los inputs
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//valido que los campos obligatorios no esten vacios
	if inp.IngresarEmpleadoInput() {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Por favor, ingrese los campos que son obligatorios."})
		return
	}

	emp := empleado.CrearEmpleado(inp.DiscID, inp.CiuID, inp.PriNombre, inp.SegNombre, inp.PriApellido, inp.SegApellido,
		inp.FechaNac, inp.NumCedula, inp.Direccion, inp.Email, inp.Telefono, inp.Genero, true,
		inp.NivelDis, inp.CargoEmpID, inp.CodigoEmp, inp.Foto)

	c.SecureJSON(http.StatusOK, gin.H{"data": emp})

}

//BuscarEmpleado ...
func BuscarEmpleado(c *gin.Context) {

	empleado := new(empleados.Empleado)

	//token.ValidarToken()

	emp, err := empleado.BuscarEmpleado(c.Param("valor"))

	if err != "" {
		c.SecureJSON(http.StatusOK, gin.H{"error": err})
		return
	}

	c.SecureJSON(http.StatusOK, gin.H{"data": emp})

}

//ActualizarEmpleado ...
func ActualizarEmpleado(c *gin.Context) {

	token.ValidarToken()

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.SecureJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.SecureJSON(http.StatusCreated, gin.H{"data": "ggdf"})

}

//EliminarEmpleado ...
func EliminarEmpleado(c *gin.Context) {

	//token.ValidarToken()

	empleado := new(empleados.Empleado)

	emp := empleado.EliminarEmpleado(c.Param("id"))

	c.SecureJSON(http.StatusOK, gin.H{"data": emp})

}
