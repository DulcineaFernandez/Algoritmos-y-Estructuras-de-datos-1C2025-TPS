package lista

type Lista[T any] interface {

	// EstaVacia devuelve true si la lista no tiene elementos, false en caso contrario
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento como primero de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista y se devuelve ese valor. Si esta vacia, entra en pánico con un mensaje "La lita esta vacia"
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista. Si esta Vacia entre en panico con un mensaje "La lista esta vacia"
	VerPrimero() T

	// VerUltimo obtiene el valor de ultimo de la lista. Si esta vacia entra en panico con un mensaje "La lista esta vacia"
	VerUltimo() T

	// Devuelve la cantidad de elementos de la lista
	Largo() int

	//Recorre los elementos de la lista y visita a la funcion
	Iterar(visitar func(T) bool)

	// Iterador() Crea un iterador para la lista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual obtiene el valor en el cual estoy parado de la lista
	VerActual() T

	//Devuelve true si el iterador actual es != nil y false sino
	HaySiguiente() bool

	//Siguiente pasa al siguiente elemento de la lista
	Siguiente()

	//Insertar agrega un elemento nuevo a la lista en la posición que se encuentra actualmente
	Insertar(T)

	//Borrar() eleimina un elemento de laposicón que se encuentra actualmente en la lista
	Borrar() T
}
