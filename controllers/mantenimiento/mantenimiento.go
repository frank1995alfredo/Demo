package mantenimiento

import (
	"net/http"

	inputsMantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento/inputsMantenimiento"
	database "github.com/frank1995alfredo/api/database"
	mantenimiento "github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/gin-gonic/gin"
)

/**************METODOS PARA PROVINCIA*******************/

//ObtenerProvincia ... metodo para obtener todas las provincias
func ObtenerProvincia(c *gin.Context) {
	var provincia []mantenimiento.Provincia

	err := database.DB.Order("provincia_id").Find(&provincia).Error
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"data": provincia})
	}
}

//CrearProvincia ... metodo para crear una provincia
func CrearProvincia(c *gin.Context) {

	var input inputsMantenimiento.ProvinciaInput
	var provinc mantenimiento.Provincia
	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Descripcion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción"})
		return
	}

	if err := database.DB.Where("descripcion=?", input.Descripcion).First(&provinc).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe esta provincia, ingrese otra."})
		return
	}

	provincia := mantenimiento.Provincia{Descripcion: input.Descripcion, Estado: input.Estado}

	tx := database.DB.Begin()
	err := tx.Create(&provincia).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusOK, gin.H{"data": provincia})
}

//BuscarProvincia ...
func BuscarProvincia(c *gin.Context) {

	var provinc mantenimiento.Provincia

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&provinc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe esta provincia."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": provinc})
}

//ActualizarProvincia ...
func ActualizarProvincia(c *gin.Context) {

	var input inputsMantenimiento.ProvinciaInput
	var provinc mantenimiento.Provincia

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
	err := tx.Model(&provinc).Where("provincia_id=?", c.Param("id")).Update(provin).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": provinc})
}

//EliminarProvincia ...
func EliminarProvincia(c *gin.Context) {

	var provincia mantenimiento.Provincia

	if err := database.DB.Where("provincia_id=?", c.Param("id")).First(&provincia).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provincia no existe"})
		return
	}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Where("provincia_id=?", c.Param("id")).Delete(&provincia).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}

/**************METODOS PARA CIUDAD*******************/

//ObtenerCiudad ...
func ObtenerCiudad(c *gin.Context) {
	var ciudad []mantenimiento.Ciudad

	err := database.DB.Order("ciudad_id").Find(&ciudad).Error
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"data": ciudad})
	}
}

//CrearCiudad ... metodo para crear una provincia
func CrearCiudad(c *gin.Context) {

	var input inputsMantenimiento.CiudadInput
	var ciu mantenimiento.Ciudad
	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Descripcion == "" || input.ProID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe ingregar una descripción y una provincia."})
		return
	}

	if err := database.DB.Where("descripcion=?", input.Descripcion).First(&ciu).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe esta ciudad, ingrese otra."})
		return
	}

	ciudad := mantenimiento.Ciudad{ProID: input.ProID, Descripcion: input.Descripcion, Estado: input.Estado}

	tx := database.DB.Begin()
	err := tx.Create(&ciudad).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusOK, gin.H{"data": ciudad})
}

//BuscarCiudad ...
func BuscarCiudad(c *gin.Context) {

	var ciu mantenimiento.Ciudad

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&ciu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe esta provincia."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": ciu})
}

//ActualizarCiudad ...
func ActualizarCiudad(c *gin.Context) {

	var input inputsMantenimiento.CiudadInput
	var ciu mantenimiento.Ciudad

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("ciudad_id=?", c.Param("id")).First(&ciu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provincia no encontrada."})
		return
	}

	ciudad := mantenimiento.Ciudad{ProID: input.ProID, Descripcion: input.Descripcion, Estado: input.Estado}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&ciu).Where("ciudad_id=?", c.Param("id")).Update(ciudad).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": ciudad})
}

//EliminarCiudad ...
func EliminarCiudad(c *gin.Context) {

	var ciudad mantenimiento.Ciudad

	if err := database.DB.Where("ciudad_id=?", c.Param("id")).First(&ciudad).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ciudad no existe."})
		return
	}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Where("ciudad_id=?", c.Param("id")).Delete(&ciudad).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}

/**************METODOS PARA DISCAPACIDAD*******************/

//ObtenerDiscapadcidad ...
func ObtenerDiscapadcidad(c *gin.Context) {
	var discapacidad []mantenimiento.Discapacidad

	err := database.DB.Order("discapacidad_id").Find(&discapacidad).Error
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"data": discapacidad})
	}
}

