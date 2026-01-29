package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

//TEST PARA LISTA EN GENERAL

func TestListaCreadaEstaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
}

func TestListsCreadaPanics(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

}

func TestInsertarPrimeroLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarPrimero(2)
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 2, lista.Largo())
	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	lista.InsertarPrimero(4)
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 4, lista.BorrarPrimero())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia(), "La lista esta vacia si el largo es 0")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")

}

func TestInsertarUltimoLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
	lista.InsertarUltimo(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarUltimo(2)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 2, lista.Largo())
	lista.InsertarUltimo(3)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	lista.InsertarUltimo(4)
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 4, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía si el largo no es 0")

	require.Equal(t, 4, lista.BorrarPrimero())
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia(), "La lista esta vacia si el largo es 0")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")
}

func TestVaciarLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia al crearse")
	require.Equal(t, 0, lista.Largo())

	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())

	lista.InsertarPrimero(2)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())

	lista.InsertarUltimo(3)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())

	lista.InsertarUltimo(5)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 4, lista.Largo())

	lista.InsertarUltimo(6)
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 6, lista.VerUltimo())
	require.Equal(t, 5, lista.Largo())

	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 2, lista.BorrarPrimero())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 6, lista.VerUltimo())
	require.Equal(t, 4, lista.Largo())

	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 6, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())

	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 5, lista.VerPrimero())
	require.Equal(t, 6, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())

	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 5, lista.BorrarPrimero())
	require.Equal(t, 6, lista.VerPrimero())
	require.Equal(t, 6, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())

	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
	require.Equal(t, 6, lista.BorrarPrimero())
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia si el largo es 0")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")

}

func TestListarEnteros(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarUltimo(0)
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, 0, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia(), "La lista debería estar vacía si el largo es 0")
	require.Equal(t, 0, lista.Largo())

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")

}

func TestListarStrings(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
	lista.InsertarPrimero("uno")
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, "uno", lista.VerPrimero())
	require.Equal(t, "uno", lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	lista.InsertarUltimo("dos")
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, "uno", lista.VerPrimero())
	require.Equal(t, "dos", lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())

	require.Equal(t, "uno", lista.BorrarPrimero())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, "dos", lista.VerPrimero())
	require.Equal(t, "dos", lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, "dos", lista.BorrarPrimero())
	require.True(t, lista.EstaVacia(), "La lista debería estar vacía si el largo es 0")
	require.Equal(t, 0, lista.Largo())

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")
}

func TestListarMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	const n = 10000
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia al crearse")
	require.Equal(t, 0, lista.Largo())
	for i := 0; i < n; i++ {
		lista.InsertarPrimero(i)
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, 0, lista.VerUltimo())
		require.Equal(t, i+1, lista.Largo())
	}

	require.Equal(t, n, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")

	for i := lista.Largo() - 1; i > 0; i-- {
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
		require.Equal(t, i, lista.BorrarPrimero())
		require.Equal(t, i-1, lista.VerPrimero())
		require.Equal(t, 0, lista.VerUltimo())
		require.Equal(t, i, lista.Largo())
	}
	require.Equal(t, 0, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia(), "La lista debería estar vacía después de borrar todos los elementos")
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")

	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia al crearse")
	require.Equal(t, 0, lista.Largo())
	for i := 0; i < n; i++ {
		lista.InsertarUltimo(i)
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
		require.Equal(t, 0, lista.VerPrimero())
		require.Equal(t, i, lista.VerUltimo())
		require.Equal(t, i+1, lista.Largo())
	}

	require.Equal(t, n, lista.Largo())
	require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")

	for i := 0; i < n-1; i++ {
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia luego de agregar un elemento")
		require.Equal(t, i, lista.BorrarPrimero())
		require.Equal(t, i+1, lista.VerPrimero())
		require.Equal(t, n-1, lista.VerUltimo())
		require.Equal(t, n-i-1, lista.Largo())
	}

	require.Equal(t, n-1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia al crearse")
	require.Equal(t, 0, lista.Largo())

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")

}

func TestInsertarYBorrarIntercalado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía al crearse")
	require.Equal(t, 0, lista.Largo())
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.Largo())

	lista.InsertarUltimo(5)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 2, lista.Largo())

	lista.InsertarPrimero(7)
	require.Equal(t, 7, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 7, lista.BorrarPrimero())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 2, lista.Largo())

	lista.InsertarUltimo(2)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	lista.InsertarPrimero(3)
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 4, lista.Largo())

	require.Equal(t, 3, lista.BorrarPrimero())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 5, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 2, lista.Largo())

	require.Equal(t, 5, lista.BorrarPrimero())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía luego de agregar un elemento")
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, 2, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía si su largo es 0")
	require.Equal(t, 0, lista.Largo())

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	}, "Ver primero en una lista vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	}, "Ver ultimo en una pila vacía debería causar panic")

	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	}, "Borrar una lista vacia deberia causar panic")

}

