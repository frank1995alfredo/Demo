package practicas

import (
	"net/http"

	inputPractica "github.com/frank1995alfredo/api/controllers/practicas/inputPractica"

	database "github.com/frank1995alfredo/api/database"
	"github.com/frank1995alfredo/api/models/clases"
	practica "github.com/frank1995alfredo/api/models/practica"

	"github.com/gin-gonic/gin"
)

//ObtenerCliente ..
func ObtenerCliente(c *gin.Context) {

	var cliente []clases.ClientePersona

	//database.DB.Order("cliente_id").Find("cliente").Error

	//result := map[string]interface{}{}
	database.DB.Table("clientes").Order("cliente_id").Find(&cliente)

	c.SecureJSON(http.StatusOK, gin.H{"data": cliente})

	/*if err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error})
	} else {
		c.SecureJSON(http.StatusOK, gin.H{"data": cliente})
	}*/
}

//CrearPersona ...
func CrearPersona(c *gin.Context) {

	var input inputPractica.PersonaInput
	var per practica.Persona

	//validas los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingrese un nombre."})
		return
	}

	if err := database.DB.Where("nombre=?", input.Nombre).First(&per).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe este nombre, ingrese otro."})
		return
	}

	persona := practica.Persona{Nombre: input.Nombre, Apellido: input.Apellido, Direccion: input.Direccion, Ciudad: input.Ciudad}

	tx := database.DB.Begin()
	err := tx.Create(&persona).Error //si no hay un error, se guarda el articulo
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusOK, gin.H{"data": persona})
}

//BuscarPersona ...
func BuscarPersona(c *gin.Context) {

	var persona practica.Persona

	if err := database.DB.Where("apellido=?", c.Param("apellido")).First(&persona).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe esa persona."})
		return
	}

	c.SecureJSON(http.StatusFound, gin.H{"data": persona})
}

//ActualizarPersona ...
func ActualizarPersona(c *gin.Context) {

	var input inputPractica.PersonaInput
	var per practica.Persona

	//validamos la entrada de los datos
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	persona := practica.Persona{Nombre: input.Nombre, Apellido: input.Apellido, Direccion: input.Direccion, Ciudad: input.Ciudad}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Model(&per).Where("persona_id=?", c.Param("id")).Update(persona).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	c.SecureJSON(http.StatusCreated, gin.H{"data": "Persona actualizada correctamente"})
}

//EliminarPersona ...
func EliminarPersona(c *gin.Context) {

	var persona practica.Persona

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Where("persona_id=?", c.Param("id")).Delete(&persona).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Registro eliminado."})

}
