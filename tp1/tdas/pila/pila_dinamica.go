package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const (
	CANTIDAD_INICIAL   = 0
	CAPACIDAD_INICIAL  = 10
	FACTOR_CRECIMIENTO = 2
	FACTOR_REDUCTOR    = 2
	LIMITE_REDUCIR     = 4
	CAPACIDAD_MINIMA   = 1
)

func CrearPilaDinamica[T any]() Pila[T] {
	arreglo := make([]T, CAPACIDAD_INICIAL)
	return &pilaDinamica[T]{datos: arreglo, cantidad: CANTIDAD_INICIAL}
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {

	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	return pila.datos[pila.cantidad-1]
}

// Redimensiona la capacidad del arreglo
func redimensionar[T any](pila *pilaDinamica[T], nuevaCapacidad int) {
	var nuevoArreglo []T

	nuevoArreglo = make([]T, nuevaCapacidad)
	copy(nuevoArreglo, pila.datos[:pila.cantidad])
	pila.datos = nuevoArreglo

}

func (pila *pilaDinamica[T]) Apilar(elemento T) {

	if pila.cantidad == cap(pila.datos) {

		if cap(pila.datos) == 0 {
			redimensionar(pila, CAPACIDAD_MINIMA)
		} else {
			redimensionar(pila, cap(pila.datos)*FACTOR_CRECIMIENTO)
		}

	}

	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {

	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	elemento := pila.VerTope()

	pila.cantidad--

	if pila.cantidad*LIMITE_REDUCIR <= cap(pila.datos) {
		redimensionar(pila, cap(pila.datos)/FACTOR_REDUCTOR)
	}
	return elemento
}
