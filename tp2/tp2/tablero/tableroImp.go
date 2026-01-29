package tablero

import (
	"bufio"
	"os"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADicionario "tdas/diccionario"
	TDALista "tdas/lista"
	"time"
	TDAVuelo "tp2/vuelo"
)

const (
	FORMATO     = "2006-01-02T15:04:05"
	ASCENDENTE  = "asc"
	DESCENDENTE = "desc"
	MAX_CODIGO  = "\xff\xff\xff"
)

type tablero struct {
	vuelosPorCódigo     TDADicionario.Diccionario[string, TDAVuelo.Vuelo]
	vuelosPorFecha      TDADicionario.DiccionarioOrdenado[FechaCodigo, TDAVuelo.Vuelo]
	vuelosPorconexiones TDADicionario.Diccionario[TDAVuelo.Conexion, TDALista.Lista[TDAVuelo.Vuelo]]
}

type FechaCodigo struct {
	Fecha  time.Time
	Codigo string
}

// Un valor negativo si el primer argumento es menor que el segundo (es decir, desde < hasta).
// Cero si son iguales.
// Un valor positivo si el primer argumento es mayor que el segundo (desde > hasta).
func CompararFechaCodigo(a, b FechaCodigo) int {
	if a.Fecha.Before(b.Fecha) {
		return -1
	}
	if a.Fecha.After(b.Fecha) {
		return 1
	}
	if a.Codigo < b.Codigo {
		return -1
	}
	if a.Codigo > b.Codigo {
		return 1
	}
	return 0
}

// Parsea la fecha de string a formato time.Time
func parsearFecha(fecha string) time.Time {
	fechaParseada, _ := time.Parse(FORMATO, fecha)
	return fechaParseada
}

// CrearTablero crea un tablero
func CrearTablero() Tablero {
	return &tablero{
		vuelosPorCódigo:     TDADicionario.CrearHash[string, TDAVuelo.Vuelo](),
		vuelosPorFecha:      TDADicionario.CrearABB[FechaCodigo, TDAVuelo.Vuelo](CompararFechaCodigo),
		vuelosPorconexiones: TDADicionario.CrearHash[TDAVuelo.Conexion, TDALista.Lista[TDAVuelo.Vuelo]](),
	}
}

func (tablero *tablero) obtenerListaConexion(conexion TDAVuelo.Conexion) TDALista.Lista[TDAVuelo.Vuelo] {
	if !tablero.vuelosPorconexiones.Pertenece(conexion) {
		tablero.vuelosPorconexiones.Guardar(conexion, TDALista.CrearListaEnlazada[TDAVuelo.Vuelo]())
	}
	return tablero.vuelosPorconexiones.Obtener(conexion)
}

func (tablero *tablero) guardarPorConexion(vuelo TDAVuelo.Vuelo) {
	listaVuelos := tablero.obtenerListaConexion(vuelo.GetConexion())
	iter := listaVuelos.Iterador()
	for iter.HaySiguiente() {
		v := iter.VerActual()
		if v.GetFlightNumber() == vuelo.GetFlightNumber() {
			iter.Borrar()
			iter.Insertar(vuelo)
			return
		}
		iter.Siguiente()
	}

	listaVuelos.InsertarUltimo(vuelo)
}

func (tablero *tablero) borrarDeConexion(vueloViejo TDAVuelo.Vuelo) {
	conexion := vueloViejo.GetConexion()
	if !tablero.vuelosPorconexiones.Pertenece(conexion) {
		return
	}
	lista := tablero.vuelosPorconexiones.Obtener(conexion)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		vuelo := iter.VerActual()
		if vuelo.GetFlightNumber() == vueloViejo.GetFlightNumber() {
			iter.Borrar()
			return
		}
		iter.Siguiente()
	}
}

// guardarVuelo Guarda un vuelo en el tablero, si ya existe lo actualiza
func (tablero *tablero) guardarVuelo(vueloNuevo TDAVuelo.Vuelo) {
	codigo := vueloNuevo.GetFlightNumber()

	if tablero.vuelosPorCódigo.Pertenece(codigo) {
		vueloViejo := tablero.vuelosPorCódigo.Obtener(codigo)
		claveVieja := FechaCodigo{
			Fecha:  vueloViejo.GetDate(),
			Codigo: codigo,
		}
		tablero.vuelosPorFecha.Borrar(claveVieja)
		tablero.borrarDeConexion(vueloViejo)
	}

	tablero.vuelosPorCódigo.Guardar(codigo, vueloNuevo)

	claveNueva := FechaCodigo{
		Fecha:  vueloNuevo.GetDate(),
		Codigo: codigo,
	}
	tablero.vuelosPorFecha.Guardar(claveNueva, vueloNuevo)

	tablero.guardarPorConexion(vueloNuevo)
}

func (tablero *tablero) AgregarArchivo(ruta string) bool {
	archivo, err := os.Open(ruta)
	if err != nil {
		return false
	}
	defer archivo.Close()
	LectorDeVuelos := bufio.NewScanner(archivo)
	for LectorDeVuelos.Scan() {
		linea := LectorDeVuelos.Text()
		infoVuelo := strings.Split(linea, ",")
		vueloNuevo := TDAVuelo.CrearVuelo(infoVuelo)
		tablero.guardarVuelo(vueloNuevo)
	}
	return true
}

