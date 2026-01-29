package cola_test

import (
	"fmt"
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaCreadaEstaVacia(t *testing.T) {
	fmt.Println("PRUEBAS DE CREACIÓN DE PILA")

	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())

	fmt.Println("✅ Una cola recién creada está vacía")

}

func TestColaCreadaPanics(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()

	require.Panics(t, func() { cola.VerPrimero() }, "PANIC: Ver el tope de una cola recien creada debería tirar pánico")
	fmt.Println("✅ No puedo ver el tope de una cola recién creada")

	require.Panics(t, func() { cola.Desencolar() }, "PANIC: Desencolar una cola recien creada debería tirar pánico")
	fmt.Println("✅ No puedo desencolar una cola recién creada")

}

func TestEncolarEnColaCreada(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)

	require.True(t, !cola.EstaVacia())
	require.Equal(t, 1, cola.VerPrimero())

	fmt.Println("✅ Una cola recién creada puede Encolar y deja de estar vacia")

	require.Equal(t, 1, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() }, "PANIC: Ver el primero de una cola recien creada debería tirar pánico")
	require.Panics(t, func() { cola.Desencolar() }, "PANIC: DesEncolar una cola recien creada debería tirar pánico")

	fmt.Println("✅ Luego de Encolar un elemento la cola se comporta correctamente")

	fmt.Print("\n")
}

func TestOrdenFIFO(t *testing.T) {

	fmt.Println("PRUEBA FIFO")

	cola := TDACola.CrearColaEnlazada[string]()
	elemA := "A"
	elemB := "B"
	elemC := "C"

	cola.Encolar(elemA)
	cola.Encolar(elemB)
	cola.Encolar(elemC)

	require.Equal(t, elemA, cola.Desencolar())
	require.Equal(t, elemB, cola.Desencolar())
	require.Equal(t, elemC, cola.Desencolar())

	fmt.Println("✅ Los elementos se apilan y desapilan en orden FIFO")

	fmt.Print("\n")

}

func TestVaciarCola(t *testing.T) {
	fmt.Println("PRUEBAS DE COLA VACÍA")

	cola := TDACola.CrearColaEnlazada[string]()
	elemA := "A"
	elemB := "B"
	elemC := "C"

	cola.Encolar(elemA)
	cola.Encolar(elemB)
	cola.Encolar(elemC)

	elemA = cola.Desencolar()
	elemB = cola.Desencolar()
	elemC = cola.Desencolar()

	require.True(t, cola.EstaVacia())
	fmt.Println("✅ Cuando desencolo todos los elementos de una cola esta queda vacia")

}

func TestColalaVaciaPanics(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[string]()
	elemA := "A"
	elemB := "B"
	elemC := "C"

	cola.Encolar(elemA)
	cola.Encolar(elemB)
	cola.Encolar(elemC)

	elemA = cola.Desencolar()
	elemB = cola.Desencolar()
	elemC = cola.Desencolar()

	require.Panics(t, func() { cola.VerPrimero() }, "PANIC: Ver el primero de una cola vacía debería tirar pánico")
	fmt.Println("✅ No puedo ver el tope de una cola vacía")

	require.Panics(t, func() { cola.Desencolar() }, "PANIC: Desencolar una cola vacía debería tirar pánico")
	fmt.Println("✅ No puedo desencolar una cola vacía")

	fmt.Print("\n")

}

func TestEncolarEnteros(t *testing.T) {
	fmt.Println("CREAR COLA DE ENTEROS")

	cola := TDACola.CrearColaEnlazada[int]()

	primerElemento := 17
	cola.Encolar(primerElemento)
	require.Equal(t, primerElemento, cola.VerPrimero())

	segundoElemento := 31
	cola.Encolar(segundoElemento)
	require.Equal(t, primerElemento, cola.VerPrimero())

	require.True(t, !cola.EstaVacia())
	fmt.Println("✅ Se puede crear y encolar una cola de enteros")

	require.Equal(t, primerElemento, cola.Desencolar())
	require.Equal(t, segundoElemento, cola.Desencolar())
	fmt.Println("✅ Se puede Desencolar una cola de enteros")

	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() }, "PANIC: Ver el tope de una cola vacía debería tirar pánico")
	require.Panics(t, func() { cola.Desencolar() }, "PANIC: Desencolar una cola vacía debería tirar pánico")

	fmt.Println("✅ La cola de enteros se puede encolar y Desencolar sin romperse")

}
func TestEncolarStrings(t *testing.T) {
	fmt.Println("CREAR COLA DE STRINGS")

	cola := TDACola.CrearColaEnlazada[string]()

	primerElemento := "17"
	cola.Encolar(primerElemento)
	require.Equal(t, primerElemento, cola.VerPrimero())

	segundoElemento := "31"
	cola.Encolar(segundoElemento)
	require.Equal(t, primerElemento, cola.VerPrimero())

	require.True(t, !cola.EstaVacia())
	fmt.Println("✅ Se puede crear y encolar una cola de strings")

	require.Equal(t, primerElemento, cola.Desencolar())
	require.Equal(t, segundoElemento, cola.Desencolar())
	fmt.Println("✅ Se puede Desencolar una cola de strings")

	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() }, "PANIC: Ver el tope de una cola vacía debería tirar pánico")
	require.Panics(t, func() { cola.Desencolar() }, "PANIC: Desencolar una cola vacía debería tirar pánico")

	fmt.Println("✅ La cola de strings se puede encolar y Desencolar sin romperse")

}

