#!/bin/bash

set -e

mkdir -p research/flamegraphs

go test -bench=BenchmarkRuntimeIntegration \
-cpuprofile=research/flamegraphs/runtime_cpu.prof \
./benchmarks/runtimeintegration

go tool pprof \
-png \
research/flamegraphs/runtime_cpu.prof \
> research/flamegraphs/runtime_cpu.png || true

go test -bench=BenchmarkRuntimeSaturation \
-cpuprofile=research/flamegraphs/saturation_cpu.prof \
./benchmarks/saturation

go tool pprof \
-png \
research/flamegraphs/saturation_cpu.prof \
> research/flamegraphs/saturation_cpu.png || true
