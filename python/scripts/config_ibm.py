from qiskit_ibm_runtime import QiskitRuntimeService
from dotenv import load_dotenv
import os

load_dotenv()

api_token = os.getenv("IBM_API_TOKEN")

QiskitRuntimeService.save_account(token=api_token, channel="ibm_quantum")
print("Account saved!")