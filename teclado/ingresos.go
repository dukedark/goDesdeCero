package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numeroUno int
var numeroDos int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println("Ingrese numero uno")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numeroUno, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese numero dos")
	if scanner.Scan() {
		numeroDos, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numeroUno*numeroDos)
}
