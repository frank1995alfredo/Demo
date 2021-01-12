package empleados

import (
	"net/http"
	"strconv"

	empleados "github.com/frank1995alfredo/api/class/empleado"
	input "github.com/frank1995alfredo/api/functions"
	"github.com/gin-gonic/gin"
)

/**************METODO PARA EMPLEADO******************/

//ObtenerEmpleado ...
func ObtenerEmpleado(c *gin.Context) {

	empleado := new(empleados.Empleado)

	//token.ValidarToken()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	emp := empleado.ObtenerEmpleados(page, limit)

	c.SecureJSON(http.StatusOK, gin.H{"data": emp})

}

//CrearEmpleado ...
func CrearEmpleado(c *gin.Context) {

	var Inp input.InputGlobal

	empleado := new(empleados.Empleado)

	//validaops los inputs
	if err := c.ShouldBindJSON(&Inp); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//token.ValidarToken()

	emp := empleado.CrearEmpleado(Inp.DiscID, Inp.CiuID, Inp.PriNombre, Inp.SegNombre, Inp.PriApellido,
		Inp.SegApellido, Inp.FechaNac, Inp.NumCedula, Inp.Direccion, Inp.Email, Inp.Telefono,
		Inp.Genero, true, Inp.NivelDis, Inp.CargoEmpID, Inp.CodigoEmp, Inp.Foto)

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

	var Inp input.InputGlobal

	empleado := new(empleados.Empleado)
	//token.ValidarToken()

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&Inp); err != nil {
		c.SecureJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	emp := empleado.ActualizarEmpleado(c.Param("id"), Inp.DiscID, Inp.CiuID, Inp.PriNombre, Inp.SegNombre,
		Inp.PriApellido, Inp.SegApellido, Inp.FechaNac, Inp.NumCedula, Inp.Direccion, Inp.Email,
		Inp.Telefono, Inp.Genero, Inp.Estado, Inp.NivelDis, Inp.CargoEmpID, Inp.CodigoEmp, Inp.Foto)

	c.SecureJSON(http.StatusOK, gin.H{"data": emp})

}

//EliminarEmpleado ...
func EliminarEmpleado(c *gin.Context) {

	//token.ValidarToken()

	empleado := new(empleados.Empleado)

	emp := empleado.EliminarEmpleado(c.Param("id"))

	c.SecureJSON(http.StatusOK, gin.H{"data": emp})

}
