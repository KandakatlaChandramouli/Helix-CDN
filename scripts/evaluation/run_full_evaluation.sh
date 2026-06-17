#!/bin/bash

set -e

mkdir -p research/evaluation

echo "=== Runtime Integration ==="

go test -bench=. -benchmem ./benchmarks/runtimeintegration \
| tee research/evaluation/runtimeintegration.txt

echo "=== Scheduler Core ==="

go test -bench=. -benchmem ./benchmarks/schedulercore \
| tee research/evaluation/schedulercore.txt

echo "=== Saturation ==="

go test -bench=. -benchmem ./benchmarks/saturation \
| tee research/evaluation/saturation.txt

echo "=== Comparison ==="

go test -bench=. -benchmem ./benchmarks/comparison \
| tee research/evaluation/comparison.txt

echo "=== Ringbuffer ==="

go test -bench=. -benchmem ./benchmarks/ringbuffer \
| tee research/evaluation/ringbuffer.txt

echo "=== Epoll Runtime ==="

go test -bench=. -benchmem ./benchmarks/epollruntime \
| tee research/evaluation/epollruntime.txt
