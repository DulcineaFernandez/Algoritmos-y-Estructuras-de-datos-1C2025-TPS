package main

import (
	"bufio"
	"fmt"
	"os"
	"tdas/cola"
	"tdas/pila"
	"tp1/operaciones"
)

// imprime todos los elementos de la cola
func imprimirCuenta(cuenta cola.Cola[string]) {
	for !cuenta.EstaVacia() {
		fmt.Printf("%s ", cuenta.Desencolar())
	}

	fmt.Print("\n")
}

// devuelve true si el token es un numero, false en caso contrario
func esUnNumero(token string) bool {
	for _, simbolo := range token {
		if simbolo < '0' || simbolo > '9' {
			return false
		}
	}

	return true
}

// Genera un arreglo con los tokens de la cuenta
func generadorTokens(cuenta string) []string {
	tokens := []string{}
	var numero string

	for _, simbolo := range cuenta {
		if esUnNumero(string(simbolo)) {
			numero += string(simbolo)

		} else {
			if numero != "" {
				tokens = append(tokens, numero)
				numero = ""
			}
			tokens = append(tokens, string(simbolo))
		}
	}

	if numero != "" {
		tokens = append(tokens, numero)
	}

	return tokens
}

// seguirDesapilando devuelve true si hay que desapilar un operador, false en caso contrario
func seguirDesapilando(operadores pila.Pila[string], token string) bool {
	if operadores.EstaVacia() {
		return false
	}

	if operadores.VerTope() == "(" {
		return false
	}

	opTope := operaciones.CrearOperacion(operaciones.Simbolo(operadores.VerTope()))
	opToken := operaciones.CrearOperacion(operaciones.Simbolo(token))

	if opTope.Precedencia() > opToken.Precedencia() {
		return true
	}

	return opTope.Precedencia() == opToken.Precedencia() && opToken.Asociatividad() == operaciones.Izquierda
}

// Devuelve la cuenta en una notación post fija
func convertirPostFija(cuenta string) cola.Cola[string] {
	salida := cola.CrearColaEnlazada[string]()
	operadores := pila.CrearPilaDinamica[string]()
	tokens := generadorTokens(cuenta)

	for _, token := range tokens {
		if token == " " {
			continue
		}

		if esUnNumero(token) {
			salida.Encolar(token)

		} else if token == "(" {
			operadores.Apilar(token)

		} else if token == ")" {
			for operadores.VerTope() != "(" {
				salida.Encolar(operadores.Desapilar())
			}
			operadores.Desapilar()

		} else if operadores.EstaVacia() {
			operadores.Apilar(token)

		} else {
			for seguirDesapilando(operadores, token) {
				salida.Encolar(operadores.Desapilar())
			}
			operadores.Apilar(token)
		}
	}

	for !operadores.EstaVacia() {
		salida.Encolar(operadores.Desapilar())
	}

	return salida
}

func main() {
	lector := bufio.NewScanner(os.Stdin)

	for lector.Scan() {
		linea := lector.Text()
		NotacionPostFija := convertirPostFija(linea)
		imprimirCuenta(NotacionPostFija)
	}

	err := lector.Err()
	if err != nil {
		fmt.Println(err)

	}
}
