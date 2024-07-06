package main

import (
	"fmt"
	"time"

	"github.com/Hoosat-Oy/HTND/domain/consensus/utils/hashes"
	"github.com/Hoosat-Oy/HTND/domain/consensus/utils/pow"
	// Import other necessary packages
)

func BenchmarkMatrixHoohashRev1() {
    input := []byte("BenchmarkMatrix_HeavyHash")
    firstPass := hashes.Blake3HashWriter()
    firstPass.InfallibleWrite(input)
    hash := firstPass.Finalize()
    matrix := pow.generateHoohashMatrix(hash)
    multiplied := matrix.HoohashMatrixMultiplication(hash)
    secondPass := hashes.Blake3HashWriter()
    secondPass.InfallibleWrite(multiplied)
    hash = secondPass.Finalize()
}


func main() {
    iterations := 0
    startTime := time.Now()

    for {
        BenchmarkMatrixHoohashRev1()
        iterations++

        if iterations%1000 == 0 {
            elapsed := time.Since(startTime)
            opsPerSecond := float64(iterations) / elapsed.Seconds()
            fmt.Printf("Iterations: %d, Time: %v, Ops/sec: %.2f\n", iterations, elapsed, opsPerSecond)
        }
    }
}