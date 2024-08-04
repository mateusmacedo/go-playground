package main

import (
	"fmt"
	"sync"

	"github.com/mateusmacedo/go-playground/pkg/concurrency/helpers"
)

type Ratios struct {
	PositiveProportion float64
	NegativeProportion float64
	ZeroProportion     float64
}

// classifyData classifies the data into positives, negatives, and zeros and calculates proportions.
func classifyData(data []int) Ratios {
	countPos, countNeg, countZero := 0, 0, 0

	for _, value := range data {
		switch {
		case value > 0:
			countPos++
		case value < 0:
			countNeg++
		default:
			countZero++
		}
	}

	total := float64(len(data))

	return Ratios{
		PositiveProportion: float64(countPos) / total,
		NegativeProportion: float64(countNeg) / total,
		ZeroProportion:     float64(countZero) / total,
	}
}

// launchWorkers starts worker goroutines to process segments of the data.
func launchWorkers(data [][]int, results chan<- Ratios, wg *sync.WaitGroup) {
	for _, segment := range data {
		wg.Add(1)
		go func(s []int) {
			defer wg.Done()
			result := classifyData(s)
			results <- result
		}(segment)
	}
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 4
	results := make(chan Ratios, numGoroutines)

	data := []int{-4, -3, -2, -1, 0, 1, 2, 3, 4, 0}
	fmt.Println("Original slice:", data)

	segments := helpers.SplitSlice[int](data, numGoroutines)

	launchWorkers(segments, results, &wg)

	wg.Wait()
	close(results)

	ratios := Ratios{}
	for result := range results {
		ratios.PositiveProportion += result.PositiveProportion
		ratios.NegativeProportion += result.NegativeProportion
		ratios.ZeroProportion += result.ZeroProportion
	}

	fmt.Printf("Proportions - Positives: %.2f, Negatives: %.2f, Zeros: %.2f\n",
		ratios.PositiveProportion, ratios.NegativeProportion, ratios.ZeroProportion)
}
