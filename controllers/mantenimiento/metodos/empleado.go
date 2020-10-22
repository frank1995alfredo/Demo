package metodos

import (
	"net/http"

	inputsmantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento/inputsmantenimiento"
	database "github.com/frank1995alfredo/api/database"
	mantenimiento "github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/gin-gonic/gin"
)

/**************METODO PARA EMPLEADO******************/

//ObtenerEmpleado ...
func ObtenerEmpleado(c *gin.Context) {

	var empleado []mantenimiento.Empleado

	err := database.DB.Order("empleado_id").Find(&empleado).Error
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"data": empleado})
	}
}

//CrearEmpleado ...
func CrearEmpleado(c *gin.Context) {

	var input inputsmantenimiento.EmpleadoInput
	var emp mantenimiento.Empleado

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
	if err := database.DB.Where("num_cedula=?", input.NumCedula).First(&emp).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe este empleado con el mismo número de cédula, ingrese otro."})
		return
	}

	empleado := mantenimiento.Empleado{DiscID: input.DiscID, CiuID: input.CiuID,
		CargoEmpID: input.CargoEmpID, PriNombre: input.PriNombre, SegNombre: input.SegNombre,
		PriApellido: input.PriApellido, SegApellido: input.SegApellido, FechNac: input.FechNac,
		NumCedula: input.NumCedula, CodigoEmp: input.CodigoEmp, Direccion: input.Direccion,
		Email: input.Email, Telefono: input.Telefono, Genero: input.Genero, Estado: input.Estado,
		Foto: input.Foto, NivelDis: input.NivelDis}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Create(&empleado).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": empleado})
}

//BuscarEmpleado ...
func BuscarEmpleado(c *gin.Context) {

	var empleado mantenimiento.Empleado

	if err := database.DB.Where("num_cedula=?", c.Param("numcedula")).First(&empleado).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe este empleado."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": empleado})
}

//ActualizarEmpleado ...
func ActualizarEmpleado(c *gin.Context) {

	var input inputsmantenimiento.EmpleadoInput
	var emp mantenimiento.Empleado

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	//valido que los camplos obligatorios no esten vacios
	if input.CiuID == 0 || input.DiscID == 0 ||
		input.CargoEmpID == 0 || input.PriNombre == "" ||
		input.SegNombre == "" || input.PriApellido == "" ||
		input.SegApellido == "" || input.NumCedula == "" ||
		input.CodigoEmp == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Por favor, ingrese los campos que son obligatorios."})
		return
	}

	if err := database.DB.Where("empleado_id=?", c.Param("id")).First(&emp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empleado no encontrada."})
		return
	}

	empleado := mantenimiento.Empleado{DiscID: input.DiscID, CiuID: input.CiuID,
		CargoEmpID: input.CargoEmpID, PriNombre: input.PriNombre, SegNombre: input.SegNombre,
		PriApellido: input.PriApellido, SegApellido: input.SegApellido, FechNac: input.FechNac,
		NumCedula: input.NumCedula, CodigoEmp: input.CodigoEmp, Direccion: input.Direccion,
		Email: input.Email, Telefono: input.Telefono, Genero: input.Genero, Estado: input.Estado,
		Foto: input.Foto, NivelDis: input.NivelDis}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&emp).Where("cargo_emp_id=?", c.Param("id")).Update(empleado).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": empleado})
}

//EliminarEmpleado ...
func EliminarEmpleado(c *gin.Context) {

	var empleado mantenimiento.CargoEmp

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Where("empleado_id=?", c.Param("id")).Delete(&empleado).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}
