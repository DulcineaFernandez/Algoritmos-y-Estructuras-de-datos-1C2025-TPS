package tablero

import (
	TDAVuelo "tp2/vuelo"
)

type Tablero interface {
	// AgregarArchivo Guarda los vuelos del archivo, en el caso que ya aparezcan los actualiza
	// Tambien devuelve un bool indicando false si hubo un error o true caso contrario
	AgregarArchivo(nombreArchivo string) bool

	//VerTablero devuelve un array de vuelos ordenados por fecha según los filtros mandados.
	// Tambien devuelve un bool indicando false si hubo un error o true caso contrario
	VerTablero(k int, modo string, desde string, hasta string) ([]TDAVuelo.Vuelo, bool)

	// InfoVuelo devuelve un vuelo con toda su informacion
	// Tambien devuelve un bool indicando false si hubo un error o true caso contrario
	InfoVuelo(codigo string) (TDAVuelo.Vuelo, bool)

	//PrioridadVuelos devuelve los k vuelos con mayor prioridad, ordenados.
	// Tambien devuelve un bool indicando false si hubo un error o true caso contrario
	PrioridadVuelos(k int) ([]TDAVuelo.Vuelo, bool)

	//SiguienteVuelo devuelve el próximo vuelo directo que conecte ambos aeropuertos (la fecha meas próxima)
	// Tambien devuelve un bool indicando false si hubo un error o true caso contrario
	SiguienteVuelo(origen string, destino string, fecha string) (TDAVuelo.Vuelo, bool)

	// Borrar elemina los vuelos que están en esas fechas y los devuelve en un array de vuelos
	// Tambien devuelve un bool indicando false si hubo un error o true caso contrario
	Borrar(desde string, hasta string) ([]TDAVuelo.Vuelo, bool)

	//Devuelve True si existe un vuelo con ese código, en caso contrario False
	Pertenece(codigo string) bool
}
