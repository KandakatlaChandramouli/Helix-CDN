#!/bin/bash

set -e

mkdir -p research/runtime

go test -bench=. \
-benchmem \
-cpuprofile=research/runtime/cpu.prof \
./benchmarks/corequeue

go tool pprof -top research/runtime/cpu.prof \
> research/runtime/cpu_top.txt