//TEST PARA ITERADORES:

//EXTERNO

func TestIteradorExtCreado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente(), "HaySiguiente debe ser false en una lista vacía")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")
}

func TestIterarExtIterarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente(), "El iterador debería estar vacío en una lista recién creada")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

}

func TestIteradorExtInsertarYBorrarUno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	iter.Insertar(42)

	require.False(t, lista.EstaVacia(), "La lista no debería estar vacía después de insertar")
	require.Equal(t, 1, lista.Largo(), "La lista debería tener un elemento")
	require.True(t, iter.HaySiguiente(), "El iterador debería tener un siguiente después de insertar")
	require.Equal(t, 42, iter.VerActual(), "El elemento actual debería ser el que insertamos")

	valorBorrado := iter.Borrar()
	require.Equal(t, 42, valorBorrado, "El valor borrado debería ser el que insertamos")

	require.True(t, lista.EstaVacia(), "La lista debería estar vacía después de borrar")
	require.False(t, iter.HaySiguiente(), "No debería haber siguiente después de borrar el único elemento")
}

func TestIteradorExtInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	iter.Insertar(10)
	require.Equal(t, 10, iter.VerActual(), "El actual debería ser el primer elemento insertado")
	require.Equal(t, 1, lista.Largo(), "La lista debería tener un elemento tras la primera inserción")

	iter.Insertar(20)
	require.Equal(t, 20, iter.VerActual(), "El actual debería ser el último insertado (nuevo primero)")
	require.Equal(t, 2, lista.Largo(), "La lista debería tener dos elementos después de la segunda inserción")

	elementos := []int{}
	iter2 := lista.Iterador()
	for iter2.HaySiguiente() {
		elementos = append(elementos, iter2.VerActual())
		iter2.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	require.Equal(t, []int{20, 10}, elementos, "Los elementos deberían estar en orden de inserción al principio")
}

func TestIteradorExtInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)

	iter := lista.Iterador()

	for iter.HaySiguiente() {
		iter.Siguiente()
	}

	iter.Insertar(40)
	require.Equal(t, 40, iter.VerActual(), "El actual debería ser el nuevo elemento insertado al final")
	require.Equal(t, 4, lista.Largo(), "La lista debería tener cuatro elementos después de la inserción")

	elementos := []int{}
	iter2 := lista.Iterador()
	for iter2.HaySiguiente() {
		elementos = append(elementos, iter2.VerActual())
		iter2.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	require.Equal(t, []int{10, 20, 30, 40}, elementos, "Los elementos deberían estar en el orden correcto después de insertar al final")
}

func TestIteradorExtInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(10)
	lista.InsertarUltimo(30)

	iter := lista.Iterador()

	require.True(t, iter.HaySiguiente(), "Debería haber un primer elemento")
	require.Equal(t, 10, iter.VerActual(), "El primer elemento debería ser 10")

	iter.Siguiente()
	require.True(t, iter.HaySiguiente(), "Debería haber un segundo elemento")
	require.Equal(t, 30, iter.VerActual(), "El segundo elemento debería ser 30")

	iter.Insertar(20)
	require.Equal(t, 20, iter.VerActual(), "El actual debería ser el 20 recién insertado")

	iter.Siguiente()
	require.True(t, iter.HaySiguiente(), "Debería haber un segundo elemento")
	require.Equal(t, 30, iter.VerActual(), "El segundo elemento debería ser 30")

	require.Equal(t, 3, lista.Largo(), "La lista debería tener tres elementos después de la inserción")

	elementos := []int{}
	iter2 := lista.Iterador()
	for iter2.HaySiguiente() {
		elementos = append(elementos, iter2.VerActual())
		iter2.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	require.Equal(t, []int{10, 20, 30}, elementos, "El orden de los elementos debería ser 10, 20, 30")
}

