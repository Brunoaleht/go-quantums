package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/helloquantum/:scriptPath", RunQuantumOnScriptHandler)
	r.GET("/hadamard/:real/:imag", HadamardHandler)
}
