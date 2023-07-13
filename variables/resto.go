package variables

import (
	"fmt"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var fecha time.Time

func RestoVariables() {
	Nombre = "Duke"
	Estado = true
	Sueldo = 150.50
	fecha = time.Now()

	fmt.Println("Nombre: ", Nombre)
	fmt.Println("Estado: ", Estado)
	fmt.Println("Sueldo: ", Sueldo)
	fmt.Println("Fecha: ", fecha)
}
