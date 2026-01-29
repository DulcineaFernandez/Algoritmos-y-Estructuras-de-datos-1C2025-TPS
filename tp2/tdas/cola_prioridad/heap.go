package cola_prioridad

type heap[T any] struct {
	datos       []T
	cant        int
	funcion_cmp func(T, T) int
}

const (
	CAPACIDAD_INICIAL             = 7
	FACTOR_CARGA_MAX              = 0.99
	MULTIPLICADOR_NUEVA_CAPACIDAD = 2
	FACTOR_CARGA_MINIMO           = 0.25
	NO_REDIMENSIONAR              = -1
	DIVISOR_NUEVA_CAPACIDAD       = 2
)

// CrearHeap, crea un heap vacío.
func CrearHeap[T any](funcion_funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{datos: make([]T, CAPACIDAD_INICIAL), cant: 0, funcion_cmp: funcion_funcion_cmp}
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *heap[T]) Cantidad() int {
	return heap.cant
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

// upheap sube un elemento hasta que encunetra la posición correcta
func upheap[T any](datos []T, posicion int, funcion_cmp func(T, T) int) {
	if posicion == 0 {
		return
	}

	padre := (posicion - 1) / 2
	if funcion_cmp(datos[posicion], datos[padre]) > 0 {
		datos[posicion], datos[padre] = datos[padre], datos[posicion]
		upheap(datos, padre, funcion_cmp)
	}
}

// downheap baja un elemento hasta encontrar su posición correcta en el heap
func downheap[T any](datos []T, posicion int, hasta int, funcion_cmp func(T, T) int) {
	hijoIzq := posicion*2 + 1
	hijoDcho := posicion*2 + 2
	mayor := posicion

	if hijoIzq < hasta && funcion_cmp(datos[hijoIzq], datos[mayor]) > 0 {
		mayor = hijoIzq
	}

	if hijoDcho < hasta && funcion_cmp(datos[hijoDcho], datos[mayor]) > 0 {
		mayor = hijoDcho
	}

	if mayor != posicion {
		datos[posicion], datos[mayor] = datos[mayor], datos[posicion]
		downheap(datos, mayor, hasta, funcion_cmp)
	}
}

// heapify ordena un arreglo de modo que respete la propiedad de heap
func heapify[T any](datos []T, funcion_cmp func(T, T) int) {
	for i := len(datos) - 1; i >= 0; i-- {
		downheap(datos, i, len(datos), funcion_cmp)
	}
}

func (heap *heap[T]) redimensionar(nuevaCapacidad int) {
	nuevo := make([]T, nuevaCapacidad)
	copy(nuevo, heap.datos)
	heap.datos = nuevo

}

func (heap *heap[T]) Encolar(dato T) {
	factorCarga := float64(heap.cant) / float64(cap(heap.datos))
	if factorCarga > FACTOR_CARGA_MAX {
		heap.redimensionar(cap(heap.datos) * MULTIPLICADOR_NUEVA_CAPACIDAD)
	}

	heap.datos[heap.cant] = dato
	heap.cant++
	upheap(heap.datos, heap.cant-1, heap.funcion_cmp)
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	datoPrioritario := heap.VerMax()
	heap.datos[0] = heap.datos[heap.cant-1]
	heap.cant--
	downheap(heap.datos, 0, heap.cant, heap.funcion_cmp)

	factorCarga := float64(heap.cant) / float64(cap(heap.datos))
	if factorCarga < FACTOR_CARGA_MINIMO {
		nuevaCapacidad := cap(heap.datos) / DIVISOR_NUEVA_CAPACIDAD
		if nuevaCapacidad >= CAPACIDAD_INICIAL {
			heap.redimensionar(nuevaCapacidad)
		}
	}

	return datoPrioritario
}

// CrearHeapArr, crea un heap desde un arreglo
func CrearHeapArr[T any](arreglo []T, funcion_funcion_cmp func(T, T) int) ColaPrioridad[T] {
	if len(arreglo) == 0 {
		return CrearHeap(funcion_funcion_cmp)
	}

	copia := make([]T, len(arreglo))
	copy(copia, arreglo)

	heapify(copia, funcion_funcion_cmp)
	return &heap[T]{datos: copia, cant: len(copia), funcion_cmp: funcion_funcion_cmp}
}

// Heapsort ordena el arreglo
func HeapSort[T any](elementos []T, funcion_funcion_cmp func(T, T) int) {
	if len(elementos) == 0 {
		return
	}

	heapify(elementos, funcion_funcion_cmp)
	for i := len(elementos) - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downheap(elementos, 0, i, funcion_funcion_cmp)
	}
}
