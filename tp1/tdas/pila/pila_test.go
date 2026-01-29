package pila_test

import (
	"fmt"
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaCreadaEstaVacia(t *testing.T) {
	fmt.Println("PRUEBAS DE CREACIÓN DE PILA")

	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	fmt.Println("✅ Una pila recién creada está vacía")

}

func TestPilaCreadaPanics(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()

	require.Panics(t, func() { pila.VerTope() }, "PANIC: Ver el tope de una pila recien creada debería tirar pánico")
	fmt.Println("✅ No puedo ver el tope de una pila recién creada")

	require.Panics(t, func() { pila.Desapilar() }, "PANIC: Desapilar una pila recien creada debería tirar pánico")
	fmt.Println("✅ No puedo desapilar una pila recién creada")

}
func TestApilarEnPilaCreada(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)

	require.True(t, !pila.EstaVacia())
	require.Equal(t, 1, pila.VerTope())

	fmt.Println("✅ Una pila recién creada puede apilar y deja de estar vacia")

	require.Equal(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.VerTope() }, "PANIC: Ver el tope de una pila recien creada debería tirar pánico")
	require.Panics(t, func() { pila.Desapilar() }, "PANIC: Desapilar una pila recien creada debería tirar pánico")

	fmt.Println("✅ Luego de apilar un elemento la pila se comporta correctamente")

	fmt.Print("\n")
}

func TestVaciarPila(t *testing.T) {
	fmt.Println("PRUEBAS DE PILA VACÍA")

	pilaAVaciar := TDAPila.CrearPilaDinamica[string]()
	elemA := "A"
	elemB := "B"
	elemC := "C"

	pilaAVaciar.Apilar(elemA)
	pilaAVaciar.Apilar(elemB)
	pilaAVaciar.Apilar(elemC)

	elemC = pilaAVaciar.Desapilar()
	elemB = pilaAVaciar.Desapilar()
	elemA = pilaAVaciar.Desapilar()

	require.True(t, pilaAVaciar.EstaVacia())
	fmt.Println("✅ Cuando desapilo todos los elementos de una pila esta queda vacia")

}

func TestPilaVaciaPanics(t *testing.T) {

	pilaAVaciar := TDAPila.CrearPilaDinamica[string]()
	elemA := "A"
	elemB := "B"
	elemC := "C"

	pilaAVaciar.Apilar(elemA)
	pilaAVaciar.Apilar(elemB)
	pilaAVaciar.Apilar(elemC)

	elemC = pilaAVaciar.Desapilar()
	elemB = pilaAVaciar.Desapilar()
	elemA = pilaAVaciar.Desapilar()

	require.Panics(t, func() { pilaAVaciar.VerTope() }, "PANIC: Ver el tope de una pila vacía debería tirar pánico")
	fmt.Println("✅ No puedo ver el tope de una pila vacía")

	require.Panics(t, func() { pilaAVaciar.Desapilar() }, "PANIC: Desapilar una pila vacía debería tirar pánico")
	fmt.Println("✅ No puedo desapilar una pila vacía")

	fmt.Print("\n")

}

func TestApilarEnteros(t *testing.T) {
	fmt.Println("CREAR PILA DE ENTEROS")

	pilaEnteros := TDAPila.CrearPilaDinamica[int]()

	pilaEnteros.Apilar(17)
	require.Equal(t, 17, pilaEnteros.VerTope())

	pilaEnteros.Apilar(31)
	require.Equal(t, 31, pilaEnteros.VerTope())

	require.True(t, !pilaEnteros.EstaVacia())
	fmt.Println("✅ Se puede crear y apilar una pila de enteros")

	require.Equal(t, 31, pilaEnteros.Desapilar())
	require.Equal(t, 17, pilaEnteros.Desapilar())
	fmt.Println("✅ Se puede desapilar una pila de enteros")

	require.Panics(t, func() { pilaEnteros.VerTope() }, "PANIC: Ver el tope de una pila vacía debería tirar pánico")
	require.Panics(t, func() { pilaEnteros.Desapilar() }, "PANIC: Desapilar una pila vacía debería tirar pánico")

	fmt.Println("✅ La pila de enteros se puede apilar y desapilar sin romperse")

}

