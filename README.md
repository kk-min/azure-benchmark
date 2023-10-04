# snapstart-benchmark
Program to perform manual benchmarking for AWS Lambda's SnapStart.

# Usage
1. `./scripts/setup.sh` (You may need to run `chmod +x ./scripts/setup.sh` to make it executable)
2. `go build main.go`
3. `./main [OPTIONS]`

# Options
- `-b`: Burst count (integer)
- `-i`: Inter-arrival time (IAT) (milliseconds)
- `-o`: Output directory (dir path)
