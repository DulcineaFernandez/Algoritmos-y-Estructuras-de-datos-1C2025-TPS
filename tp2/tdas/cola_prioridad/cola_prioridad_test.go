package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"
)

func comparacionInt(clave int, clave2 int) int {
	if clave < clave2 {
		return -1
	} else if clave > clave2 {
		return 1
	}

	return 0
}

// HEAP CREADO VACIO
func TestHeapCreado(t *testing.T) {
	t.Log("Prueba que un Heap vacío funcione de forma correcta")
	heap := TDAHeap.CrearHeap[string](strings.Compare)
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

}

func TestHeapUnElemento(t *testing.T) {
	t.Log("Prueba que un Heap con un elemento funcione de forma correcta")

	heap := TDAHeap.CrearHeap[string](strings.Compare)
	dato := "A"
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())

	heap.Encolar(dato)
	require.EqualValues(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, dato, heap.VerMax())

	heap.Desencolar()
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

}

func TestHeapGuardar(t *testing.T) {
	t.Log("Prueba que un Heap encole varios elementos de forma correcta")

	heap := TDAHeap.CrearHeap[string](strings.Compare)
	elementos := []string{"C", "B", "A", "D"}

	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

	heap.Encolar(elementos[0])
	require.EqualValues(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, elementos[0], heap.VerMax())

	heap.Encolar(elementos[1])
	require.EqualValues(t, 2, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, elementos[0], heap.VerMax())

	heap.Encolar(elementos[2])
	require.EqualValues(t, 3, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, elementos[0], heap.VerMax())

	heap.Encolar(elementos[3])
	require.EqualValues(t, 4, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, elementos[3], heap.VerMax())
}