func TestEncolarSlices(t *testing.T) {
	fmt.Println("CREAR COLA DE SLICES")

	cola := TDACola.CrearColaEnlazada[[]int]()

	primerSlices := []int{1, 2, 3}

	cola.Encolar(primerSlices)
	require.Equal(t, primerSlices, cola.VerPrimero())

	segundoSlices := []int{4, 5, 6}

	cola.Encolar(segundoSlices)
	require.Equal(t, primerSlices, cola.VerPrimero())

	require.True(t, !cola.EstaVacia())
	fmt.Println("✅ Se puede crear y encolar una cola de slices")

	require.Equal(t, primerSlices, cola.Desencolar())
	require.Equal(t, segundoSlices, cola.Desencolar())
	fmt.Println("✅ Se puede Desencolar una cola de slices")

	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() }, "PANIC: Ver el tope de una cola vacía debería tirar pánico")
	require.Panics(t, func() { cola.Desencolar() }, "PANIC: Desencolar una cola vacía debería tirar pánico")

	fmt.Println("✅ La cola de slices se puede encolar y Desencolar sin romperse")

}

func TestEncolarMuchosElementos(t *testing.T) {
	fmt.Println("PRUEBAS DE VOLUMEN")

	cola := TDACola.CrearColaEnlazada[int]()

	for i := 0; i <= 10000; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}

	for i := 0; i <= 10000; i++ {
		require.Equal(t, i, cola.Desencolar())
	}

	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() }, "PANIC: Ver el tope de una cola vacía debería tirar pánico")
	require.Panics(t, func() { cola.Desencolar() }, "PANIC: desencolar una cola vacía debería tirar pánico")

	fmt.Println("✅ Puedo encolar 10000 elementos y se desencolan en el orden correcto ")
	fmt.Println("✅ Puedo encolar y desencolar 10000 elementos y la cola sigue funcionando correctamente ")

	colaMasGrande := TDACola.CrearColaEnlazada[int]()

	for i := 0; i <= 1000000; i++ {
		colaMasGrande.Encolar(i)
		require.Equal(t, 0, colaMasGrande.VerPrimero())
	}

	for i := 0; i <= 1000000; i++ {
		require.Equal(t, i, colaMasGrande.Desencolar())
	}

	require.True(t, colaMasGrande.EstaVacia())
	require.Panics(t, func() { colaMasGrande.VerPrimero() }, "PANIC: Ver el tope de una cola vacía debería tirar pánico")
	require.Panics(t, func() { colaMasGrande.Desencolar() }, "PANIC: desencolar una cola vacía debería tirar pánico")

	fmt.Println("✅ Puedo encolar 1000000 elementos y se desencolan en el orden correcto ")
	fmt.Println("✅ Puedo encolar y desencolar 1000000 elementos y la cola sigue funcionando correctamente ")
}

func TestEncolarYDesencolarIntercalado(t *testing.T) {
	fmt.Println("PRUEBAS DE ENCOLAR Y DESENCOLAR ALTERNADO")

	cola := TDACola.CrearColaEnlazada[int]()
	primerElemento := 1
	cola.Encolar(primerElemento)
	segundoElemento := 2
	cola.Encolar(segundoElemento)

	require.Equal(t, primerElemento, cola.Desencolar())

	tercerElemento := 3
	cola.Encolar(tercerElemento)
	require.Equal(t, segundoElemento, cola.Desencolar())
	require.Equal(t, tercerElemento, cola.VerPrimero())
	require.Equal(t, tercerElemento, cola.Desencolar())

	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() }, "PANIC: Ver el tope de una cola vacía debería tirar pánico")
	require.Panics(t, func() { cola.Desencolar() }, "PANIC: Desencolar una cola vacía debería tirar pánico")

	fmt.Println("✅ Puedo encolar y desencolar alternadamente sin romper la cola")

	fmt.Print("\n")

}
