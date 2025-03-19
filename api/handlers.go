package api

import (
	"encoding/json"
	"fmt"
	"go-quantums/bloch"
	"go-quantums/quantum"
	"net/http"
	"os/exec"
	"sync"

	"github.com/gin-gonic/gin"
)

type Coords struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Z     float64 `json:"z"`
	Phi   float64 `json:"phi"`
	Theta float64 `json:"theta"`
}

type Result struct {
	Q0    quantum.Complex `json:"q0"`
	Q1    quantum.Complex `json:"q1"`
	Cord0 Coords          `json:"cord0"`
	Cord1 Coords          `json:"cord1"`
}

// func parseParam(param string) float64 {
// 	value, err := strconv.ParseFloat(param, 64)
// 	if err != nil {
// 		return 0.0
// 	}
// 	return value
// }

func HadamardHandler(c *gin.Context) {
	var input quantum.Complex
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	q := quantum.Complex{Real: input.Real, Imag: input.Imag}

	results := make(chan Result)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		q0, q1 := quantum.Hadamard(q)
		x0, y0, z0, theta0, phi0 := bloch.BlochCoords(q0)
		x1, y1, z1, theta1, phi1 := bloch.BlochCoords(q1)

		result := Result{
			Q0:    quantum.Complex{Real: q0.Real, Imag: q0.Imag},
			Q1:    quantum.Complex{Real: q1.Real, Imag: q1.Imag},
			Cord0: Coords{X: x0, Y: y0, Z: z0, Phi: phi0, Theta: theta0},
			Cord1: Coords{X: x1, Y: y1, Z: z1, Phi: phi1, Theta: theta1},
		}

		results <- result
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	result := <-results
	c.JSON(http.StatusOK, gin.H{
		"q0":           fmt.Sprintf("(%.2f, %.2f)", result.Q0.Real, result.Q0.Imag),
		"theta0, phi0": fmt.Sprintf("(%.2f, %.2f)", result.Cord0.Theta, result.Cord0.Phi),
		"x0, y0, z0":   fmt.Sprintf("(%.2f, %.2f, %.2f)", result.Cord0.X, result.Cord0.Y, result.Cord0.Z),
		"q1":           fmt.Sprintf("(%.2f, %.2f)", result.Q1.Real, result.Q1.Imag),
		"theta1, phi1": fmt.Sprintf("(%.2f, %.2f)", result.Cord1.Theta, result.Cord1.Phi),
		"x1, y1, z1":   fmt.Sprintf("(%.2f, %.2f, %.2f)", result.Cord1.X, result.Cord1.Y, result.Cord1.Z),
	})
}

func PauliXHandler(c *gin.Context) {

	var input quantum.Complex
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	q := quantum.Complex{Real: input.Real, Imag: input.Imag}

	results := make(chan struct {
		QX   quantum.Complex
		Cord Coords
	})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		qX := quantum.PauliX(q)
		x, y, z, phi, theta := bloch.BlochCoords(qX)

		results <- struct {
			QX   quantum.Complex
			Cord Coords
		}{
			QX:   qX,
			Cord: Coords{X: x, Y: y, Z: z, Phi: phi, Theta: theta},
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

func PauliYHandler(c *gin.Context) {
	var input quantum.Complex
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	q := quantum.Complex{Real: input.Real, Imag: input.Imag}

	results := make(chan struct {
		QY   quantum.Complex
		Cord Coords
	})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		qY := quantum.PauliY(q)
		x, y, z, phi, theta := bloch.BlochCoords(qY)

		results <- struct {
			QY   quantum.Complex
			Cord Coords
		}{
			QY:   qY,
			Cord: Coords{X: x, Y: y, Z: z, Phi: phi, Theta: theta},
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	result := <-results
	c.JSON(http.StatusOK, gin.H{
		"qY":      fmt.Sprintf("(%.2f, %.2f)", result.QY.Real, result.QY.Imag),
		"x, y, z": fmt.Sprintf("(%.2f, %.2f, %.2f)", result.Cord.X, result.Cord.Y, result.Cord.Z),
	})
}

func PauliZHandler(c *gin.Context) {
	var input quantum.Complex
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	q := quantum.Complex{Real: input.Real, Imag: input.Imag}

	results := make(chan struct {
		QZ   quantum.Complex
		Cord Coords
	})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		qZ := quantum.PauliZ(q)
		x, y, z, phi, theta := bloch.BlochCoords(qZ)

		results <- struct {
			QZ   quantum.Complex
			Cord Coords
		}{
			QZ:   qZ,
			Cord: Coords{X: x, Y: y, Z: z, Phi: phi, Theta: theta},
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	result := <-results
	c.JSON(http.StatusOK, gin.H{
		"qZ":      fmt.Sprintf("(%.2f, %.2f)", result.QZ.Real, result.QZ.Imag),
		"x, y, z": fmt.Sprintf("(%.2f, %.2f, %.2f)", result.Cord.X, result.Cord.Y, result.Cord.Z),
	})
}

func PhaseShiftHandler(c *gin.Context) {
	var input struct {
		Real float64 `json:"real"`
		Imag float64 `json:"imag"`
		Phi  float64 `json:"phi"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	q := quantum.Complex{Real: input.Real, Imag: input.Imag}

	results := make(chan struct {
		QPS  quantum.Complex
		Cord Coords
	})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		qPS := quantum.PhaseShift(q, input.Phi)
		x, y, z, phi, theta := bloch.BlochCoords(qPS)

		results <- struct {
			QPS  quantum.Complex
			Cord Coords
		}{
			QPS:  qPS,
			Cord: Coords{X: x, Y: y, Z: z, Phi: phi, Theta: theta},
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	result := <-results
	c.JSON(http.StatusOK, gin.H{
		"qPS":     fmt.Sprintf("(%.2f, %.2f)", result.QPS.Real, result.QPS.Imag),
		"phi":     fmt.Sprintf("%.2f", input.Phi),
		"theta":   fmt.Sprintf("%.2f", result.Cord.Theta),
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
