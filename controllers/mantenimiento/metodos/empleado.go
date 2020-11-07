package metodos

import (
	"net/http"
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	inputsmantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento/inputsMantenimiento"
	database "github.com/frank1995alfredo/api/database"
	mantenimiento "github.com/frank1995alfredo/api/models/mantenimiento"
	token "github.com/frank1995alfredo/api/token"
	"github.com/gin-gonic/gin"
)

/**************METODO PARA EMPLEADO******************/

//ObtenerEmpleado ...
func ObtenerEmpleado(c *gin.Context) {

	var empleado []mantenimiento.Empleado

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

	err2 := database.DB.Find(&empleado).Error
	if err2 != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "3"))

	db := database.DB.Where("estado=?", true).Find(&empleado)

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"empleado_id asc"},
		ShowSQL: false,
	}, &empleado)
	c.SecureJSON(http.StatusOK, gin.H{"data": paginator})

}

//CrearEmpleado ...
func CrearEmpleado(c *gin.Context) {

	var input inputsmantenimiento.EmpleadoInput
	var emp mantenimiento.Empleado

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

	//validams los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//valido que los campos obligatorios no esten vacios
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
	err2 := tx.Create(&empleado).Error //si no hay un error, se guarda el articulo
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": empleado})
}

//BuscarEmpleado ...
func BuscarEmpleado(c *gin.Context) {

	var empleado mantenimiento.Empleado

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

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
	err2 := tx.Model(&emp).Where("empleado_id=?", c.Param("id")).Update(empleado).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": empleado})
}

//EliminarEmpleado ...
func EliminarEmpleado(c *gin.Context) {

	var emp mantenimiento.Empleado

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

	tx := database.DB.Begin()
	err2 := tx.Model(&emp).Where("empleado_id=?", c.Param("id")).Update("estado", false).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}