func TestApilarStrings(t *testing.T) {
	fmt.Println("CREAR PILA DE STRINGS")

	pilaStrings := TDAPila.CrearPilaDinamica[string]()

	pilaStrings.Apilar("17")
	require.Equal(t, "17", pilaStrings.VerTope())

	pilaStrings.Apilar("31")
	require.Equal(t, "31", pilaStrings.VerTope())

	require.True(t, !pilaStrings.EstaVacia())
	fmt.Println("✅ Se puede crear y apilar una pila de Strings")

	require.Equal(t, "31", pilaStrings.Desapilar())
	require.Equal(t, "17", pilaStrings.Desapilar())
	fmt.Println("✅ Se puede desapilar una pila de Strings")

	require.Panics(t, func() { pilaStrings.VerTope() }, "PANIC: Ver el tope de una pila vacía debería tirar pánico")
	require.Panics(t, func() { pilaStrings.Desapilar() }, "PANIC: Desapilar una pila vacía debería tirar pánico")

	fmt.Println("✅ La pila de Strings se puede apilar y desapilar sin romperse")

}

func TestApilarSlices(t *testing.T) {
	fmt.Println("CREAR PILA DE SLICE")

	pilaSlices := TDAPila.CrearPilaDinamica[[]int]()
	slice1 := []int{1, 2, 3}

	pilaSlices.Apilar(slice1)
	require.Equal(t, slice1, pilaSlices.VerTope())

	slice2 := []int{4, 5, 6}
	pilaSlices.Apilar(slice2)
	require.Equal(t, slice2, pilaSlices.VerTope())

	require.True(t, !pilaSlices.EstaVacia())
	fmt.Println("✅ Se puede crear y apilar una pila de Slices")

	require.Equal(t, slice2, pilaSlices.Desapilar())
	require.Equal(t, slice1, pilaSlices.Desapilar())
	fmt.Println("✅ Se puede desapilar una pila de Slices")

	require.Panics(t, func() { pilaSlices.VerTope() }, "PANIC: Ver el tope de una pila vacía debería tirar pánico")
	require.Panics(t, func() { pilaSlices.Desapilar() }, "PANIC: Desapilar una pila vacía debería tirar pánico")

	fmt.Println("✅ La pila de Slices se puede apilar y desapilar sin romperse")

}

func TestOrdenLIFO(t *testing.T) {

	fmt.Println("PRUEBA LIFO")

	pila := TDAPila.CrearPilaDinamica[string]()
	elemA := "A"
	elemB := "B"
	elemC := "C"

	pila.Apilar(elemA)
	pila.Apilar(elemB)
	pila.Apilar(elemC)

	require.Equal(t, elemC, pila.Desapilar())
	require.Equal(t, elemB, pila.Desapilar())
	require.Equal(t, elemA, pila.Desapilar())

	fmt.Println("✅ Los elementos se apilan y desapilan en orden LIFO")

	fmt.Print("\n")

}

func TestApilarMuchosElementos(t *testing.T) {
	fmt.Println("PRUEBAS DE VOLUMEN")

	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i <= 10000; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}

	for j := 10000; j >= 0; j-- {
		require.Equal(t, j, pila.Desapilar())
	}

	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.VerTope() }, "PANIC: Ver el tope de una pila vacía debería tirar pánico")
	require.Panics(t, func() { pila.Desapilar() }, "PANIC: Desapilar una pila vacía debería tirar pánico")

	fmt.Println("✅ Puedo apilar 10000 elementos y se desapilan en el orden correcto ")
	fmt.Println("✅ Puedo apilar y desapilar 10000 elementos y la pila sigue funcionando correctamente ")

	pilaMasGrande := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i <= 1000000; i++ {
		pilaMasGrande.Apilar(i)
		require.Equal(t, i, pilaMasGrande.VerTope())

	}

	for j := 1000000; j >= 0; j-- {
		require.Equal(t, j, pilaMasGrande.Desapilar())
	}

	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pilaMasGrande.VerTope() }, "PANIC: Ver el tope de una pila vacía debería tirar pánico")
	require.Panics(t, func() { pilaMasGrande.Desapilar() }, "PANIC: Desapilar una pila vacía debería tirar pánico")

	fmt.Println("✅ Puedo apilar 1000000 elementos y se desapilan en el orden correcto")
	fmt.Println("✅ Puedo apilar y desapilar 1000000 elementos y la pila sigue funcionando correctamente ")
	fmt.Print("\n")
}

func TestApilarYDesapilarIntercalado(t *testing.T) {
	fmt.Println("PRUEBAS DE APILAR Y DESAPILAR ALTERNADO")

	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)

	require.Equal(t, 2, pila.Desapilar())

	pila.Apilar(3)
	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 1, pila.VerTope())
	require.Equal(t, 1, pila.Desapilar())

	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.VerTope() }, "PANIC: Ver el tope de una pila vacía debería tirar pánico")
	require.Panics(t, func() { pila.Desapilar() }, "PANIC: Desapilar una pila vacía debería tirar pánico")

	fmt.Println("✅ Puedo apilar y desapilar alternadamente sin romper la pila")

	fmt.Print("\n")

}
