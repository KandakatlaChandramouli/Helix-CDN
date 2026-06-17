#!/bin/bash

set -e

mkdir -p research/memory

go test -bench=. \
-benchmem \
-memprofile=research/memory/mem.prof \
./benchmarks/ringbuffer

go tool pprof -top research/memory/mem.prof \
> research/memory/mem_top.txt
