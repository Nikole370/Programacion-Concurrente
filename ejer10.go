package main

import (
	"fmt"
	"time"
)

//Segundo intento de solucion de problema de SC -> Deadlock

var wantp bool = false
var wantq bool = false

func p() {
	for {
		fmt.Println("Linea 1: SNC -> P")
		fmt.Println("Linea 2: SNC -> P")
		for wantq != false {
			//esperar hasta que wantq sea igual a false
		}
		wantp = true
		fmt.Println("Linea 1: SC -> P")
		fmt.Println("Linea 2: SC -> P")
		wantp = false // Post Protocol
	}
}

func q() {
	for {
		fmt.Println("Linea 1: SNC -> Q")
		fmt.Println("Linea 2: SNC -> Q")
		for wantp != false {
			//esperar hasta que wantp sea igual a false
		}
		wantq = true
		fmt.Println("Linea 1: SC -> P")
		fmt.Println("Linea 2: SC -> P")
		wantq = false // Post Condicion   o Post Protocol
	}
}

func main() {

	go p()
	go q()

	time.Sleep(time.Millisecond * 100)

}
