package main

import (
	"go-quantums/api"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Conecta ao banco de dados
	// database.ConnectDatabase()

	// Registra rotas
	api.RegisterRoutes(r)

	r.Run(":8080")

}
