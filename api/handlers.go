package api

import (
	"fmt"
	"go-quantums/bloch"
	"go-quantums/quantum"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseParam(param string) float64 {
	value, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0.0
	}
	return value
}

func HadamardHandler(c *gin.Context) {
	real := c.Param("real")
	imag := c.Param("imag")

	q := quantum.Complex{
		Real: parseParam(real),
		Imag: parseParam(imag),
	}

	q0, q1 := quantum.Hadamard(q)
	x0, y0, z0 := bloch.BlochCoords(q0)
	x1, y1, z1 := bloch.BlochCoords(q1)

	c.JSON(http.StatusOK, gin.H{
		"q0":         fmt.Sprintf("(%.2f, %.2f)", q0.Real, q0.Imag),
		"x0, y0, z0": fmt.Sprintf("(%.2f, %.2f, %.2f)", x0, y0, z0),
		"q1":         fmt.Sprintf("(%.2f, %.2f)", q1.Real, q1.Imag),
		"x1, y1, z1": fmt.Sprintf("(%.2f, %.2f, %.2f)", x1, y1, z1),
	})
}

func PauliXHandler(c *gin.Context) {
	real := c.Param("real")
	imag := c.Param("imag")

	q := quantum.Complex{
		Real: parseParam(real),
		Imag: parseParam(imag),
	}

	qX := quantum.PauliX(q)
	x, y, z := bloch.BlochCoords(qX)

	c.JSON(http.StatusOK, gin.H{
		"qX":      fmt.Sprintf("(%.2f, %.2f)", qX.Real, qX.Imag),
		"x, y, z": fmt.Sprintf("(%.2f, %.2f, %.2f)", x, y, z),
	})
}
