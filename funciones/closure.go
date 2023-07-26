package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {
	tablaDel2 := 2
	MiTabla := tabla(tablaDel2)
	for i := 0; i < 10; i++ {
		fmt.Println(MiTabla())
	}
	// fmt.Println("Tabla del 2", MiTabla())
}