func TestIteradorExtIteracionCompleta(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	elementosEsperados := []int{1, 2, 3, 4, 5}
	for _, elem := range elementosEsperados {
		lista.InsertarUltimo(elem)
	}

	iter := lista.Iterador()

	obtenidos := []int{}
	for iter.HaySiguiente() {
		obtenidos = append(obtenidos, iter.VerActual())
		iter.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	require.Equal(t, elementosEsperados, obtenidos, "La lista debería ser recorrida en el orden de inserción")
	require.False(t, iter.HaySiguiente(), "El iterador debería estar al final después de recorrer toda la lista")

}

func TestIteradorExtBorrarPrimerElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	elementosIniciales := []int{10, 20, 30}
	for _, elem := range elementosIniciales {
		lista.InsertarUltimo(elem)
	}

	require.Equal(t, 3, lista.Largo(), "La lista debería tener tres elementos inicialmente")

	iter := lista.Iterador()
	require.True(t, iter.HaySiguiente(), "El iterador debería estar en el primer elemento")
	require.Equal(t, 10, iter.VerActual(), "El primer elemento debería ser 10")

	valorBorrado := iter.Borrar()
	require.Equal(t, 10, valorBorrado, "El elemento borrado debería ser 10")

	require.Equal(t, 20, iter.VerActual(), "Luego de borrar, el nuevo primer elemento debería ser 20")
	require.Equal(t, 2, lista.Largo(), "La lista debería tener dos elementos después de borrar uno")

	elementosRestantes := []int{}
	iter2 := lista.Iterador()
	for iter2.HaySiguiente() {
		elementosRestantes = append(elementosRestantes, iter2.VerActual())
		iter2.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	require.Equal(t, []int{20, 30}, elementosRestantes, "La lista debería contener los elementos restantes en orden")
}

func TestIteradorExtBorrarUltimoElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	elementosIniciales := []int{1, 2, 3}
	for _, elem := range elementosIniciales {
		lista.InsertarUltimo(elem)
	}

	require.Equal(t, 3, lista.Largo(), "La lista debería tener tres elementos inicialmente")

	iter := lista.Iterador()

	for iter.HaySiguiente() {
		if iter.VerActual() == 3 {
			break
		}
		iter.Siguiente()
	}

	require.Equal(t, 3, iter.VerActual(), "El iterador debería estar en el último elemento (3)")
	valorBorrado := iter.Borrar()
	require.Equal(t, 3, valorBorrado, "El valor borrado debería ser 3")

	require.False(t, iter.HaySiguiente(), "No debería haber más elementos después de borrar el último")
	require.Equal(t, 2, lista.Largo(), "La lista debería tener dos elementos después de borrar el último")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	iter2 := lista.Iterador()
	elementosRestantes := []int{}
	for iter2.HaySiguiente() {
		elementosRestantes = append(elementosRestantes, iter2.VerActual())
		iter2.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	require.Equal(t, []int{1, 2}, elementosRestantes, "La lista debería contener los elementos restantes en orden")
}

func TestIteradorExtBorrarElementoMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	elementosIniciales := []int{1, 2, 3, 4}
	for _, elem := range elementosIniciales {
		lista.InsertarUltimo(elem)
	}

	require.Equal(t, 4, lista.Largo(), "La lista debería tener cuatro elementos inicialmente")

	iter := lista.Iterador()

	for iter.HaySiguiente() {
		if iter.VerActual() == 2 {
			break
		}
		iter.Siguiente()
	}

	require.Equal(t, 2, iter.VerActual(), "El iterador debería estar en el segundo elemento (2)")
	valorBorrado := iter.Borrar()
	require.Equal(t, 2, valorBorrado, "El valor borrado debería ser 2")

	require.Equal(t, 3, iter.VerActual(), "Después de borrar, el iterador debería estar en el siguiente elemento (3)")
	require.Equal(t, 3, lista.Largo(), "La lista debería tener tres elementos después de borrar uno")

	iter2 := lista.Iterador()
	elementosRestantes := []int{}
	for iter2.HaySiguiente() {
		elementosRestantes = append(elementosRestantes, iter2.VerActual())
		iter2.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	require.Equal(t, []int{1, 3, 4}, elementosRestantes, "La lista debería contener los elementos correctos después del borrado")
}

func TestIteradorExtOperacionesCombinadasInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	iter := lista.Iterador()

	iter.Insertar(10)
	require.True(t, iter.HaySiguiente(), "Debería haber siguiente antes de VerActual()")
	require.Equal(t, 10, iter.VerActual())
	iter.Insertar(20)
	require.True(t, iter.HaySiguiente(), "Debería haber siguiente antes de VerActual()")
	require.Equal(t, 20, iter.VerActual())
	iter.Insertar(30)
	require.True(t, iter.HaySiguiente(), "Debería haber siguiente antes de VerActual()")
	require.Equal(t, 30, iter.VerActual())
	require.Equal(t, 3, lista.Largo(), "La lista debería tener tres elementos")
	valorBorrado := iter.Borrar()
	require.Equal(t, 30, valorBorrado, "Debería borrarse el 30")
	require.True(t, iter.HaySiguiente(), "Debería haber siguiente antes de VerActual()")
	require.Equal(t, 20, iter.VerActual(), "Ahora el actual debería ser 20")

	iter.Siguiente()
	require.Equal(t, 10, iter.VerActual(), "Debería estar en 10 ahora")
	iter.Insertar(15)

	require.Equal(t, 15, iter.VerActual(), "Ahora debería estar en 15")
	iter.Siguiente()
	require.Equal(t, 10, iter.VerActual(), "Ahora debería estar en 10")

	valorBorrado = iter.Borrar()
	require.Equal(t, 10, valorBorrado, "Se debería haber borrado el 10")
	require.False(t, iter.HaySiguiente(), "No debería haber siguiente luego de borrar el ultimo")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	iter2 := lista.Iterador()
	var elementos []int
	for iter2.HaySiguiente() {
		elementos = append(elementos, iter2.VerActual())
		iter2.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	require.Equal(t, []int{20, 15}, elementos, "La lista debería terminar con [20, 15]")
}

func TestIteradorExtOperacionesCombinadasString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	iter := lista.Iterador()

	iter.Insertar("perro")
	require.Equal(t, "perro", iter.VerActual())
	iter.Insertar("gato")
	require.Equal(t, "gato", iter.VerActual())
	iter.Insertar("pez")
	require.Equal(t, "pez", iter.VerActual())

	require.Equal(t, 3, lista.Largo(), "La lista debería tener tres elementos")

	require.Equal(t, "pez", iter.VerActual(), "El actual debería ser 'pez'")
	iter.Siguiente()
	require.Equal(t, "gato", iter.VerActual(), "El actual debería ser 'gato'")

	iter.Insertar("loro")
	require.Equal(t, 4, lista.Largo(), "La lista debería tener cuatro elementos")
	require.Equal(t, "loro", iter.VerActual(), "El actual debería ser 'loro' después de insertar")
	valorBorrado := iter.Borrar()
	require.Equal(t, "loro", valorBorrado, "Debería haberse borrado 'loro'")
	require.Equal(t, 3, lista.Largo(), "La lista debería tener tres elementos")
	require.Equal(t, "gato", iter.VerActual(), "Después de borrar debería estar en 'gato'")

	iter.Siguiente()
	require.Equal(t, "perro", iter.VerActual(), "Debería estar en 'perro'")
	valorBorrado = iter.Borrar()
	require.Equal(t, "perro", valorBorrado, "Debería haberse borrado 'perro'")
	require.Equal(t, 2, lista.Largo(), "La lista debería tener dos elementos")

	iter2 := lista.Iterador()
	var elementos []string
	for iter2.HaySiguiente() {
		elementos = append(elementos, iter2.VerActual())
		iter2.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	require.Equal(t, []string{"pez", "gato"}, elementos, "La lista final debería ser ['pez', 'gato']")
}

func TestIteradorExtVolumenInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	iter := lista.Iterador()

	const n = 10000
	for i := 0; i < n; i++ {
		iter.Insertar(i)
	}
	require.Equal(t, n, lista.Largo(), "La lista debería tener n elementos")
	iter2 := lista.Iterador()
	for i := n - 1; iter2.HaySiguiente() == true; i-- {
		require.Equal(t, i, iter2.VerActual(), "El valor actual debería ser el índice durante la iteración")
		iter2.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter2.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	for iter.HaySiguiente() {
		iter.Borrar()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.VerActual()
	}, "VerActual en un iterador vacío debería causar panic")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Borrar()
	}, "Borrar en un iterador vacío debería causar panic")

	require.Equal(t, 0, lista.Largo(), "La lista debería estar vacía después de borrar todos los elementos")
}

//INTERNO

func TestIterarIntIterarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	var result []int
	lista.Iterar(func(valor int) bool {
		result = append(result, valor)
		return true
	})

	require.Empty(t, result, "La iteración sobre una lista vacía no debería agregar elementos")
}

func TestIteradorIntIterarTodos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)

	var result []int
	lista.Iterar(func(valor int) bool {
		result = append(result, valor)
		return true
	})

	require.Equal(t, []int{3, 2, 1}, result, "La iteración debería visitar los elementos en el orden esperado")
}

func TestIteradorStringIterarTodos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarPrimero("1")
	lista.InsertarPrimero("2")
	lista.InsertarPrimero("3")

	var result []string
	lista.Iterar(func(valor string) bool {
		result = append(result, valor)
		return true
	})

	require.Equal(t, []string{"3", "2", "1"}, result, "La iteración debería visitar los elementos en el orden esperado")
}

func TestIterarIntIterarConCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)

	var result []int
	lista.Iterar(func(valor int) bool {
		result = append(result, valor)
		return valor != 2 // Cortar cuando el valor sea 2
	})

	require.Equal(t, []int{3, 2}, result, "La iteración debería detenerse cuando se encuentra el valor 2")
}
