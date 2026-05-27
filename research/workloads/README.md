# Runtime Workloads

This directory documents runtime workload scenarios used for evaluation.

## Workloads

### Burst Workload
Simulates sudden scheduler spikes with high fanout execution bursts.

### Socket Fanout
Simulates event fanout behavior under runtime queue pressure.

### Runtime Saturation
Measures scheduler degradation under sustained overload.

### Core Scaling
Measures runtime throughput as worker count increases.

## Metrics

Collected metrics include:
- throughput
- ns/op
- allocations/op
- queue pressure
- tail latency
- runtime saturation
