# GPU Numba and CUDA and Cupy:
# https://www.youtube.com/watch?v=wa0EmEq5Otw

from qiskit import QuantumCircuit, Aer, execute
import numpy as np

# Define the number to be factored
N = 21

# Define the number of qubits needed to represent N
n = np.ceil(np.log2(N)).astype(int)

# Define the quantum circuit
qc = QuantumCircuit(n*2, n)

# Apply Hadamard gates to the first n qubits
for i in range(n):
    qc.h(i)

# Define the function f(a) = a^x mod N, where x = 2^n / r for some r
def f(a):
    return pow(a, 2**(n//2), N)

# Apply the function f(a) to the second n qubits using modular exponentiation
for i in range(n):
    qc.swap(i, i+n)
    qc.barrier()
    qc.x(i+n)
    qc.swap(i, i+n)
    qc.barrier()
    qc.compose(QuantumCircuit(n*2, n).from_function(f), [i+n for i in range(n)], inplace=True)
    qc.barrier()
    qc.swap(i, i+n)

# Apply inverse Fourier transform to the first n qubits
for i in range(n):
    qc.h(i)
    for j in range(i+1, n):
        qc.cu1(-2*np.pi/2**(j-i), j, i)
    qc.barrier()

# Measure the first n qubits
for i in range(n):
    qc.measure(i, i)

# Simulate the quantum circuit using the statevector simulator
backend = Aer.get_backend('statevector_simulator')
job = execute(qc, backend)
result = job.result()
psi = result.get_statevector()

# Find the period r by measuring the first n qubits
counts = result.get_counts()
for c in counts:
    r = int(c, 2)
    if np.gcd(r, N) != 1:
        break

# Factor the number using the period r
x = pow(2, n//2)
a = pow(x + r, N//2, N)
factor1 = np.gcd(a - 1, N)
factor2 = np.gcd(a + 1, N)

# Print the factors
print(f"Factors of {N}: {factor1} and {factor2}")
