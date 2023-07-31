package mapas

import "fmt"

func MostrarMpas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	paises["Colombia"] = "Bogota"
	paises["Venezuela"] = "Caracas"
	paises["Chile"] = "Santiago"

	// fmt.Println(paises)
	// fmt.Println(paises["Mexico"])

	campeonato := map[string]int{
		"barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	fmt.Println(campeonato)

	// for equipo, puntaje := range campeonato {
	// 	fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", equipo, puntaje)
	// }

	delete(campeonato, "Real Madrid")

	fmt.Println(campeonato)

	puntaje, existe := campeonato["JUVENTUS"]

	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)

	puntaje2, existe2 := campeonato["Chivas"]

	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje2, existe2)
}
