# HelixCDN Benchmark Plan

## Metrics

- Throughput
- P50 latency
- P99 latency
- Memory allocation
- Goroutine count
- Reactor scalability
- Fanout throughput
- Queue pressure
- Broadcast latency
- Packet reuse efficiency

## Benchmarks

- Reactor router
- Coalescer
- Broadcast fanout
- Buffer pool
- Allocation pressure
- Shard distribution
- Registry scaling
- Event-loop throughput

## Target Scale

- 10K concurrent clients
- 100K concurrent clients
- 1M packets/minute

## Profiling

- CPU profile
- Heap profile
- Allocation profile
- Flamegraph analysis

## Research Direction

- Zero-copy fanout
- Lock-free queue design
- Reactor scheduling
- Shard-aware multicast
- Kernel-level optimization
