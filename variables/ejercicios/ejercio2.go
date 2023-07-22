package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaDeMultiplicar() {
	var numero int
	var err error
	fmt.Println("Ingrese un numero: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato es incorrecto" + err.Error())
		}
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(numero, "x", i, "=", numero*i)
	}
}
