package Lab8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculateY(b float64, x float64) float64 {
	denominator := math.Cbrt(math.Pow(b, 3) + math.Pow(x, 3))
	if denominator == 0 {
		return math.Inf(1)
	}
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / denominator
}

func taskA(bA float64, xStart float64, xEnd float64, deltaX float64) []float64 {
	results := []float64{}
	for x := xStart; x <= xEnd; x += deltaX {
		y := calculateY(bA, x)
		results = append(results, y)
	}
	return results
}

func taskB(bB float64, xValues []float64) []float64 {
	results := []float64{}
	for _, x := range xValues {
		y := calculateY(bB, x)
		results = append(results, y)
	}
	return results
}

func RunLab8() {
	filename := "input.txt"

	values, err := readFile(filename)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	if len(values) < 2 {
		fmt.Println("Недостаточно значений в файле")
		return
	}

	bA := values[0]
	bB := values[1]
	xValues := values[2:]

	xStart := 1.28
	xEnd := 3.28
	deltaX := 0.4

	resultsA := taskA(bA, xStart, xEnd, deltaX)
	fmt.Println("Результаты задания A:")
	for i, y := range resultsA {
		x := xStart + float64(i)*deltaX
		fmt.Printf("x: %.2f, y: %.4f\n", x, y)
	}

	resultsB := taskB(bB, xValues)
	fmt.Println("\nРезультаты задания B:")
	for i, y := range resultsB {
		x := xValues[i]
		fmt.Printf("x: %.2f, y: %.4f\n", x, y)
	}
}

func readFile(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var values []float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Ошибка при чтении числа:", err)
			continue
		}
		values = append(values, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return values, nil
}