func TestHeapElementosRepetidos(t *testing.T) {
	t.Log("Prueba que un Heap funcione con elementos repetidos")

	heap := TDAHeap.CrearHeap[string](strings.Compare)
	elementos := []string{"A", "B", "C", "A", "A", "B", "C"}
	require.True(t, heap.EstaVacia())

	i := heap.Cantidad()
	for _, elem := range elementos {
		require.EqualValues(t, i, heap.Cantidad())
		heap.Encolar(elem)
		require.False(t, heap.EstaVacia())
		i++
	}

	resultado := []string{"C", "C", "B", "B", "A", "A", "A"}
	for _, rta := range resultado {
		require.False(t, heap.EstaVacia())
		require.EqualValues(t, rta, heap.VerMax())
		require.EqualValues(t, rta, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapVaciar(t *testing.T) {
	t.Log("Prueba que vaciar un Heap con varios elementos no rompa su funcionamiento")

	cant := 30
	heap := TDAHeap.CrearHeap[int](comparacionInt)
	for i := 0; i < cant; i++ {
		require.EqualValues(t, i, heap.Cantidad())
		heap.Encolar(i)
		require.False(t, heap.EstaVacia())
	}

	require.EqualValues(t, cant, heap.Cantidad())
	require.EqualValues(t, cant-1, heap.VerMax())

	for i := cant - 1; i >= 0; i-- {
		require.EqualValues(t, i, heap.Desencolar())
		require.EqualValues(t, i, heap.Cantidad())
	}

	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapVolumenOrdenado(t *testing.T) {

	cant := 10000
	heap := TDAHeap.CrearHeap[int](comparacionInt)
	for i := 0; i < cant; i++ {
		require.EqualValues(t, i, heap.Cantidad())
		heap.Encolar(i)
		require.False(t, heap.EstaVacia())
	}
	require.EqualValues(t, cant, heap.Cantidad())
	require.EqualValues(t, cant-1, heap.VerMax())

	for i := cant - 1; i >= 0; i-- {
		require.EqualValues(t, i, heap.Desencolar())
		require.EqualValues(t, i, heap.Cantidad())
	}
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapVolumenSinRandMilElementos(t *testing.T) {
	t.Log("Prueba de volumen con 10000 elementos desordenados")

	cant := 10000
	valores := make([]int, cant)

	// Llenamos el arreglo con un patrón fijo que "mezcla" los valores
	for i := 0; i < cant; i++ {
		valores[i] = (i * 337) % cant // 337 es coprimo con 1000, garantiza permutación completa
	}

	heap := TDAHeap.CrearHeap[int](comparacionInt)

	for _, val := range valores {
		heap.Encolar(val)
		require.False(t, heap.EstaVacia())
	}

	require.EqualValues(t, cant, heap.Cantidad())
	require.EqualValues(t, cant-1, heap.VerMax())

	for i := cant - 1; i >= 0; i-- {
		require.False(t, heap.EstaVacia())
		require.EqualValues(t, i, heap.Desencolar())
	}

	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestVolumenCrearHeapArrCreciente(t *testing.T) {
	t.Log("Prueba de volumen: CrearHeapArr con arreglo creciente con heap de maximo")

	cant := 10000
	arr := make([]int, cant)
	for i := 0; i < cant; i++ {
		arr[i] = i
	}

	heap := TDAHeap.CrearHeapArr[int](arr, comparacionInt)
	require.EqualValues(t, cant, heap.Cantidad())
	require.EqualValues(t, cant-1, heap.VerMax())

	for i := cant - 1; i >= 0; i-- {
		require.False(t, heap.EstaVacia())
		require.EqualValues(t, i, heap.Desencolar())
	}
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSortVolumen(t *testing.T) {
	t.Log("Prueba de volumen: HeapSort con arreglo desordenado")

	cant := 10000
	arr := make([]int, cant)
	for i := 0; i < cant; i++ {
		arr[i] = (i * 337) % cant
	}

	TDAHeap.HeapSort[int](arr, comparacionInt)

	// Verificamos que está ordenado de menor a mayor
	for i := 1; i < cant; i++ {
		require.LessOrEqual(t, arr[i-1], arr[i])
	}
}

func TestGuardarYBorrar(t *testing.T) {
	t.Log("Prueba guardar y borrar elementos de forma alternada")

	heap := TDAHeap.CrearHeap[string](strings.Compare)
	elementos := []string{"C", "B", "A", "D"}

	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())

	heap.Encolar(elementos[0])
	require.EqualValues(t, elementos[0], heap.VerMax())
	heap.Encolar(elementos[1])
	require.EqualValues(t, elementos[0], heap.VerMax())

	require.EqualValues(t, elementos[0], heap.Desencolar())
	require.EqualValues(t, elementos[1], heap.VerMax())

	heap.Encolar(elementos[2])
	require.EqualValues(t, elementos[1], heap.VerMax())
	heap.Encolar(elementos[3])
	require.EqualValues(t, elementos[3], heap.VerMax())

	require.EqualValues(t, elementos[3], heap.Desencolar())
	require.EqualValues(t, elementos[1], heap.VerMax())
	require.EqualValues(t, elementos[1], heap.Desencolar())
	require.EqualValues(t, elementos[2], heap.VerMax())
	require.EqualValues(t, elementos[2], heap.Desencolar())
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())

	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

// HEAP DESDE ARREGLO
func TestHeapConArregloCrearInt(t *testing.T) {
	t.Log("Prueba que se pueda  crear un heap con un arreglo de enteros")
	elementos := []int{5, 3, 8, 1, 2}
	heap := TDAHeap.CrearHeapArr(elementos, comparacionInt)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, 8, heap.VerMax())
}

func TestHeapConArregloNumerosRepetidosInt(t *testing.T) {
	arr := []int{5, 3, 7, 3, 2, 5, 8, 5, 1}
	TDAHeap.HeapSort[int](arr, comparacionInt)
	for i := 1; i < len(arr); i++ {
		require.LessOrEqual(t, arr[i-1], arr[i])
	}
}

func TestHeapConArregloCrearString(t *testing.T) {
	t.Log("Prueba que se pueda  crear un heap con un arreglo de strings")
	elementos := []string{"5", "3", "8", "1", "2"}
	heap := TDAHeap.CrearHeapArr(elementos, strings.Compare)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, "8", heap.VerMax())
}

func TestHeapConArregloGuardar(t *testing.T) {
	t.Log("Prueba que se pueda guardar datos correctamenteen un heap creado con un arreglo")

	elementos := []int{5, 3, 8, 1, 2}
	heap := TDAHeap.CrearHeapArr(elementos, comparacionInt)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, 8, heap.VerMax())

	diez := 10
	cuatro := 4
	veinte := 20

	heap.Encolar(cuatro)
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, 8, heap.VerMax())

	heap.Encolar(diez)
	require.EqualValues(t, 7, heap.Cantidad())
	require.EqualValues(t, 10, heap.VerMax())

	heap.Encolar(veinte)
	require.EqualValues(t, 8, heap.Cantidad())
	require.EqualValues(t, 20, heap.VerMax())

	heap.Encolar(cuatro)
	require.EqualValues(t, 9, heap.Cantidad())
	require.EqualValues(t, 20, heap.VerMax())

	heap.Encolar(veinte)
	require.EqualValues(t, 10, heap.Cantidad())
	require.EqualValues(t, 20, heap.VerMax())
}

func TestHeapConArregloVaciar(t *testing.T) {
	t.Log("Prueba que se pueda vaciar correctamente un heap creado con un arreglo")
	elementos := []int{5, 3, 8, 1, 2}
	heap := TDAHeap.CrearHeapArr(elementos, comparacionInt)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, 8, heap.VerMax())

	heap.Desencolar()
	heap.Desencolar()
	heap.Desencolar()
	heap.Desencolar()
	heap.Desencolar()

	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapConArregloVacio(t *testing.T) {
	t.Log("Prueba que se pueda crear un heap con un arreglo vacio")
	arregloVacio := make([]int, 0, 0)
	heap := TDAHeap.CrearHeapArr(arregloVacio, comparacionInt)
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapConArregloVacioGuardarYBorrar(t *testing.T) {
	t.Log("Prueba que se pueda crear un heap con un arreglo vacio y ande correctamente")
	arregloVacio := make([]string, 0, 0)
	heap := TDAHeap.CrearHeapArr(arregloVacio, strings.Compare)
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

	elementos := []string{"C", "B", "A", "D"}

	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())

	heap.Encolar(elementos[0])
	require.EqualValues(t, elementos[0], heap.VerMax())
	heap.Encolar(elementos[1])
	require.EqualValues(t, elementos[0], heap.VerMax())

	require.EqualValues(t, elementos[0], heap.Desencolar())
	require.EqualValues(t, elementos[1], heap.VerMax())

	heap.Encolar(elementos[2])
	require.EqualValues(t, elementos[1], heap.VerMax())
	heap.Encolar(elementos[3])
	require.EqualValues(t, elementos[3], heap.VerMax())

	require.EqualValues(t, elementos[3], heap.Desencolar())
	require.EqualValues(t, elementos[1], heap.VerMax())
	require.EqualValues(t, elementos[1], heap.Desencolar())
	require.EqualValues(t, elementos[2], heap.VerMax())
	require.EqualValues(t, elementos[2], heap.Desencolar())
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())

	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

// HEAPSORT
func TestHeapSortInts(t *testing.T) {
	t.Log("Prueba HeapSort con enteros")
	elementos := []int{7, 3, 9, 1, 6, 2}
	resultado := []int{1, 2, 3, 6, 7, 9}
	TDAHeap.HeapSort(elementos, comparacionInt)
	require.EqualValues(t, resultado, elementos)
}

func TestHeapSortLetras(t *testing.T) {
	t.Log("Prueba HeapSort con strings")
	elementos := []string{"D", "B", "A", "C", "E"}
	resultado := []string{"A", "B", "C", "D", "E"}
	TDAHeap.HeapSort(elementos, strings.Compare)
	require.EqualValues(t, resultado, elementos)
}

func TestHeapSortVacio(t *testing.T) {
	t.Log("Prueba HeapSort con arreglo vacío")
	elementos := []int{}
	resultado := []int{}
	TDAHeap.HeapSort(elementos, comparacionInt)
	require.EqualValues(t, resultado, elementos)
}
