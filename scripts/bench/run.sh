#!/bin/bash

echo "running helix benchmarks"

go test \
./benchmarks/... \
-bench=. \
-benchmem \
-run=^$
