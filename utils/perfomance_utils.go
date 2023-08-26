package utils

import (
	"log"
	"time"
)

const performanceFormat = "performance.%s - took %s"

func MeasurePerformance(nameOperation string, fn func()) {
	start := time.Now()

	fn()
	elapsed := time.Since(start)
	log.Printf(performanceFormat, nameOperation, elapsed.String())
}
