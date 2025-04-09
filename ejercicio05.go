// algoritmo concurrente
package main

import (
	"fmt"
	"time"
)

var n int // var global

func p() {
	k1 := 1
	n = k1
}

func q() {
	k2 := 2
	n = k2
}
func main() { // proceso mandatorio (m)
	// secuencial
	go p()
	go q()
	//tiempo d retardo para el proceso main
	time.Sleep(time.Millisecond * 1000)
	// el valor final de n
	fmt.Printf("El valor final de n es : %d\n", n)
}
