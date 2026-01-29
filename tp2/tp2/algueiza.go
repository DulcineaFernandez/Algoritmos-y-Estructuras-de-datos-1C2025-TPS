package main

import (
	programa "tp2/comandos"
	TDATablero "tp2/tablero"
)

func main() {
	tablero := TDATablero.CrearTablero()
	programa.Leer(tablero)
}
