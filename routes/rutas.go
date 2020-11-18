package rutas

import (
	"github.com/frank1995alfredo/api/config"

	clientes "github.com/frank1995alfredo/api/controllers/clientes"
	empleados "github.com/frank1995alfredo/api/controllers/empleados"
	mantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento"
	token "github.com/frank1995alfredo/api/controllers/token"
	usuarios "github.com/frank1995alfredo/api/controllers/usuarios"
	database "github.com/frank1995alfredo/api/database"

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
		usuario.POST("/crearUsuario", token.TokenAuthMiddleware(), usuarios.RegistrarUsuario)
		usuario.PATCH("/activarUsuario/:id", token.TokenAuthMiddleware(), usuarios.ActivarUsuario)
		usuario.PATCH("/desactivarUsuario/:id", token.TokenAuthMiddleware(), usuarios.DesactivarUsuario)
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
		cargo.GET("/obtenerCargo", token.TokenAuthMiddleware(), empleados.ObtenerCargo)
		cargo.POST("/crearCargo", token.TokenAuthMiddleware(), empleados.CrearCargo)
		cargo.GET("/buscarCargo/:descripcion", token.TokenAuthMiddleware(), empleados.BuscarCargo)
		cargo.PATCH("/actualizarCargo/:id", token.TokenAuthMiddleware(), empleados.ActualizarCargo)
		cargo.DELETE("/eliminarCargo/:id", token.TokenAuthMiddleware(), empleados.EliminarCargo)
	}

	cliente := r.Group("/clientes")
	{
		cliente.GET("/obtenerCliente", token.TokenAuthMiddleware(), clientes.ObtenerCliente)
		cliente.POST("/crearCliente", token.TokenAuthMiddleware(), clientes.CrearCliente)
		cliente.GET("/buscarCliente/:numcedula", token.TokenAuthMiddleware(), clientes.BuscarCliente)
		cliente.PATCH("/actualizarCliente/:id", token.TokenAuthMiddleware(), clientes.ActualizarCliente)
		cliente.DELETE("/eliminarCliente/:id", token.TokenAuthMiddleware(), clientes.EliminarCliente)
	}

	empleado := r.Group("/empleados")
	{
		empleado.GET("/obtenerEmpleado", token.TokenAuthMiddleware(), empleados.ObtenerEmpleado)
		empleado.POST("/crearEmpleado", token.TokenAuthMiddleware(), empleados.CrearEmpleado)
		empleado.GET("/buscarEmpleado/:numcedula", token.TokenAuthMiddleware(), empleados.BuscarEmpleado)
		empleado.PATCH("/actualizarEmplado/:id", token.TokenAuthMiddleware(), empleados.ActualizarEmpleado)
		empleado.DELETE("/eliminarEmpleado/:id", token.TokenAuthMiddleware(), empleados.EliminarEmpleado)
	}

	r.Run()
}
