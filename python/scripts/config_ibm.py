from qiskit import IBMQ

IBMQ.save_account('YOUR_IBM_QUANTUM_API_TOKEN', overwrite=True)
print("Account saved!")