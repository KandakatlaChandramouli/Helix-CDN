#!/bin/bash

go test \
./benchmarks \
-bench=. \
-cpuprofile cpu.prof
