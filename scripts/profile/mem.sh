#!/bin/bash

go test \
./benchmarks \
-bench=. \
-memprofile mem.prof
