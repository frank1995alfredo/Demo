package metodos

import (
	"net/http"

	inputsmantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento/inputsMantenimiento"
	database "github.com/frank1995alfredo/api/database"
	"github.com/frank1995alfredo/api/models/mantenimiento"
	"github.com/gin-gonic/gin"
)

//RegistrarUsuario ...
func RegistrarUsuario(c *gin.Context) {
	var input inputsmantenimiento.UserInput
	var user []mantenimiento.User

	//validamos los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ValidarEntrada() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingrese una descripción."})
		return
	}

	//pregunto si ese usuario existe en la base de datos
	if err := database.DB.Where("usuario=?", input.Usuario).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe este usuario, ingrese otro."})
		return
	}

	usuario := mantenimiento.User{Usuario: input.Usuario, Password: input.Password, EmpID: input.EmpID, Estado: input.Estado}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err := tx.Create(&usuario).Error //si no hay un error, se guarda el usuario
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": usuario})
}

//DarBajaUsuario

//DesactivarUsuario ...
func DesactivarUsuario(c *gin.Context) {
	var user mantenimiento.User

	tx := database.DB.Begin()
	err := tx.Model(&user).Where("usuario_id=?", c.Param("id")).Update("estado", false).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Usuario desactivado."})
}

//ActivarUsuario ...
func ActivarUsuario(c *gin.Context) {
	var user mantenimiento.User

	tx := database.DB.Begin()
	err := tx.Model(&user).Where("usuario_id=?", c.Param("id")).Update("estado", true).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Usuario Activado."})
}
