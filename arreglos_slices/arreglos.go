package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{3, 5, 10, 7, 99, 0, 0, 0, 4}
var matriz [20][30]int

func MustraArreglos() {
	tabla[2] = 33
	tabla[8] = 15

	tabla2 := [15]string{"pedro", "juan", "", "Sol", "", "luna"}
	fmt.Println(tabla)
	fmt.Println("Tabla", tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)

}
