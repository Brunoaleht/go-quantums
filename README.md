# **Quantum Go: Simulação Quântica em Go**

Este repositório contém uma implementação de portas quânticas e cálculos relacionados à esfera de Bloch utilizando a linguagem Go. Ele fornece estruturas e funções fundamentais para manipulação de números complexos e operações comuns em computação quântica.

## **📌 Funcionalidades**
- **Representação de números complexos** via a estrutura `Complex`.
- **Portas quânticas** implementadas:
  - **Hadamard (H)**: Cria superposição.
  - **Pauli (X, Y, Z)**: Operadores fundamentais na mecânica quântica.
  - **Phase Shift**: Aplica um deslocamento de fase.
- **Conversão para a Esfera de Bloch**: Representação visual dos estados quânticos.

## **🚀 Como Usar**
Basta importar os pacotes `quantum` e `bloch` para utilizar as portas e converter estados para coordenadas na esfera de Bloch.

### **Exemplo: Aplicando a porta Hadamard e obtendo coordenadas na esfera de Bloch**
```go
package main

import (
	"fmt"
	"go-quantums/quantum"
	"go-quantums/bloch"
	"math"
)

func main() {
	// Definição de um estado quântico |0⟩
	q := quantum.Complex{Real: 1, Imag: 0}

	// Aplicando a porta Hadamard
	qH, _ := quantum.Hadamard(q)

	// Conversão para coordenadas na esfera de Bloch
	x, y, z := bloch.BlochCoords(qH)

	// Exibição do resultado
	fmt.Printf("Coordenadas de |+⟩ na esfera de Bloch: (x=%.3f, y=%.3f, z=%.3f)\n", x, y, z)
}
```

## **📚 Requisitos**
- Go 1.18+
- Biblioteca `gonum.org/v1/gonum/cmplxs` para operações complexas

## **🛠️ Contribuições**
Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

---

### **Licença**
Este projeto está licenciado sob a [MIT License](LICENSE).

