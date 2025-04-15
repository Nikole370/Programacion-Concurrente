package main

import (
	"fmt"
	"time"
)

var (
	n = 0     // Variable compartida, inicialmente 0
	K = 10000 // Número de incrementos/decrementos que hará cada gorutina
)

// p() incrementa n K veces
func p() {
	for i := 0; i < K; i++ {
		temp := n    // temp <- n
		n = temp + 1 // n <- temp + 1
	}
}

// q() decrementa n K veces
func q() {
	for i := 0; i < K; i++ {
		temp := n    // temp <- n
		n = temp - 1 // n <- temp - 1
	}
}

func main() {
	// Lanzamos ambas funciones en paralelo
	go p()
	go q()

	// Damos tiempo a que ambas gorutinas terminen sus bucles
	//time.Sleep(1 * time.Second)
	time.Sleep(time.Millisecond * 1000)

	// Imprimimos el resultado final
	fmt.Printf("El valor final de n es %d\n", n)
}
