package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Preguntar valor mínimo
	fmt.Print("Ingrese el valor mínimo: ")
	min, err := strconv.ParseFloat(getInput(), 64)
	if err != nil {
		fmt.Println("Error: Valor mínimo incorrecto")
		return
	}

	// Preguntar valor máximo
	fmt.Print("Ingrese el valor máximo: ")
	max, err := strconv.ParseFloat(getInput(), 64)
	if err != nil {
		fmt.Println("Error: Valor máximo incorrecto")
		return
	}

	// Verificar que el mínimo sea menor que el máximo
	if min >= max {
		fmt.Println("Error: El mínimo debe ser menor que el máximo")
		return
	}

	// Obtener la lista de valores
	fmt.Print("Ingrese la lista de valores (separados por espacios): ")
	valuesInput := getInput()

	// Convertir la entrada en números
	var values []float64
	for _, str := range strings.Fields(valuesInput) {
		if num, err := strconv.ParseFloat(str, 64); err == nil {
			values = append(values, num)
		}
	}

	// Filtrar los valores dentro del rango
	resultado := minmax(min, max, values...)

	// Mostrar el resultado
	fmt.Printf("Los valores dentro del rango son: %v\n", resultado)
}

func minmax(min, max float64, values ...float64) []float64 {
	var dentroRango []float64
	for _, valor := range values {
		if valor >= min && valor <= max {
			dentroRango = append(dentroRango, valor)
		}
	}
	return dentroRango
}

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
