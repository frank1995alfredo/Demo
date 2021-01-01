package rutas

import (
	"github.com/frank1995alfredo/api/config"

	clientes "github.com/frank1995alfredo/api/controllers/clientes"
	empleados "github.com/frank1995alfredo/api/controllers/empleados"
	mantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento"
	token "github.com/frank1995alfredo/api/controllers/usuarios"
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

	usuario := r.Group("/security")
	{
		//usuario.GET("/obtenerUsuario", token.TokenAuthMiddleware(), mantenimiento.ObtenerProvincia)
		usuario.POST("/user", token.TokenAuthMiddleware(), usuarios.RegistrarUsuario)
		usuario.PATCH("/user/:id", token.TokenAuthMiddleware(), usuarios.ActivarUsuario)
		usuario.PATCH("/iuser/:id", token.TokenAuthMiddleware(), usuarios.DesactivarUsuario)
	}
	rutes := r.Group("/api/index.php/model")
	{

		//PROVINCIA
		rutes.GET("/provinces", token.TokenAuthMiddleware(), mantenimiento.ObtenerProvincia)
		rutes.POST("/sendprovinces" /*token.TokenAuthMiddleware(),*/, mantenimiento.CrearProvincia)
		rutes.GET("/getprovinces/:descripcion", token.TokenAuthMiddleware(), mantenimiento.BuscarProvincia)
		rutes.PATCH("/patchprovinces/:id", token.TokenAuthMiddleware(), mantenimiento.ActualizarProvincia)
		rutes.DELETE("/delprovinces/:id", token.TokenAuthMiddleware(), mantenimiento.EliminarProvincia)

		//CIUDAD
		rutes.GET("/cities", mantenimiento.ObtenerCiudad)
		rutes.POST("/sendcities", mantenimiento.CrearCiudad)
		rutes.GET("/getcities/:descripcion", token.TokenAuthMiddleware(), mantenimiento.BuscarCiudad)
		rutes.PATCH("/patchcities/:id", mantenimiento.ActualizarCiudad)
		rutes.DELETE("/delcities/:id", mantenimiento.EliminarCiudad)

		//DISCAPACIDAD
		rutes.GET("/disabilities", token.TokenAuthMiddleware(), mantenimiento.ObtenerDiscapacidad)
		rutes.POST("/senddisabilities", token.TokenAuthMiddleware(), mantenimiento.CrearDiscapacidad)
		rutes.GET("/getdisabilities/:descripcion", token.TokenAuthMiddleware(), mantenimiento.BuscarDiscapacidad)
		rutes.PATCH("/patchdisabilities/:id", token.TokenAuthMiddleware(), mantenimiento.ActualizarDiscapacidad)
		rutes.DELETE("/deldisabilities/:id", token.TokenAuthMiddleware(), mantenimiento.EliminarDiscapacidad)

		//CARGO
		rutes.GET("/positions" /*token.TokenAuthMiddleware(),*/, empleados.ObtenerCargo)
		rutes.POST("/sendpositions", token.TokenAuthMiddleware(), empleados.CrearCargo)
		rutes.GET("/getpositions/:descripcion", token.TokenAuthMiddleware(), empleados.BuscarCargo)
		rutes.PATCH("/patchpositions/:id", token.TokenAuthMiddleware(), empleados.ActualizarCargo)
		rutes.DELETE("/delpositions/:id", token.TokenAuthMiddleware(), empleados.EliminarCargo)

		//CLIENTE
		rutes.GET("/clients", token.TokenAuthMiddleware(), clientes.ObtenerCliente)
		rutes.POST("/sendclients", token.TokenAuthMiddleware(), clientes.CrearCliente)
		rutes.GET("/getclients/:numcedula", token.TokenAuthMiddleware(), clientes.BuscarCliente)
		rutes.PATCH("/patchclients/:id", token.TokenAuthMiddleware(), clientes.ActualizarCliente)
		rutes.DELETE("/delclients/:id", token.TokenAuthMiddleware(), clientes.EliminarCliente)

		//EMPLEADO
		rutes.GET("/employees" /*token.TokenAuthMiddleware(),*/, empleados.ObtenerEmpleado)
		rutes.POST("/sendemployees" /*token.TokenAuthMiddleware(),*/, empleados.CrearEmpleado)
		rutes.GET("/getemployees/:valor" /*token.TokenAuthMiddleware(),*/, empleados.BuscarEmpleado)
		rutes.PATCH("/patchemployees/:id" /*token.TokenAuthMiddleware(),*/, empleados.ActualizarEmpleado)
		rutes.DELETE("/delemployees/:id" /*token.TokenAuthMiddleware(),*/, empleados.EliminarEmpleado)

		/*	rutes.GET("/persona", personas.ObtenerPersona)
			rutes.POST("/persona", personas.CrearPersona)
			rutes.GET("/buscarpersona/:id", personas.BuscarPersona)
			rutes.PATCH("/modificarpersona/:id", personas.ModificarPersona)
			rutes.DELETE("/persona/:id", personas.EliminarPersona) */
	}
	r.Run()
}
