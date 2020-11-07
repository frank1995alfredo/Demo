package models

import (
	"log"

	"github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/frank1995alfredo/api/models/practica"

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
	bd.Model(&mantenimiento.Ciudad{}).AddForeignKey("pro_id", "provincia(provincia_id)", "SET NULL", "CASCADE")

	bd.AutoMigrate(&mantenimiento.Discapacidad{})

	bd.AutoMigrate(&mantenimiento.Cliente{})
	bd.Model(&mantenimiento.Cliente{}).AddForeignKey("ciu_id", "ciudads(ciudad_id)", "SET NULL", "CASCADE")
	bd.Model(&mantenimiento.Cliente{}).AddForeignKey("disc_id", "discapacidads(discapacidad_id)", "SET NULL", "CASCADE")

	bd.AutoMigrate(&mantenimiento.CargoEmp{})

	bd.AutoMigrate(&mantenimiento.Empleado{})
	bd.Model(&mantenimiento.Empleado{}).AddForeignKey("cargo_emp_id", "cargo_emps(cargo_emp_id)", "SET NULL", "CASCADE")
	bd.Model(&mantenimiento.Empleado{}).AddForeignKey("ciu_id", "ciudads(ciudad_id)", "SET NULL", "CASCADE")
	bd.Model(&mantenimiento.Empleado{}).AddForeignKey("disc_id", "discapacidads(discapacidad_id)", "SET NULL", "CASCADE")

	bd.AutoMigrate(&practica.Contacto{}, &practica.Persona{})
	bd.Model(&practica.Contacto{}).AddForeignKey("per_id", "personas(persona_id)", "SET NULL", "CASCADE")

	bd.AutoMigrate(&mantenimiento.User{})
	bd.Model(&mantenimiento.User{}).AddForeignKey("emp_id", "empleados(empleado_id)", "SET NULL", "CASCADE")

	log.Println("Migracion satisfactoria")

	DB = bd //toma la conexion de la base de datos
}
