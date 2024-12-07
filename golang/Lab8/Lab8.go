package Lab8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateY(b float64, x float64) float64 {
	return (1 + math.Pow(math.Sin(math.Pow(b, 3)+math.Pow(x, 3)), 2)) / math.Cbrt(math.Pow(b, 3)+math.Pow(x, 3))
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

	createAndWriteFile(filename)

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

	var searchTerm string
	fmt.Print("\nВведите текст для поиска в файле: ")
	fmt.Scanln(&searchTerm)
	searchInFile(filename, searchTerm)
}

func createAndWriteFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Println("Введите значения для записи в файл (введите 'exit' для завершения):")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			break
		}
		writer.WriteString(text + "\n")
	}

	writer.Flush()
	fmt.Println("Данные успешно записаны в файл:", filename)
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

func searchInFile(filename string, searchTerm string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchTerm) {
			fmt.Println("Найдено:", line)
			found = true
		}
	}

	if !found {
		fmt.Println("Текст не найден в файле.")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}
