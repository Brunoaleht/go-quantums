package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/quantum/:scriptPath", RunQuantumOnScriptHandler)
	r.POST("/hadamard", HadamardHandler)
	r.POST("/paulix", PauliXHandler)
	r.POST("/pauliy", PauliYHandler)
	r.POST("/pauliz", PauliZHandler)
	r.POST("/phaseshift", PhaseShiftHandler)
}
