package cola

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato      T
	siguiente *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{primero: nil, ultimo: nil}
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

// Crea un nuevo nodo
func crearNodo[T any](datoNuevo T) *nodoCola[T] {
	nuevoNodo := nodoCola[T]{dato: datoNuevo, siguiente: nil}
	return &nuevoNodo
}

func (cola *colaEnlazada[T]) Encolar(dato T) {
	nuevoNodo := crearNodo(dato)

	if cola.EstaVacia() {
		cola.primero = nuevoNodo

	} else {

		cola.ultimo.siguiente = nuevoNodo
	}

	cola.ultimo = nuevoNodo

}

func (cola *colaEnlazada[T]) Desencolar() T {

	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	nodoAux := cola.primero
	cola.primero = cola.primero.siguiente

	if nodoAux.siguiente == nil {
		cola.ultimo = nil
	}

	return nodoAux.dato
}
