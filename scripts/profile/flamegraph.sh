#!/bin/bash

set -e

mkdir -p research/flamegraphs

go test -bench=. \
-cpuprofile=research/flamegraphs/cpu.prof \
./benchmarks/corequeue

go tool pprof -top research/flamegraphs/cpu.prof \
> research/flamegraphs/flamegraph.txt
