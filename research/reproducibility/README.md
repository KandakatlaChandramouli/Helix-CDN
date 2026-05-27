# Helix Runtime Reproducibility

## Runtime Environment

This repository stores:
- benchmark outputs
- flamegraphs
- cpu profiles
- heap profiles
- runtime traces

## Validation Commands

### Run Full Tests

go test ./...

### Run Race Detection

go test -race ./...

### Run Benchmarks

go test -bench=. -benchmem ./...

### Generate CPU Profiles

./scripts/profile/cpu_profile.sh

### Generate Heap Profiles

./scripts/profile/heap_profile.sh

### Generate Runtime Traces

./scripts/profile/trace_runtime.sh
