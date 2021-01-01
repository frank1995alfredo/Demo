package usuarios

import (
	"net/http"

	config "github.com/frank1995alfredo/api/config"

	database "github.com/frank1995alfredo/api/database"
	usuario "github.com/frank1995alfredo/api/models/usuarios"
	"github.com/gin-gonic/gin"
)

//RegistrarUsuario ...
func RegistrarUsuario(c *gin.Context) {
	var input UsuarioInput
	var user usuario.User

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.SecureJSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

	//validamos los inputs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ValidarEntrada() {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Ingrese un usuario y una contraceña."})
		return
	}

	//pregunto si ese usuario existe en la base de datos
	if err := database.DB.Where("usuario LIKE ?", input.Usuario).First(&user).Error; err == nil {
		c.SecureJSON(http.StatusBadRequest, gin.H{"error": "Ya existe este usuario, ingrese otro."})
		return
	}

	password, _ := config.HashPassword(input.Password)
	usuario := usuario.User{Usuario: input.Usuario, Password: password, Estado: true}

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
	var user usuario.User

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.SecureJSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

	tx := database.DB.Begin()
	err2 := tx.Model(&user).Where("usuario_id=?", c.Param("id")).Update("estado", false).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Usuario ha desactivado correctamente."})
}

//ActivarUsuario ...
func ActivarUsuario(c *gin.Context) {
	var user usuario.User

	//se extrae los metadatos del token, si se esta autenticado, se presentaran los datos
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.SecureJSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}

	tx := database.DB.Begin()
	err2 := tx.Model(&user).Where("usuario_id=?", c.Param("id")).Update("estado", true).Error
	if err2 != nil {
		tx.Rollback()
	}
	tx.Commit()
	//fin de la transaccion

	c.SecureJSON(http.StatusOK, gin.H{"data": "Usuario ha sido activado correctamente."})
}
