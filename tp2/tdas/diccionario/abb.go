package diccionario

import (
	TDAPila "tdas/pila"
)

type nodo[K comparable, V any] struct {
	par     par[K, V]
	abbIzq  *nodo[K, V]
	abbDcho *nodo[K, V]
}

type abb[K comparable, V any] struct {
	raiz        *nodo[K, V]
	comparacion func(K, K) int
	cantidad    int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	arbol := &abb[K, V]{
		raiz:        nil,
		comparacion: funcion_cmp,
		cantidad:    0,
	}
	return arbol
}

func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func crearNodo[K comparable, V any](clave K, valor V) *nodo[K, V] {

	nuevoNodo := &nodo[K, V]{
		par: par[K, V]{
			clave: clave,
			valor: valor,
		},
		abbIzq:  nil,
		abbDcho: nil,
	}

	return nuevoNodo
}

func buscarPosicion[K comparable, V any](arbol *abb[K, V], nodo **nodo[K, V], clave K) (**nodo[K, V], bool) {
	for *nodo != nil {
		comparacion := arbol.comparacion(clave, (*nodo).par.clave)
		if comparacion < 0 {
			nodo = &(*nodo).abbIzq
		} else if comparacion > 0 {
			nodo = &(*nodo).abbDcho
		} else {
			return nodo, true // aca lo encontro, devuelve un puntero al nodo y un true encontrado
		}
	}
	return nodo, false // aca no, entonces nodo apunta a donde deberia guardarse (es solo para guardar esta parte)
	//devuelve un puntero a nil, pero el nil que toca para guardar, y un false de encontrado
}

func (arbol *abb[K, V]) Guardar(clave K, valor V) {
	nodoPosicion, existe := buscarPosicion(arbol, &arbol.raiz, clave)
	if existe {
		(*nodoPosicion).par.valor = valor
	} else {
		*nodoPosicion = crearNodo(clave, valor)
		arbol.cantidad++
	}
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {
	_, nodoEncontrado := buscarPosicion(arbol, &arbol.raiz, clave)
	return nodoEncontrado
}

func (arbol *abb[K, V]) Obtener(clave K) V {
	nodoPosicion, nodoEncontrado := buscarPosicion(arbol, &arbol.raiz, clave)
	if !nodoEncontrado {
		panic("La clave no pertenece al diccionario")
	}
	return (*nodoPosicion).par.valor
}

// borrarMaxYDevolverPar borra un nodo con 2 hijos, estrategia: mayor entre los menores
func borrarMaxYDevolverPar[K comparable, V any](nodo *nodo[K, V]) (*nodo[K, V], *par[K, V]) {
	if nodo.abbDcho == nil {
		return nodo.abbIzq, &nodo.par
	}
	var parMax *par[K, V]
	nodo.abbDcho, parMax = borrarMaxYDevolverPar(nodo.abbDcho)
	return nodo, parMax
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	nodoPosicion, nodoEncontrado := buscarPosicion(arbol, &arbol.raiz, clave)
	if !nodoEncontrado {
		panic("La clave no pertenece al diccionario")
	}

	nodo := *nodoPosicion
	valorBorrado := nodo.par.valor
	arbol.cantidad--

	// Caso 0 o 1 hijo
	if nodo.abbIzq == nil {
		*nodoPosicion = nodo.abbDcho
		return valorBorrado
	}
	if nodo.abbDcho == nil {
		*nodoPosicion = nodo.abbIzq
		return valorBorrado
	}

	// Caso 2 hijos: reemplazar por el mayor entre los menores
	var reemplazo *par[K, V]
	nodo.abbIzq, reemplazo = borrarMaxYDevolverPar(nodo.abbIzq)
	nodo.par = *reemplazo
	*nodoPosicion = nodo

	return valorBorrado
}

func (arbol *abb[K, V]) Iterar(visitar func(clave K, valor V) bool) {
	arbol.IterarRango(nil, nil, visitar)
}

func inorderPorRango[K comparable, V any](arbol *abb[K, V], nodo *nodo[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}

	if desde == nil || arbol.comparacion(nodo.par.clave, *desde) > 0 {
		if !inorderPorRango(arbol, nodo.abbIzq, desde, hasta, visitar) {
			return false
		}
	}

	if (desde == nil || arbol.comparacion(nodo.par.clave, *desde) >= 0) &&
		(hasta == nil || arbol.comparacion(nodo.par.clave, *hasta) <= 0) {
		if !visitar(nodo.par.clave, nodo.par.valor) {
			return false
		}
	}

	if hasta == nil || arbol.comparacion(nodo.par.clave, *hasta) < 0 {
		if !inorderPorRango(arbol, nodo.abbDcho, desde, hasta, visitar) {
			return false
		}
	}

	return true
}

func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	inorderPorRango(arbol, arbol.raiz, desde, hasta, visitar)
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return arbol.IteradorRango(nil, nil)
}

type iterABBporRangos[K comparable, V any] struct {
	pila         TDAPila.Pila[*nodo[K, V]]
	desde, hasta *K
	comparar     func(K, K) int
}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodo[K, V]]()
	apilarRangoInicial(arbol.raiz, pila, desde, hasta, arbol.comparacion)

	return &iterABBporRangos[K, V]{
		pila: pila, desde: desde, hasta: hasta, comparar: arbol.comparacion,
	}
}
func apilarRangoInicial[K comparable, V any](nodo *nodo[K, V], pila TDAPila.Pila[*nodo[K, V]], desde, hasta *K, comparar func(K, K) int) {
	for nodo != nil {
		if desde != nil && comparar(nodo.par.clave, *desde) < 0 {
			nodo = nodo.abbDcho
		} else if hasta != nil && comparar(nodo.par.clave, *hasta) > 0 {
			nodo = nodo.abbIzq
		} else {
			pila.Apilar(nodo)
			nodo = nodo.abbIzq
		}
	}
}

func (iter *iterABBporRangos[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterABBporRangos[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.VerTope()
	return nodo.par.clave, nodo.par.valor
}

func (iter *iterABBporRangos[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	nodo := iter.pila.Desapilar()
	hijo := nodo.abbDcho
	apilarRangoInicial(hijo, iter.pila, iter.desde, iter.hasta, iter.comparar)
}
