package funciones

import "fmt"

func Exponenecia(valor int) {
	if valor > 100000000 {
		return
	}
	fmt.Println(valor)
	Exponenecia(valor * 2)
}
