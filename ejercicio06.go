package main

import (
	"fmt"
	"time"
)

var temp int

func p() {
	t := 0
	n := 0
	for t <= 10 {
		temp = n
		n = temp + 1
		t++
	}

}

func main() {

	go p() // gorotina = instancia = proceso
	go p() // gorotina = instancia = proceso

	time.Sleep(time.Millisecond * 1000)

	fmt.Printf("El valor final de temp es %d\n", temp)
}
