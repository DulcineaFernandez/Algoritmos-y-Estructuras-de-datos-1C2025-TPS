package diccionario_test

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

//DE LAS PRUEBAS QUE NO SON DE ITERADORES ME FAETAN LAS DE BORRAR PORQUE NO TERMINE DE HACER BORRAR.
//LAS PRUEBAS LAS COPIE DE MARTIN. NO SE SI SE PODIA HABRIA QU CONSULTAR. PERO TAMPOCO ES TAN DIFICIL HACERLAS. VOY A HACER UN ARCHIVO APARTE DE PRUEBAS POR LAS DUDAS.

// este no lo entendixd
func TestTImprimirMensaje(t *testing.T) {
	fmt.Println("TEST DICCIONARIO ORDENADO")
}

// este funca? tipo no se si se comparan asi los strings en go
func comparacionString(clave string, clave2 string) int {
	if clave < clave2 {
		return -1
	} else if clave > clave2 {
		return 1
	}

	return 0
}

func comparacionInt(clave int, clave2 int) int {
	if clave < clave2 {
		return -1
	} else if clave > clave2 {
		return 1
	}

	return 0
}

// ESTO LOS PASO

func TestDiccionarioOrdenadoCreado(t *testing.T) {
	t.Log("Diccionario Ordenado Vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.Panics(t, func() { dic.Obtener("A") }, "PANIC: La clave no pertenece al diccionario")
	require.Panics(t, func() { dic.Borrar("A") }, "PANIC: La clave no pertenece al diccionario")

	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](comparacionInt)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })

}

// ESTA LAS PASA
func TestDiccionarioOrdenadoUnElemento(t *testing.T) {
	t.Log("Probar funcionamiento de un elemento en el diccionario")
	dic := TDADiccionario.CrearABB[string, int](comparacionString)
	require.EqualValues(t, 0, dic.Cantidad())

	dic.Guardar("A", 1)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 1, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })

	dic.Borrar("A")
	require.EqualValues(t, 0, dic.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.Panics(t, func() { dic.Obtener("A") }, "PANIC: La clave no pertenece al diccionario")
	require.Panics(t, func() { dic.Borrar("A") }, "PANIC: La clave no pertenece al diccionario")

}

// ESTA LAS PASA
func TestDiccionarioOrdenadoGuardar(t *testing.T) {

	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](comparacionString)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestDccionarioOrdenadoReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

//ACÁ van las de borrado con cantidad, no las hice porque me falta el caso d dos hijos

// pasa esta prueba
func TestDiccionarioOrdenadoClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

// pasa esta prueba
func TestDiccionarioOrdenadoValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](comparacionString)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestDiccionarioOrdenadoCadenaLargaParticular(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionario.CrearABB[string, string](comparacionString)
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func TestDiccionarioOrdenadoGuardarYBorrarRepetidasVeces(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces. Esto lo hacemos porque un error comun es no considerar " +
		"los borrados para agrandar en un Hash Cerrado. Si no se agranda, muy probablemente se quede en un ciclo " +
		"infinito")

	dic := TDADiccionario.CrearABB[int, int](comparacionInt)
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}

func TestDiccionarioVacioo(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](func(a, b string) int {
		return strings.Compare(a, b)
	})
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

//Borrar

func TestBorrarElementosABB(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](func(a, b string) int {
		return strings.Compare(a, b)
	})

	dic.Guardar("b", 2)
	dic.Guardar("a", 1)
	dic.Guardar("c", 3)

	require.Equal(t, 3, dic.Cantidad())

	// Borrar hoja
	require.EqualValues(t, 1, dic.Borrar("a"))
	require.False(t, dic.Pertenece("a"))

	// Borrar nodo con un hijo
	require.EqualValues(t, 3, dic.Borrar("c"))
	require.False(t, dic.Pertenece("c"))

	// Borrar nodo con dos hijos
	require.EqualValues(t, 2, dic.Borrar("b"))
	require.False(t, dic.Pertenece("b"))

	require.Equal(t, 0, dic.Cantidad())
}

// ////
// Iteradores
func TestIteradorExternoABBVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
}

func TestIteradorExternoABBUnElemento(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("clave", 10)

	iter := dic.Iterador()
	require.True(t, iter.HaySiguiente())

	clave, valor := iter.VerActual()
	require.Equal(t, "clave", clave)
	require.Equal(t, 10, valor)

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestIteradorExternoABBInOrder(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	valores := []int{50, 30, 70, 20, 40, 60, 80}

	for _, v := range valores {
		dic.Guardar(v, fmt.Sprintf("valor-%d", v))
	}

	iter := dic.Iterador()
	esperado := []int{20, 30, 40, 50, 60, 70, 80}
	i := 0

	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		require.Equal(t, esperado[i], clave)
		require.Equal(t, fmt.Sprintf("valor-%d", esperado[i]), valor)
		iter.Siguiente()
		i++
	}
	require.Equal(t, len(esperado), i)

}

func TestIteradorExternoRecorridoInorderABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)

	for _, clave := range []string{"d", "b", "a", "c", "e"} {
		abb.Guardar(clave, 1)
	}

	iter := abb.Iterador()
	var recorrido []string
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		recorrido = append(recorrido, clave)
		iter.Siguiente()
	}

	require.Equal(t, []string{"a", "b", "c", "d", "e"}, recorrido)
}

func TestIteradorVerActualPanic(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()

	require.Panics(t, func() {
		iter.VerActual()
	})
}

func TestIteradorSiguientePanic(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()

	require.Panics(t, func() {
		iter.Siguiente()
	})
}

func TestIteradorExternoRangoABB(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)

	for _, clave := range []string{"a", "b", "c", "d", "e", "f"} {
		abb.Guardar(clave, 1)
	}

	desde := "c"
	hasta := "e"

	iter := abb.IteradorRango(&desde, &hasta)
	var clavesRango []string
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clavesRango = append(clavesRango, clave)
		iter.Siguiente()
	}

	require.Equal(t, []string{"c", "d", "e"}, clavesRango)

}

//INTERNO

func TestIteradorInternoInorderABB(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	dic.Guardar("b", "valor_b")
	dic.Guardar("a", "valor_a")
	dic.Guardar("c", "valor_c")

	var recorrido []string
	dic.Iterar(func(clave string, valor string) bool {
		recorrido = append(recorrido, clave)
		return true
	})

	require.Equal(t, []string{"a", "b", "c"}, recorrido)
}

func TestIteradorInternoRangoABB(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	for _, clave := range []string{"a", "b", "c", "d", "e"} {
		dic.Guardar(clave, 1)
	}

	desde := "b"
	hasta := "d"

	var clavesDentro []string
	dic.IterarRango(&desde, &hasta, func(clave string, valor int) bool {
		clavesDentro = append(clavesDentro, clave)
		return true
	})

	require.Equal(t, []string{"b", "c", "d"}, clavesDentro)
}
