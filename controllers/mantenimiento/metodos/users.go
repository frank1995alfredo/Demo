package metodos

import (
	"net/http"

	config "github.com/frank1995alfredo/api/config"
	inputsmantenimiento "github.com/frank1995alfredo/api/controllers/mantenimiento/inputsMantenimiento"
	database "github.com/frank1995alfredo/api/database"
	"github.com/frank1995alfredo/api/models/mantenimiento"
	token "github.com/frank1995alfredo/api/token"

	"github.com/gin-gonic/gin"
)

//RegistrarUsuario ...
func RegistrarUsuario(c *gin.Context) {
	var input inputsmantenimiento.UserInput
	var user mantenimiento.User

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

	//validamos los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ValidarEntrada() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingrese un usuario y una contraceña."})
		return
	}

	//pregunto si ese usuario existe en la base de datos
	if err := database.DB.Where("usuario LIKE ?", input.Usuario).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe este usuario, ingrese otro."})
		return
	}

	password, _ := config.HashPassword(input.Password)
	usuario := mantenimiento.User{Usuario: input.Usuario, Password: password, EmpID: input.EmpID, Estado: input.Estado}

	//inicio de la transaccion
	tx := database.DB.Begin()
	err2 := tx.Create(&usuario).Error //si no hay un error, se guarda el usuario
	if err2 != nil {
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

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

	tx := database.DB.Begin()
	err2 := tx.Model(&user).Where("usuario_id=?", c.Param("id")).Update("estado", false).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Usuario desactivado."})
}

//ActivarUsuario ...
func ActivarUsuario(c *gin.Context) {
	var user mantenimiento.User

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

	tx := database.DB.Begin()
	err2 := tx.Model(&user).Where("usuario_id=?", c.Param("id")).Update("estado", true).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Usuario Activado."})
}
