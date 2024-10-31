package main

import (
    "fmt"
    "math"
)

func main(){
    bA := 2.5
    xStart := 1.28
    xEnd := 3.28
    deltaX := 0.4

    fmt.Println("Результаты задания A:")
    for x := xStart; x <= xEnd; x += deltaX {
        y := (1 + math.Pow(math.Sin(math.Pow(bA, 3)+math.Pow(x, 3)), 2)) / math.Cbrt(math.Pow(bA, 3)+math.Pow(x, 3))
        fmt.Printf("x: %.2f, y: %.4f\n", x, y)
    }

    bB := 2.5
    xValues := []float64{1.1, 2.4, 3.6, 1.7, 3.9}

    fmt.Println("\nРезультаты задания B:")
    for _, x := range xValues {
        y := (1 + math.Pow(math.Sin(math.Pow(bB, 3)+math.Pow(x, 3)), 2)) / math.Cbrt(math.Pow(bB, 3)+math.Pow(x, 3))
        fmt.Printf("x: %.2f, y: %.4f\n", x, y)
    }
}
