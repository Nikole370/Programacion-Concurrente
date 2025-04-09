package main

import "fmt"

func main() {
	// arreglos / estructura repetitiva
	arreglo := [7]int{5, 3, 4, 1, 6, 2, 0}

	// 1. forma de uso dl for (estructura repetitiva)
	for i, v := range arreglo {
		fmt.Printf("El valor de v es %d en el indice %d\n", v, i)
	}

	// 2 . forma
	// similar al while(1)--> infinito
	fmt.Printf("Uso del for infinito\n")
	i := 5
	for {
		fmt.Printf("Valor=%d\n", i)
		i--
		if i < 0 {
			break
		}
	}

	//3. forma
	// Usando condicion
	fmt.Printf("Uso de for --> condicional\n")

	j := 0
	for j <= 6 {
		fmt.Printf("Valor = %d\n", j)
		j++
	}

}
