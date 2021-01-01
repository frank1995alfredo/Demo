package functions

import (
	"net/http"

	token "github.com/frank1995alfredo/api/controllers/usuarios"
	"github.com/gin-gonic/gin"
)

// ValidarToken ... VALIDAMOS EL TOKEN
func ValidarToken() {
	c := gin.Context{}
	_, err := token.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.SecureJSON(http.StatusUnauthorized, "No tiene permisos necesarios.")
		return
	}
}
