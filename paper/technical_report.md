
# Helix-CDN
## A Research-Oriented Event-Driven Runtime for Scalable Distributed Systems

---

# Abstract

[Insert abstract.md]

---

# 1. Introduction

[Insert introduction.md]

Figure 13: Overall Runtime Architecture

Figure 14: Request Lifecycle

---

# 2. Motivation

The project originated from studying large-scale streaming and content delivery systems and investigating how core runtime primitives can be implemented from first principles.

---

# 3. System Architecture

## Runtime Architecture

Figure 13

## Request Lifecycle

Figure 14

## Reactor Event Flow

Figure 15

## Work-Stealing Scheduler

Figure 16

## Storage Engine

Figure 17

---

# 4. Methodology

[Insert methodology.md]

Environment:

- Linux
- Go Runtime
- Intel Xeon 2.20 GHz

Validation Targets:

- Reactor
- Mempool
- Batcher
- WorkQueue
- WorkStealing
- SocketRuntime
- EpollCore
- EventFD

---

# 5. Benchmark Design

Benchmarks:

- Runtime Integration
- Channel Baseline
- Goroutine Baseline
- Saturation
- Tail Latency
- Scalability
- Socket Fanout
- Burst Workload

---

# 6. Evaluation

## Runtime Comparison

Figure 1

## Scalability

Figure 2
Figure 5
Figure 8
Figure 11

## Saturation

Figure 3
Figure 7

## Workload

Figure 4

## Runtime Efficiency

Figure 6
Figure 9

## Linux Validation

Figure 10

## Benchmark Overview

Figure 12

---

# 7. Results

Table: benchmark_table.md

CSV: benchmark_results.csv

Observed Results:

- Runtime Integration: 467.6 ns/op
- Channel Baseline: 108.3 ns/op
- Goroutine Baseline: 424.9 ns/op
- Queue Pressure: 24.44 ns/op
- Runtime Saturation: 234.9 ns/op
- Tail Latency: 269.7 ns/op
- Socket Fanout: 7.95 ns/op
- Burst Workload: 26512 ns/op

Scalability:

- 1 Worker: 243.2 ns/op
- 2 Workers: 206.1 ns/op
- 4 Workers: 201.6 ns/op
- 8 Workers: 188.9 ns/op
- 16 Workers: 201.7 ns/op

---

# 8. Discussion

[Insert discussion.md]

Key Observation:

Performance improves consistently up to 8 workers before showing diminishing returns at 16 workers.

---

# 9. Related Work

[Insert related_work.md]

Relevant Systems:

- NGINX
- Envoy
- Seastar
- Netty
- Tokio
- gnet
- libuv

---

# 10. Limitations

[Insert limitations.md]

---

# 11. Reproducibility

[Insert reproducibility.md]

Figure 18: Research Artifact Pipeline

Repository Structure:

repository_structure.txt

Environment:

environment.txt

---

# 12. Conclusion

[Insert conclusion.md]

---

# Artifact Inventory

Figures:
18

Tables:
1+

Architecture Diagrams:
6

Benchmark Suites:
5

Validation Targets:
8

Runtime Packages:
Hundreds of source files

Git Tag:
stable-recovery-v1

---

# Appendix

benchmark_summary.md

benchmark_results.csv

figure_manifest.md
