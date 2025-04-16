package main

import (
	"fmt"
	"time"
)

//tercer intento

var wantp bool = false
var wantq bool = false

func p() {
	for {
		fmt.Println("Linea 1: SNC -> P")
		fmt.Println("Linea 2: SNC -> P")
		wantp = true
		for wantq != false {
			//esperar hasta que wantq sea igual a false
		}

		fmt.Println("Linea 1: SC -> P")
		fmt.Println("Linea 2: SC -> P")
		wantp = false // Post Protocol
	}
}

func q() {
	for {
		fmt.Println("Linea 1: SNC -> Q")
		fmt.Println("Linea 2: SNC -> Q")
		wantq = true
		for wantp != false {
			//esperar hasta que wantp sea igual a false
		}
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
