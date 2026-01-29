package lista

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type iteradorExterno[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]

	//Para acordarme (despues borrar comentario) : le pse listaItera para no confundir con lista al poner lista.alg
	listaIterar *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	borrado := lista.primero.dato

	if lista.Largo() == 1 {
		lista.ultimo = nil
	}

	lista.primero = lista.primero.siguiente
	lista.largo--

	return borrado
}

func crearNodo[T any](datoNuevo T) *nodoLista[T] {
	return &nodoLista[T]{dato: datoNuevo, siguiente: nil}
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodo(dato)

	nuevoNodo.siguiente = lista.primero
	lista.primero = nuevoNodo

	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}

	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodo(dato)

	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo

	lista.largo++
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoActual := lista.primero

	for nodoActual != nil {
		if !visitar(nodoActual.dato) {
			return
		}

		nodoActual = nodoActual.siguiente
	}

}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorExterno[T]{actual: lista.primero, anterior: nil, listaIterar: lista}
}

func (iterador *iteradorExterno[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iteradorExterno[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	return iterador.actual.dato

}

func (iterador *iteradorExterno[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iteradorExterno[T]) Insertar(dato T) {
	nuevoNodo := crearNodo(dato)

	//Insertar Primero
	if iterador.anterior == nil {

		nuevoNodo.siguiente = iterador.listaIterar.primero
		iterador.listaIterar.primero = nuevoNodo

		if iterador.listaIterar.ultimo == nil {
			iterador.listaIterar.ultimo = nuevoNodo
		}
		// Insertar al final
	} else if !iterador.HaySiguiente() {

		iterador.anterior.siguiente = nuevoNodo
		iterador.listaIterar.ultimo = nuevoNodo
		// en el medio
	} else {

		nuevoNodo.siguiente = iterador.actual
		iterador.anterior.siguiente = nuevoNodo
	}

	iterador.actual = nuevoNodo

	iterador.listaIterar.largo++

}

func (iterador *iteradorExterno[T]) Borrar() T {

	//caso vacia
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")

	}

	borrado := iterador.actual.dato

	//caso de un solo elemento
	if iterador.listaIterar.primero == iterador.listaIterar.ultimo {
		iterador.actual = nil
		iterador.listaIterar.primero = nil
		iterador.listaIterar.ultimo = nil

		//caso borrar primero
	} else if iterador.anterior == nil {

		iterador.listaIterar.primero = iterador.actual.siguiente
		iterador.actual = iterador.actual.siguiente

		//Caso ultimo
	} else if iterador.actual.siguiente == nil {
		iterador.anterior.siguiente = nil
		iterador.listaIterar.ultimo = iterador.anterior
		iterador.actual = nil

		//medio
	} else {
		iterador.actual = iterador.actual.siguiente
		iterador.anterior.siguiente = iterador.actual
	}

	iterador.listaIterar.largo--

	return borrado
}
