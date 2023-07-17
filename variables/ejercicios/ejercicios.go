package ejercicios

import (
	"fmt"
	"strconv"
)

func Ejercicio1(valorNumericoText string) (int, string) {
	textoConvertido, error := strconv.Atoi(valorNumericoText)
	var valorTexto string
	if error != nil {
		fmt.Println("Error al convertir la cadena: ", error)
		return 0, "Error"
	}
	if textoConvertido > 100 {
		valorTexto = "El entero es mayor a 100"
	} else {
		valorTexto = "El entero es menor a 100"
	}
	return textoConvertido, valorTexto
}
