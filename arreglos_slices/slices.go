package arreglos_slices

import "fmt"

var tablaSlice []int = []int{2, 5, 4}
var arreglos_slices [10]int = [10]int{5, 6, 7, 34, 2, 43, 6, 8, 9, 0}

func MustrasSlice() {
	fmt.Println(tablaSlice)

	porcion := arreglos_slices[2:]  //slice creado desde un vector desde a posicion 2 hasta el final
	porcion2 := arreglos_slices[:6] //slice creado desde un vector desde el inicio hasta la posicion 4
	porcion3 := arreglos_slices[2:5]
	fmt.Println("################################################")
	fmt.Println(porcion)
	fmt.Println("################################################")
	fmt.Println(porcion2)
	fmt.Println("################################################")
	fmt.Println(porcion3)
	fmt.Println("################################################")
}

func Capacidad() {
	elementos := make([]int, 5, 20)

	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	numero := make([]int, 0, 0)

	for i := 0; i < 350; i++ {
		numero = append(numero, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(numero), cap(numero))
}
