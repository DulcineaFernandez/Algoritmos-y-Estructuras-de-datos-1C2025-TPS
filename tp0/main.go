package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const FILE1PATH = "archivo1.in"
const FILE2PATH = "archivo2.in"

// Imprime un vector en formato columna
func print_vector(vector []int) {
	for _, valor := range vector {
		fmt.Println(valor)
	}
}

// Rellena un vector con los datos de un archivo
// Debe recibir una ruta de archivo valida
func load_vector(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s\n", err, file)
		return nil, err
	}

	defer file.Close()

	vector := []int{}

	lector := bufio.NewScanner(file)
	for lector.Scan() {
		linea := lector.Text()
		if linea != "" {
			numero, err := strconv.Atoi(linea)
			if err != nil {
				fmt.Println(err)
				continue
			}
			vector = append(vector, numero)
		} else {
			vector = append(vector, 0)
		}

	}

	read_error := lector.Err()
	if read_error != nil {
		return nil, read_error
	}

	return vector, nil
}

func main() {

	numbers_list1, err := load_vector(FILE1PATH)

	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s\n", err, FILE1PATH)
		return
	}

	numbers_list2, err := load_vector(FILE2PATH)

	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s\n", err, FILE2PATH)
		return
	}

	large_array := ejercicios.Comparar(numbers_list1, numbers_list2)


	if large_array == 0 || large_array == 1 {
		ejercicios.Seleccion(numbers_list1)
		print_vector(numbers_list1)
	} else {
		ejercicios.Seleccion(numbers_list2)
		print_vector(numbers_list2)
	}
}
