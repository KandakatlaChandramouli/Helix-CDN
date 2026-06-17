#!/bin/bash

set -e

mkdir -p research/csv

go test -bench=. -benchmem ./benchmarks/runtimeintegration \
| tee research/csv/runtimeintegration.csv

go test -bench=. -benchmem ./benchmarks/comparison \
| tee research/csv/comparison.csv

go test -bench=. -benchmem ./benchmarks/saturation \
| tee research/csv/saturation.csv

go test -bench=. -benchmem ./benchmarks/scalability \
| tee research/csv/scalability.csv

go test -bench=. -benchmem ./benchmarks/workload \
| tee research/csv/workload.csv
