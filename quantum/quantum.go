package quantum

import "math"

// Estrutura para representar um número complexo
type Complex struct {
	Real float64 `json:"real"`
	Imag float64 `json:"imag"`
}

// Porta Hadamard
func Hadamard(q Complex) (Complex, Complex) {
	h := 1.0 / math.Sqrt(2)
	return Complex{
			Real: h * (q.Real + q.Imag),
			Imag: h * (q.Imag + q.Real),
		}, Complex{
			Real: h * (q.Real - q.Imag),
			Imag: h * (q.Imag - q.Real),
		}
}

// Porta Pauli-X CNOT
func PauliX(q Complex) Complex {
	return Complex{Real: q.Imag, Imag: q.Real}
}

// Porta Pauli-Y
func PauliY(q Complex) Complex {
	return Complex{Real: -q.Imag, Imag: q.Real}
}

// Porta Pauli-Z
func PauliZ(q Complex) Complex {
	return Complex{Real: q.Real, Imag: -q.Imag}
}

// Porta de Fase
func PhaseShift(q Complex, phi float64) Complex {
	return Complex{
		Real: q.Real*math.Cos(phi) - q.Imag*math.Sin(phi),
		Imag: q.Real*math.Sin(phi) + q.Imag*math.Cos(phi),
	}
}