func (tablero *tablero) obtenerVuelosPorRango(desde, hasta *time.Time) []TDAVuelo.Vuelo {
	claveDesde := FechaCodigo{Fecha: *desde, Codigo: ""}         // El menor código posible
	claveHasta := FechaCodigo{Fecha: *hasta, Codigo: MAX_CODIGO} // Mayor string posible

	iter := tablero.vuelosPorFecha.IteradorRango(&claveDesde, &claveHasta)
	var vuelos []TDAVuelo.Vuelo

	for iter.HaySiguiente() {
		_, vuelo := iter.VerActual()
		vuelos = append(vuelos, vuelo)
		iter.Siguiente()
	}
	return vuelos

}

func (tablero *tablero) VerTablero(k int, modo string, desde string, hasta string) ([]TDAVuelo.Vuelo, bool) {
	fDesde := parsearFecha(desde)
	fHasta := parsearFecha(hasta)

	if k <= 0 || (modo != ASCENDENTE && modo != DESCENDENTE) || fDesde.After(fHasta) {
		return nil, false
	}

	vuelosLista := tablero.obtenerVuelosPorRango(&fDesde, &fHasta)
	var vuelos []TDAVuelo.Vuelo

	if modo == ASCENDENTE {
		for i := 0; i < k && i < len(vuelosLista); i++ {
			vuelos = append(vuelos, vuelosLista[i])
		}
	} else {
		for i := len(vuelosLista) - 1; i >= 0 && k > 0; i-- {
			vuelos = append(vuelos, vuelosLista[i])
			k--
		}
	}
	return vuelos, true
}

// compararPrioridad Devuelve 1 si el primer vuelo es mas prioritario que el segundo, -1 en caso contrario y 0 si son iguales
func compararPrioridad(vuelo1, vuelo2 TDAVuelo.Vuelo) int {
	p1 := vuelo1.GetPriority()
	p2 := vuelo2.GetPriority()

	if p1 > p2 {
		return 1
	} else if p1 < p2 {
		return -1
	}

	return strings.Compare(vuelo2.GetFlightNumber(), vuelo1.GetFlightNumber())
}

func (tablero *tablero) PrioridadVuelos(k int) ([]TDAVuelo.Vuelo, bool) {
	if k <= 0 {
		return nil, false
	}
	var resultado []TDAVuelo.Vuelo

	desde := time.Time{}
	hasta := time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)
	arregloVuelos := tablero.obtenerVuelosPorRango(&desde, &hasta)
	heapVuelos := TDAHeap.CrearHeapArr(arregloVuelos, compararPrioridad)
	for i := 0; i < k && !heapVuelos.EstaVacia(); i++ {
		resultado = append(resultado, heapVuelos.Desencolar())
	}

	return resultado, true
}

func (tablero *tablero) InfoVuelo(codigo string) (TDAVuelo.Vuelo, bool) {
	vuelo := tablero.vuelosPorCódigo.Obtener(codigo)
	return vuelo, true
}

func (tablero *tablero) SiguienteVuelo(origen string, destino string, fecha string) (TDAVuelo.Vuelo, bool) {
	fechaComienzo := parsearFecha(fecha)
	conexionBuscada := TDAVuelo.CrearConexion(origen, destino)
	listaVuelos := tablero.obtenerListaConexion(*conexionBuscada)
	iter := listaVuelos.Iterador()
	var vueloMasProximo TDAVuelo.Vuelo = nil
	for iter.HaySiguiente() {
		vuelo := iter.VerActual()
		iter.Siguiente()
		if vuelo.GetCancelled() == "1" || vuelo.GetDate().Before(fechaComienzo) {
			continue
		}

		if vueloMasProximo == nil || vuelo.GetDate().Before(vueloMasProximo.GetDate()) {
			vueloMasProximo = vuelo
		}
	}

	if vueloMasProximo == nil {
		return nil, false
	}

	return vueloMasProximo, true
}

func (tablero *tablero) borrarVuelo(vuelo TDAVuelo.Vuelo) {
	clave := FechaCodigo{
		Fecha:  vuelo.GetDate(),
		Codigo: vuelo.GetFlightNumber(),
	}
	tablero.vuelosPorFecha.Borrar(clave)
	tablero.vuelosPorCódigo.Borrar(vuelo.GetFlightNumber())

	conexion := vuelo.GetConexion()
	lista := tablero.vuelosPorconexiones.Obtener(conexion)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		v := iter.VerActual()
		if v.GetFlightNumber() == vuelo.GetFlightNumber() {
			iter.Borrar()
			break
		}
		iter.Siguiente()
	}
	if lista.EstaVacia() {
		tablero.vuelosPorconexiones.Borrar(conexion)
	}
}

func (tablero *tablero) Borrar(desde string, hasta string) ([]TDAVuelo.Vuelo, bool) {
	Desde := parsearFecha(desde)
	Hasta := parsearFecha(hasta)
	if Desde.After(Hasta) {
		return nil, false
	}

	var resultado []TDAVuelo.Vuelo
	vuelosABorrar := tablero.obtenerVuelosPorRango(&Desde, &Hasta)
	for _, vuelo := range vuelosABorrar {
		resultado = append(resultado, vuelo)
		tablero.borrarVuelo(vuelo)
	}

	return resultado, true
}

func (tablero *tablero) Pertenece(codigo string) bool {
	return tablero.vuelosPorCódigo.Pertenece(codigo)
}
