package rutas

import (
	"github.com/frank1995alfredo/api/config"

	mantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento/metodos"
	practicas "github.com/frank1995alfredo/api/controllers/practicas"
	database "github.com/frank1995alfredo/api/database"
	token "github.com/frank1995alfredo/api/token"

	"github.com/gin-gonic/gin"
)

//Rutas ...
func Rutas() {
	r := gin.Default()
	r.Use(config.CORS)

	database.ConectorBD()
	defer database.DB.Close()

	r.POST("/login", token.Login)
	r.POST("/logout", token.TokenAuthMiddleware(), token.Logout)

	usuario := r.Group("/usuarios")
	{
		//usuario.GET("/obtenerUsuario", token.TokenAuthMiddleware(), mantenimiento.ObtenerProvincia)
		usuario.POST("/crearUsuario", token.TokenAuthMiddleware(), mantenimiento.RegistrarUsuario)
		usuario.PATCH("/activarUsuario/:id", token.TokenAuthMiddleware(), mantenimiento.ActivarUsuario)
		usuario.PATCH("/desactivarUsuario/:id", token.TokenAuthMiddleware(), mantenimiento.DesactivarUsuario)
	}

	provincia := r.Group("/provincias")
	{
		provincia.GET("/obtenerProvincia", token.TokenAuthMiddleware(), mantenimiento.ObtenerProvincia)
		provincia.POST("/crearProvincia", token.TokenAuthMiddleware(), mantenimiento.CrearProvincia)
		provincia.GET("/buscarProvincia/:descripcion", token.TokenAuthMiddleware(), mantenimiento.BuscarProvincia)
		provincia.PATCH("/actualizarProvincia/:id", token.TokenAuthMiddleware(), mantenimiento.ActualizarProvincia)
		provincia.DELETE("/eliminarProvincia/:id", token.TokenAuthMiddleware(), mantenimiento.EliminarProvincia)
	}

	ciudad := r.Group("/ciudades")
	{
		ciudad.GET("/obtenerCiudad", token.TokenAuthMiddleware(), mantenimiento.ObtenerCiudad)
		ciudad.POST("/crearCiudad", token.TokenAuthMiddleware(), mantenimiento.CrearCiudad)
		ciudad.GET("/buscarCiudad/:descripcion", token.TokenAuthMiddleware(), mantenimiento.BuscarCiudad)
		ciudad.PATCH("/actualizarCiudad/:id", token.TokenAuthMiddleware(), mantenimiento.ActualizarCiudad)
		ciudad.DELETE("/eliminarCiudad/:id", token.TokenAuthMiddleware(), mantenimiento.EliminarCiudad)
	}

	discapacidad := r.Group("/discapacidades")
	{
		discapacidad.GET("/obtenerDiscapacidad", token.TokenAuthMiddleware(), mantenimiento.ObtenerDiscapacidad)
		discapacidad.POST("/crearDiscapacidad", token.TokenAuthMiddleware(), mantenimiento.CrearDiscapacidad)
		discapacidad.GET("/buscarDiscapacidad/:descripcion", token.TokenAuthMiddleware(), mantenimiento.BuscarDiscapacidad)
		discapacidad.PATCH("/actualizarDiscapacidad/:id", token.TokenAuthMiddleware(), mantenimiento.ActualizarDiscapacidad)
		discapacidad.DELETE("/eliminarDiscapacidad/:id", token.TokenAuthMiddleware(), mantenimiento.EliminarDiscapacidad)
	}

	cargo := r.Group("/cargos")
	{
		cargo.GET("/obtenerCargo", token.TokenAuthMiddleware(), mantenimiento.ObtenerCargo)
		cargo.POST("/crearCargo", token.TokenAuthMiddleware(), mantenimiento.CrearCargo)
		cargo.GET("/buscarCargo/:descripcion", token.TokenAuthMiddleware(), mantenimiento.BuscarCargo)
		cargo.PATCH("/actualizarCargo/:id", token.TokenAuthMiddleware(), mantenimiento.ActualizarCargo)
		cargo.DELETE("/eliminarCargo/:id", token.TokenAuthMiddleware(), mantenimiento.EliminarCargo)
	}

	cliente := r.Group("/clientes")
	{
		cliente.GET("/obtenerCliente", token.TokenAuthMiddleware(), mantenimiento.ObtenerCliente)
		cliente.POST("/crearCliente", token.TokenAuthMiddleware(), mantenimiento.CrearCliente)
		cliente.GET("/buscarCliente/:numcedula", token.TokenAuthMiddleware(), mantenimiento.BuscarCliente)
		cliente.PATCH("/actualizarCliente/:id", token.TokenAuthMiddleware(), mantenimiento.ActualizarCliente)
		cliente.DELETE("/eliminarCliente/:id", token.TokenAuthMiddleware(), mantenimiento.EliminarCliente)
	}

	empleado := r.Group("/empleados")
	{
		empleado.GET("/obtenerEmpleado", token.TokenAuthMiddleware(), mantenimiento.ObtenerEmpleado)
		empleado.POST("/crearEmpleado", token.TokenAuthMiddleware(), mantenimiento.CrearEmpleado)
		empleado.GET("/buscarEmpleado/:numcedula", token.TokenAuthMiddleware(), mantenimiento.BuscarEmpleado)
		empleado.PATCH("/actualizarEmplado/:id", token.TokenAuthMiddleware(), mantenimiento.ActualizarEmpleado)
		empleado.DELETE("/eliminarEmpleado/:id", token.TokenAuthMiddleware(), mantenimiento.EliminarEmpleado)
	}

	persona := r.Group("/personas")
	{
		persona.GET("/obtenerCliente", practicas.ObtenerCliente)
		persona.POST("/crearPersona", practicas.CrearPersona)
		persona.GET("/buscarPersona/:apellido", practicas.BuscarPersona)
		persona.PATCH("/actualizarPersona/:id", practicas.ActualizarPersona)
		persona.DELETE("/eliminarPersona/:id", practicas.EliminarPersona)
	}

	r.Run()
}
