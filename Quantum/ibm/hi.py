from qiskit import QuantumCircuit, BasicAer, execute

# create a quantum circuit with one qubit and one bit
qc = QuantumCircuit(1,1)
# apply not gate to the qubit with index 0
qc.x(0)
# measue the qubit with index 0 and store the result to bit with index 0
qc.measure(0,0)
# plot the circuit
qc.draw('mpl')
# simulation: run the quantum circuit 1000 times and print all results
backend = BasicAer.get_backend('qasm_simulator')
result = execute(qc, backend, shots=1000).result()
counts  = result.get_counts(qc)
print(counts)
