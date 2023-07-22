package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaDeMultiplicar() string {
	var numero int
	var err error
	var texto string
	fmt.Println("Ingrese un numero: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato es incorrecto" + err.Error())
		}
	}
	for i := 1; i <= 10; i++ {

		texto += fmt.Sprintln(numero, "x", i, "=", numero*i)
	}
	return texto
}
