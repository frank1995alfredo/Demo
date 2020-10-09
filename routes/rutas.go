package rutas

import (
	mantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento"

	database "github.com/frank1995alfredo/api/database"

	config "github.com/frank1995alfredo/api/config"
	"github.com/gin-gonic/gin"
)

//Rutas ...
func Rutas() {
	r := gin.Default()
	r.Use(config.CORS)

	database.ConectorBD()
	defer database.DB.Close()

	provincia := r.Group("/provincias")
	{
		provincia.GET("/obtenerProvincia", mantenimiento.ObtenerProvincia)
		provincia.POST("/crearProvincia", mantenimiento.CrearProvincia)
		provincia.GET("/buscarProvincia/:descripcion", mantenimiento.BuscarProvincia)
		provincia.PATCH("/actualizarProvincia/:id", mantenimiento.ActualizarProvincia)
		provincia.DELETE("/eliminarProvincia/:id", mantenimiento.EliminarProvincia)
	}

	ciudad := r.Group("/ciudades")
	{
		ciudad.GET("/obtenerCiudad", mantenimiento.ObtenerCiudad)
		ciudad.POST("/crearCiudad", mantenimiento.CrearCiudad)
		ciudad.GET("/buscarCiudad/:descripcion", mantenimiento.BuscarCiudad)
		ciudad.PATCH("/actualizarCiudad/:id", mantenimiento.ActualizarCiudad)
		ciudad.DELETE("/eliminarCiudad/:id", mantenimiento.EliminarCiudad)
	}

	discapacidad := r.Group("/discapacidades")
	{
		discapacidad.GET("/obtenerDiscapacidad", mantenimiento.ObtenerDiscapadcidad)
		discapacidad.POST("/crearDiscapacidad", mantenimiento.CrearDiscapacidad)
		discapacidad.GET("/buscarDiscapacidad/:descripcion", mantenimiento.BuscarDiscapacidad)
		discapacidad.PATCH("/actualizarDiscapacidad/:id", mantenimiento.ActualizarDiscapacidad)
		discapacidad.DELETE("/eliminarDiscapacidad/:id", mantenimiento.EliminarDiscapacidad)
	}

	cargo := r.Group("/cargos")
	{
		cargo.GET("/obtenerCargo", mantenimiento.ObtenerCargo)
		cargo.POST("/crearCargo", mantenimiento.CrearCargo)
		cargo.GET("/buscarCargo/:descripcion", mantenimiento.BuscarCargo)
		cargo.PATCH("/actualizarCargo/:id", mantenimiento.ActualizarCargo)
		cargo.DELETE("/eliminarCargo/:id", mantenimiento.EliminarCargo)
	}

	cliente := r.Group("/clientes")
	{
		cliente.GET("/obtenerCliente", mantenimiento.ObtenerCliente)
		cliente.POST("/crearCliente", mantenimiento.CrearCliente)
		cliente.GET("/buscarCliente/:numcedula", mantenimiento.BuscarCliente)
		cliente.PATCH("/actualizarCliente/:id", mantenimiento.ActualizarCliente)
		cliente.DELETE("/eliminarCliente/:id", mantenimiento.EliminarCliente)
	}

	empleado := r.Group("/empleados")
	{
		empleado.GET("/obtenerEmpleado", mantenimiento.ObtenerEmpleado)
		empleado.POST("/crearEmpleado", mantenimiento.CrearEmpleado)
		empleado.GET("/buscarEmpleado/:numcedula", mantenimiento.BuscarEmpleado)
		empleado.PATCH("/actualizarEmplado/:id", mantenimiento.ActualizarEmpleado)
		empleado.DELETE("/eliminarEmpleado/:id", mantenimiento.EliminarEmpleado)
	}

	r.Run()
}
