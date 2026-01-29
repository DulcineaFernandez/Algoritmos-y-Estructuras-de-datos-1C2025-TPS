package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const CAPACIDAD_INICIAL = 7

const FACTOR_CARGA_MAX = 2.5
const FACTOR_CARGA_MINIMO = 0.25
const MULTIPLICADOR_NUEVA_CAPACIDAD = 2
const DIVISOR_NUEVA_CAPACIDAD = 2

const NO_REDIMENSIONAR = -1

type par[K comparable, V any] struct {
	clave K
	valor V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[par[K, V]]
	cantidad int
}

// crearTabla, crea la tabla abierta para el hash
func crearTabla[K comparable, V any](capacidad int) []TDALista.Lista[par[K, V]] {
	tabla := make([]TDALista.Lista[par[K, V]], capacidad)
	for i := 0; i < capacidad; i++ {
		tabla[i] = TDALista.CrearListaEnlazada[par[K, V]]()
	}

	return tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tablaNueva := crearTabla[K, V](CAPACIDAD_INICIAL)
	return &hashAbierto[K, V]{tabla: tablaNueva, cantidad: 0}
}

func (diccionario *hashAbierto[K, V]) Cantidad() int {
	return diccionario.cantidad
}

// convertirABytes convierte una clave en un arreglo de bytes
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// funcionDeHash realiza una función de hash para la clave
func funcionDeHash(bytes []byte, capacidad int) int {
	var hash uint64 = 5381

	for i := 0; i < len(bytes); i++ {
		hash = ((hash << 5) + hash) + uint64(bytes[i]) // hash * 33 + c
	}

	return int(hash % uint64(capacidad))
}

// obtenerIndice obtiene la posicion de la tabla para la clave
func obtenerIndice[K comparable](clave K, capacidad int) int {
	claveDeHash := convertirABytes(clave)
	return funcionDeHash(claveDeHash, capacidad)
}

// Redimensionar redimensiona el diccionario
func redimensionar[K comparable, V any](diccionario *hashAbierto[K, V], nuevaCapacidad int) {
	tablaNueva := crearTabla[K, V](nuevaCapacidad)

	for _, lista := range diccionario.tabla {
		iter := lista.Iterador()

		for iter.HaySiguiente() {
			parActual := iter.VerActual()
			indice := obtenerIndice(parActual.clave, nuevaCapacidad)
			tablaNueva[indice].InsertarUltimo(parActual)
			iter.Siguiente()
		}
	}

	diccionario.tabla = tablaNueva
}

// buscarEnLista itera una lista buscando donde esta la clave mandada, si la encuentra devuelve el iterador y true, en caso contrario false
func buscarEnLista[K comparable, V any](diccionario *hashAbierto[K, V], clave K) (TDALista.IteradorLista[par[K, V]], bool) {
	indice := obtenerIndice(clave, cap(diccionario.tabla))
	iter := diccionario.tabla[indice].Iterador()

	for iter.HaySiguiente() {
		if iter.VerActual().clave == clave {
			return iter, true
		}
		iter.Siguiente()
	}
	return iter, false
}

func (diccionario *hashAbierto[K, V]) Guardar(clave K, valor V) {
	factorCarga := float64(diccionario.cantidad) / float64(cap(diccionario.tabla))
	if factorCarga > FACTOR_CARGA_MAX {
		redimensionar(diccionario, cap(diccionario.tabla)*MULTIPLICADOR_NUEVA_CAPACIDAD)
	}

	iter, pertenece := buscarEnLista(diccionario, clave)
	if pertenece {
		iter.Borrar()
		iter.Insertar(par[K, V]{clave: clave, valor: valor})

	} else {
		iter.Insertar(par[K, V]{clave: clave, valor: valor})
		diccionario.cantidad++
	}
}

func (diccionario *hashAbierto[K, V]) Pertenece(clave K) bool {
	_, pertenece := buscarEnLista(diccionario, clave)
	return pertenece
}

func (diccionario *hashAbierto[K, V]) Obtener(clave K) V {
	iter, pertenece := buscarEnLista(diccionario, clave)
	if pertenece {
		return iter.VerActual().valor
	}

	panic("La clave no pertenece al diccionario")
}

func (diccionario *hashAbierto[K, V]) Borrar(clave K) V {
	iter, pertenece := buscarEnLista(diccionario, clave)
	if !pertenece {
		panic("La clave no pertenece al diccionario")
	}

	valorBorrado := iter.VerActual().valor
	iter.Borrar()
	diccionario.cantidad--

	factorCarga := float64(diccionario.cantidad) / float64(cap(diccionario.tabla))
	if factorCarga < FACTOR_CARGA_MAX {
		nuevaCapacidad := cap(diccionario.tabla) / DIVISOR_NUEVA_CAPACIDAD
		if nuevaCapacidad >= CAPACIDAD_INICIAL {
			redimensionar(diccionario, nuevaCapacidad)
		}
	}

	return valorBorrado
}

func (diccionario *hashAbierto[K, V]) Iterar(visitar func(clave K, valor V) bool) {
	for _, lista := range diccionario.tabla {
		if lista.EstaVacia() {
			continue
		}

		iter := lista.Iterador()
		for iter.HaySiguiente() {
			clave := iter.VerActual().clave
			valor := iter.VerActual().valor
			if !visitar(clave, valor) {
				return
			}
			iter.Siguiente()
		}
	}
}

type iterHashAbierto[K comparable, V any] struct {
	diccionario *hashAbierto[K, V]
	indiceTabla int
	iterLista   TDALista.IteradorLista[par[K, V]]
}

// avanzarASiguienteListaNoVacia avanza al siguiente índice no vacío
func (iter *iterHashAbierto[K, V]) avanzarASiguienteListaNoVacia() {
	for iter.indiceTabla < len(iter.diccionario.tabla) {
		listaActual := iter.diccionario.tabla[iter.indiceTabla]
		if !listaActual.EstaVacia() {
			iter.iterLista = listaActual.Iterador()
			return
		}

		iter.indiceTabla++
	}

	iter.iterLista = TDALista.CrearListaEnlazada[par[K, V]]().Iterador()
}

func (diccionario *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHashAbierto[K, V]{
		diccionario: diccionario,
		indiceTabla: 0,
		iterLista:   nil,
	}

	iter.avanzarASiguienteListaNoVacia()
	return iter
}

func (iter *iterHashAbierto[K, V]) HaySiguiente() bool {
	return iter.iterLista != nil && iter.iterLista.HaySiguiente()
}

func (iter *iterHashAbierto[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	parActual := iter.iterLista.VerActual()
	return parActual.clave, parActual.valor
}

func (iter *iterHashAbierto[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	iter.iterLista.Siguiente()
	if !iter.HaySiguiente() {
		iter.indiceTabla++
		iter.avanzarASiguienteListaNoVacia()
	}
}
