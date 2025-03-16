package bloch

import (
	"go-quantums/quantum"
	"math"
	"math/cmplx"
)

// Calcula as coordenadas na esfera de Bloch
func BlochCoords(q quantum.Complex) (float64, float64, float64) {
	complexNum := complex(q.Real, q.Imag)
	magnitude := cmplx.Abs(complexNum)
	if magnitude == 0 {
		return 0, 0, 1 // Estado |0⟩ na esfera de Bloch
	}

	theta := 2 * math.Acos(q.Real/magnitude)
	phi := math.Atan2(q.Imag, q.Real)
	x := math.Sin(theta) * math.Cos(phi)
	y := math.Sin(theta) * math.Sin(phi)
	z := math.Cos(theta)
	return x, y, z
}
