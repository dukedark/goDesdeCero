package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dukedark/goDesdeCero/variables/ejercicios"
)

var tablaTexto string = ejercicios.TablaDeMultiplicar()
var fileName string = "./files/txt/tabla.txt"

func FileSave() {
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Hubo un error", err.Error())
		return
	}
	fmt.Fprintln(archivo, tablaTexto)
	archivo.Close()
}

func AddTabla() {
	if !append(fileName, tablaTexto) {
		fmt.Println("Error en la escritura Suma tabla")
	}

}

func append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error en append", err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error en la escritura", err.Error())
		return false
	}
	arch.Close()
	return true
}

func FileRead() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Hubo un error al leer el arvhivo", err.Error())
		return
	}
	fmt.Println(string(archivo))
}

func FileRead2() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Hubo un error al leer el archivo", err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("Linea > " + registro)
	}
	archivo.Close()
}
