package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	menu := `
		Bienvenido
		[1] -> Pizza
		[2] -> Empanada
		Seleccione la opcion de comida de tu preferencia:
	`
	fmt.Printf(menu)
	//lectura de datos por la consola o por buffer
	bufferin := bufio.NewReader(os.Stdin) // creo el buffer para el almacenamiento de datos en la consola
	opc, _ := bufferin.ReadString('\n')   // leer desde la consola
	opc = strings.TrimRight(opc, "\r\n")  // elimina por la derecha el retorno de carro y salto de linea

	// el fmt tiene la opcion de ingreso de consola y salida de consola
	// scan es una de las opciones de fmt
	// le profe no recomienda scan porque es poible que se mezclen caraceres porque e lbuffer es mas exacto para hacer la lectura

	opcn, _ := strconv.Atoi(opc) // ignora el error

	switch opcn {
	case 1:
		fmt.Printf("Elejiste Pizza\n")
	case 2:
		fmt.Printf("Elejiste Empanada\n")
	default:
		fmt.Printf("Opcion invalida\n")
	}
	if opcn == 1 || opcn == 2 {
		fmt.Printf("Ingrese la cantidad de su pedido:\n")
	}

	var cant int
	fmt.Scanf("%d", &cant) // scanf para calcular un valor numerico

	var cpizza int = 20
	var cemp int = 5
	//var total int=0
	total := 0

	if opcn == 1 {
		total = cant * cpizza
	} else {
		total = cant * cemp
	}
	fmt.Printf("El importe pagar es: %d\n", total)
}
