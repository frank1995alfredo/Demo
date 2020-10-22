package interfaces

import "github.com/gin-gonic/gin"

// ICliente ...
type ICliente interface {
	ObtenerCliente(c *gin.Context)
}
