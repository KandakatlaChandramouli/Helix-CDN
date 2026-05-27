#!/bin/bash

set -e

mkdir -p research/runtime

go test -trace=research/runtime/trace.out \
./internal/core/mpsc

go tool trace research/runtime/trace.out