//CrearDiscapacidad ...
func CrearDiscapacidad(c *gin.Context) {

	var input inputsMantenimiento.DiscapacidadInput
	var discap mantenimiento.Discapacidad

	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//valido que la descripcion no este vacia
	if input.Descripcion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción."})
		return
	}

	//pregunto si esa discapacidad existe en la base de datos
	if err := database.DB.Where("descripcion=?", input.Descripcion).First(&discap).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe esta discapacidad, ingrese otra."})
		return
	}

	discapacidad := mantenimiento.Discapacidad{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Create(&discapacidad).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": discapacidad})
}

//BuscarDiscapacidad ...
func BuscarDiscapacidad(c *gin.Context) {

	var discap mantenimiento.Discapacidad

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&discap).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe esta discapacidad."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": discap})
}

//ActualizarDiscapacidad ...
func ActualizarDiscapacidad(c *gin.Context) {

	var input inputsMantenimiento.DiscapacidadInput
	var disc mantenimiento.Discapacidad

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if input.Descripcion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción."})
		return
	}

	if err := database.DB.Where("discapacidad_id=?", c.Param("id")).First(&disc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Discapacidad no encontrada."})
		return
	}

	discapacidad := mantenimiento.Discapacidad{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&disc).Where("discapacidad_id=?", c.Param("id")).Update(discapacidad).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": discapacidad})
}

//EliminarDiscapacidad ...
func EliminarDiscapacidad(c *gin.Context) {

	var discapacidad mantenimiento.Discapacidad

	if err := database.DB.Where("discapacidad_id=?", c.Param("id")).First(&discapacidad).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Esta discapacidad no existe"})
		return
	}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Where("discapacidad_id=?", c.Param("id")).Delete(&discapacidad).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Resgistro eliminado."})

}

/***************METODO PARA CARGO EMPLEADO********************/

//ObtenerCargo ...
func ObtenerCargo(c *gin.Context) {
	var cargo []mantenimiento.CargoEmp

	err := database.DB.Order("cargo_emp_id").Find(&cargo).Error
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"data": cargo})
	}
}

//CrearCargo ...
func CrearCargo(c *gin.Context) {

	var input inputsMantenimiento.CargoInput
	var carg mantenimiento.CargoEmp

	//validaops los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//valido que la descripcion no este vacia
	if input.Descripcion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción."})
		return
	}

	//pregunto si este cargo existe en la base de datos
	if err := database.DB.Where("descripcion=?", input.Descripcion).First(&carg).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe este cargo, ingrese otro."})
		return
	}

	cargo := mantenimiento.CargoEmp{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Create(&cargo).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": cargo})
}

//BuscarCargo ...
func BuscarCargo(c *gin.Context) {

	var cargo mantenimiento.CargoEmp

	if err := database.DB.Where("descripcion=?", c.Param("descripcion")).First(&cargo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe este cargo."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": cargo})
}

//ActualizarCargo ...
func ActualizarCargo(c *gin.Context) {

	var input inputsMantenimiento.CargoInput
	var carg mantenimiento.CargoEmp

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if input.Descripcion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción."})
		return
	}

	if err := database.DB.Where("cargo_emp_id=?", c.Param("id")).First(&carg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cargo no encontrada."})
		return
	}

	cargo := mantenimiento.CargoEmp{Descripcion: input.Descripcion}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&carg).Where("cargo_emp_id=?", c.Param("id")).Update(cargo).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": cargo})
}

//EliminarCargo ...
func EliminarCargo(c *gin.Context) {

	var cargo mantenimiento.CargoEmp

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Where("cargo_emp_id=?", c.Param("id")).Delete(&cargo).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}

/**************METODO PARA EMPLEADO******************/

//ObtenerEmpleado ...
func ObtenerEmpleado(c *gin.Context) {

	var empleado []mantenimiento.CargoEmp

	err := database.DB.Order("empleado_id").Find(&empleado).Error
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"data": empleado})
	}
}

//CrearEmpleado ...
func CrearEmpleado(c *gin.Context) {

	var input inputsMantenimiento.EmpleadoInput
	var emp mantenimiento.Empleado

	//validams los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	var input inputsMantenimiento.EmpleadoInput
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

/************MODELO DE CLIENTE********************/

//ObtenerCliente ...
func ObtenerCliente(c *gin.Context) {

	var cliente []mantenimiento.Cliente

	err := database.DB.Order("cliente_id").Find(&cliente).Error
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
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
	if input.CiuID == 0 || input.DiscID == 0 ||
		input.PriNombre == "" || input.SegNombre == "" ||
		input.PriApellido == "" || input.SegApellido == "" ||
		input.NumCedula == "" || input.CodigoCli == "" ||
		input.Direccion == "" {
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

	//inicio de la transaccion
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
