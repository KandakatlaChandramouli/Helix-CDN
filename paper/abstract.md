
# Abstract

Helix-CDN is a lightweight event-driven runtime designed to explore the architectural foundations of high-performance distributed systems. The system combines reactor-driven event processing, adaptive scheduling, lock-free queues, memory pooling, batching, and Linux-native I/O primitives to evaluate runtime behavior under scalable workloads.

Linux validation was performed using epoll and eventfd backed implementations. Benchmark evaluation demonstrated sub-microsecond runtime integration overhead, scalable worker scheduling, low-latency queue operations, and efficient socket fanout behavior.

This artifact presents the implementation, evaluation methodology, benchmark results, reproducibility package, and engineering lessons learned during development.
