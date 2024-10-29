package Lab4

import (
	"fmt"
	"math"
	"Lab4"
)

func calculateY(b float64, x float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / math.Cbrt(math.Pow(b, 3)+math.Pow(x, 3))
}

func taskA(bA float64, xStart float64, xEnd float64, deltaX float64) {
	fmt.Println("Результаты задания A:")
	for x := xStart; x <= xEnd; x += deltaX {
		y := calculateY(bA, x)
		fmt.Printf("x: %.2f, y: %.4f\n", x, y)
	}
}

func taskB(bB float64, xValues []float64) {
	fmt.Println("\nРезультаты задания B:")
	for _, x := range xValues {
		y := calculateY(bB, x)
		fmt.Printf("x: %.2f, y: %.4f\n", x, y)
	}
}

func RunLab4() {
	bA := 2.5
	xStart := 1.28
	xEnd := 3.28
	deltaX := 0.4

	taskA(bA, xStart, xEnd, deltaX)

	bB := 2.5
	xValues := []float64{1.1, 2.4, 3.6, 1.7, 3.9}

	taskB(bB, xValues)
}
