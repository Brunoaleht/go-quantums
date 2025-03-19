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

- Primeiramente Executar o `config_ibm.py`, para configurar as chaves da IBM
  - comando: `python config_ibm.py`
- Observar os script em python, para usar na rota
- Após executar `main.go`, e executar as rotas desejadas
  -comando: `go run main.go`

## **📚 Requisitos**

- Python3
- Go 1.18+
- Biblioteca:  
  `gonum.org/v1/gonum/cmplxs`
  `pip install qiskit`
  `github.com/gin-gonic/gin`
  `pip install qiskit-ibm-runtime`

## **🛠️ Contribuições**

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

---

### **Licença**

Este projeto está licenciado sob a [MIT License](LICENSE).
