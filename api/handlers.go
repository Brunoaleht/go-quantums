package api

import (
	"encoding/json"
	"fmt"
	"go-quantums/bloch"
	"go-quantums/quantum"
	"net/http"
	"os/exec"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Coords struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Result struct {
	Q0    quantum.Complex `json:"q0"`
	Q1    quantum.Complex `json:"q1"`
	Cord0 Coords          `json:"cord0"`
	Cord1 Coords          `json:"cord1"`
}

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

	results := make(chan Result)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		q0, q1 := quantum.Hadamard(q)
		x0, y0, z0 := bloch.BlochCoords(q0)
		x1, y1, z1 := bloch.BlochCoords(q1)

		result := Result{
			Q0:    quantum.Complex{Real: q0.Real, Imag: q0.Imag},
			Q1:    quantum.Complex{Real: q1.Real, Imag: q1.Imag},
			Cord0: Coords{X: x0, Y: y0, Z: z0},
			Cord1: Coords{X: x1, Y: y1, Z: z1},
		}

		results <- result
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	result := <-results
	c.JSON(http.StatusOK, gin.H{
		"q0":         fmt.Sprintf("(%.2f, %.2f)", result.Q0.Real, result.Q0.Imag),
		"x0, y0, z0": fmt.Sprintf("(%.2f, %.2f, %.2f)", result.Cord0.X, result.Cord0.Y, result.Cord0.Z),
		"q1":         fmt.Sprintf("(%.2f, %.2f)", result.Q1.Real, result.Q1.Imag),
		"x1, y1, z1": fmt.Sprintf("(%.2f, %.2f, %.2f)", result.Cord1.X, result.Cord1.Y, result.Cord1.Z),
	})
}

func PauliXHandler(c *gin.Context) {
	real := c.Param("real")
	imag := c.Param("imag")

	q := quantum.Complex{
		Real: parseParam(real),
		Imag: parseParam(imag),
	}

	results := make(chan struct {
		QX   quantum.Complex
		Cord Coords
	})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		qX := quantum.PauliX(q)
		x, y, z := bloch.BlochCoords(qX)

		results <- struct {
			QX   quantum.Complex
			Cord Coords
		}{
			QX:   qX,
			Cord: Coords{X: x, Y: y, Z: z},
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	// Lendo o resultado do canal antes de responder
	result := <-results
	c.JSON(http.StatusOK, gin.H{
		"qX":      fmt.Sprintf("(%.2f, %.2f)", result.QX.Real, result.QX.Imag),
		"x, y, z": fmt.Sprintf("(%.2f, %.2f, %.2f)", result.Cord.X, result.Cord.Y, result.Cord.Z),
	})
}

func RunQuantumOnScriptHandler(c *gin.Context) {
	scriptPath := c.Param("scriptPath")

	currentPath := fmt.Sprintf("/python/scripts/%s", scriptPath)

	cmd := exec.Command("python", currentPath)
	output, err := cmd.Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(output, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse result"})
		return
	}

	// Salvar no banco de dados
	// database.SaveToDatabase(database.DB, "Quantum Simulation", fmt.Sprintf("%v", result))

	c.JSON(http.StatusOK, result)
}
