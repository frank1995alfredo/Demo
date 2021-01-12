package models

import (
	"log"

	"github.com/frank1995alfredo/api/class/persona"
	cliente "github.com/frank1995alfredo/api/models/clientes"
	cargo "github.com/frank1995alfredo/api/models/empleados"
	empleado "github.com/frank1995alfredo/api/models/empleados"
	mantenimiento "github.com/frank1995alfredo/api/models/mantenimiento"
	usuario "github.com/frank1995alfredo/api/models/usuarios"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //asdads

	//github.com/lib/pq ... libreria para manejar los pq, controla los orm
	_ "github.com/lib/pq"
)

//DB ... variable global
var DB *gorm.DB

//ConectorBD ... permite conectar a la base de datos
func ConectorBD() {

	bd, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=Demo password=1234 sslmode=disable")

	if err != nil {
		log.Println(err.Error())
	}

	bd.AutoMigrate(&mantenimiento.Provincia{}, &mantenimiento.Ciudad{})
	bd.Model(&mantenimiento.Ciudad{}).AddForeignKey("pro_id", "provincia(provincia_id)", "SET NULL", "")

	bd.AutoMigrate(&mantenimiento.Discapacidad{})

	bd.AutoMigrate(&cliente.Cliente{})
	bd.Model(&cliente.Cliente{}).AddForeignKey("ciu_id", "ciudads(ciudad_id)", "SET NULL", "CASCADE")
	bd.Model(&cliente.Cliente{}).AddForeignKey("disc_id", "discapacidads(discapacidad_id)", "SET NULL", "")

	bd.AutoMigrate(&cargo.CargoEmp{})

	bd.AutoMigrate(&empleado.Empleado{})
	bd.Model(&empleado.Empleado{}).AddForeignKey("cargo_emp_id", "cargo_emps(cargo_emp_id)", "SET NULL", "")
	bd.Model(&empleado.Empleado{}).AddForeignKey("ciu_id", "ciudads(ciudad_id)", "SET NULL", "")
	bd.Model(&empleado.Empleado{}).AddForeignKey("disc_id", "discapacidads(discapacidad_id)", "SET NULL", "")

	bd.AutoMigrate(&usuario.User{})
	bd.Model(&usuario.User{}).AddForeignKey("emp_id", "empleados(empleado_id)", "SET NULL", "")

	bd.AutoMigrate(&persona.Persona{})

	log.Println("Migracion satisfactoria")

	DB = bd //toma la conexion de la base de datos

}
