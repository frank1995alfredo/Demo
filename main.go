package main

import (
	config "github.com/frank1995alfredo/api/config"
	database "github.com/frank1995alfredo/api/database"
	rutas "github.com/frank1995alfredo/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(config.CORS)

	database.ConectorBD()
	defer database.DB.Close()

	rutas.Rutas()
	r.Run()
}
