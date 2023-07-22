package iteraciones

import (
	"fmt"
)

func Iteraciones() {

	for i := 0; i < 100; i += 5 {
		fmt.Println("Iteracion: ", i)
	}
}
