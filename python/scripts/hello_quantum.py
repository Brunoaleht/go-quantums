from qiskit import QuantumCircuit, Aer, transpile
from qiskit.visualization import plot_histogram
import sys
import json

def run_quantum_circuit():
    circuit = QuantumCircuit(1, 1)
    circuit.h(0)  # Aplica a porta Hadamard
    circuit.measure([0], [0])

    # Use o simulador local
    simulator = Aer.get_backend('qasm_simulator')

    # Execute o circuito quântico
    compiled_circuit = transpile(circuit, simulator)
    job = simulator.run(compiled_circuit)
    result = job.result()

    counts = result.get_counts(circuit)
    print(json.dumps(counts))  # Retornar os resultados como JSON para integração

if __name__ == "__main__":
    run_quantum_circuit()