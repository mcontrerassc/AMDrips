import matplotlib.pyplot as plt
import numpy as np
import re
import subprocess

# Run the Go benchmark and capture the output
result = subprocess.run(['go', 'test', '-bench=.', '-benchtime=10s'], capture_output=True, text=True)
output = result.stdout

# Parse the benchmark results
addition_times = []
multiplication_times = []
inversion_times = []

for line in output.split('\n'):
    if line.startswith('BenchmarkFieldAddition'):
        match = re.search(r'BenchmarkFieldAddition-\d+\s+(\d+) ns/op', line)
        if match:
            addition_times.append(int(match.group(1)))
    elif line.startswith('BenchmarkFieldMultiplication'):
        match = re.search(r'BenchmarkFieldMultiplication-\d+\s+(\d+) ns/op', line)
        if match:
            multiplication_times.append(int(match.group(1)))
    elif line.startswith('BenchmarkFieldInversion'):
        match = re.search(r'BenchmarkFieldInversion-\d+\s+(\d+) ns/op', line)
        if match:
            inversion_times.append(int(match.group(1)))

# Create the plot
plt.figure(figsize=(10, 6))
plt.plot(addition_times, label='Addition')
plt.plot(multiplication_times, label='Multiplication')
plt.plot(inversion_times, label='Inversion')
plt.xlabel('Iteration')
plt.ylabel('Time (ns)')
plt.title('Field Operations Performance')
plt.legend()
plt.grid(True)

plt.tight_layout()
plt.savefig('field_operations_benchmark.png')
plt.show()
