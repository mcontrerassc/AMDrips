import matplotlib.pyplot as plt
import numpy as np
import re
import subprocess

# Run the Go benchmark and capture the output
result = subprocess.run(['go', 'test', '-bench=.', '-benchtime=10s'], capture_output=True, text=True)
output = result.stdout

# Parse the benchmark results
sizes = []
times = []
for line in output.split('\n'):
    if line.startswith('BenchmarkMSM/Size-'):
        match = re.search(r'Size-(\d+)-\d+\s+\d+\s+([\d.]+)\s+ns/op', line)
        if match:
            size = int(match.group(1))
            time = float(match.group(2))
            sizes.append(size)
            times.append(time)

# Convert times to milliseconds
times = [t / 1e6 for t in times]

# Create the plot
plt.figure(figsize=(10, 6))
plt.plot(sizes, times, 'bo-')
plt.xscale('log')
plt.yscale('log')
plt.xlabel('Number of Points')
plt.ylabel('Time per Operation (ms)')
plt.title('BLS12-377 Multi-Scalar Multiplication Performance')
plt.grid(True)

# Add annotations
for i, (size, time) in enumerate(zip(sizes, times)):
    plt.annotate(f'{time:.2f} ms', (size, time), textcoords="offset points", xytext=(0,10), ha='center')

plt.tight_layout()
plt.savefig('msm_benchmark.png')
plt.show()
