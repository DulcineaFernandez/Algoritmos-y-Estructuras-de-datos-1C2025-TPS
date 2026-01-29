package comandos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDATablero "tp2/tablero"
	TDAVuelo "tp2/vuelo"
)

const (
	AGREGAR                = "agregar_archivo"
	INFOVUELOS             = "info_vuelo"
	VERTABLERO             = "ver_tablero"
	PRIORIDADVUELOS        = "prioridad_vuelos"
	SIGUIENTEVUELO         = "siguiente_vuelo"
	BORRAR                 = "borrar"
	VUELO_COMPLETO         = "completo"
	VUELO_PRIORIDAD_CODIGO = "prioridad"
	VUELO_TABLERO          = "tablero"
	FORMATO                = "2006-01-02T15:04:05"
)

func Leer(tablero TDATablero.Tablero) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := scanner.Text()
		comando := strings.Split(linea, " ")

		switch comando[0] {
		case AGREGAR:
			ok := tablero.AgregarArchivo(comando[1])
			if !ok {
				fmt.Fprintln(os.Stderr, "Error en comando agregar_archivo")
				continue
			}
			fmt.Println("OK")
		case INFOVUELOS:
			if len(comando) != 2 || !tablero.Pertenece(comando[1]) {
				fmt.Fprintln(os.Stderr, "Error en comando info_vuelo")
				continue
			}
			vuelo, ok := tablero.InfoVuelo(comando[1])
			if !ok {
				fmt.Fprintln(os.Stderr, "Error en comando info_vuelo")
			}
			imprimirVuelo(vuelo, VUELO_COMPLETO)
			fmt.Println("OK")
		case PRIORIDADVUELOS:
			cantidad, err := strconv.Atoi(comando[1])
			if err != nil {
				continue
			}
			vuelos, ok := tablero.PrioridadVuelos(cantidad)
			if !ok {
				fmt.Fprintln(os.Stderr, "Error en comando prioridad_vuelos")
				continue
			}
			for _, vuelo := range vuelos {
				imprimirVuelo(vuelo, VUELO_PRIORIDAD_CODIGO)
			}
			fmt.Println("OK")
		case VERTABLERO:
			cantidad, err := strconv.Atoi(comando[1])
			if len(comando) != 5 || err != nil {
				fmt.Fprintln(os.Stderr, "Error en comando ver_tablero")
				continue
			}

			vuelos, ok := tablero.VerTablero(cantidad, comando[2], comando[3], comando[4])
			if !ok {
				fmt.Fprintln(os.Stderr, "Error en comando ver_tablero")
				continue
			}
			for _, vuelo := range vuelos {
				imprimirVuelo(vuelo, VUELO_TABLERO)
			}
			fmt.Println("OK")
		case SIGUIENTEVUELO:
			if len(comando) != 4 {
				fmt.Fprintln(os.Stderr, "Error en comando siguiente_vuelo")
				continue
			}
			origen := comando[1]
			destino := comando[2]
			fecha := comando[3]

			vuelo, ok := tablero.SiguienteVuelo(origen, destino, fecha)
			if !ok {
				fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fecha)
				fmt.Println("OK")
				continue
			}
			imprimirVuelo(vuelo, VUELO_COMPLETO)
			fmt.Println("OK")
		case BORRAR:
			if len(comando) != 3 {
				fmt.Fprintln(os.Stderr, "Error en comando borrar")
				continue
			}
			vuelos, ok := tablero.Borrar(comando[1], comando[2])
			if !ok {
				fmt.Fprintln(os.Stderr, "Error en comando borrar")
				continue
			}
			for _, vuelo := range vuelos {
				imprimirVuelo(vuelo, VUELO_COMPLETO)
			}
			fmt.Println("OK")
		default:
			fmt.Fprintln(os.Stdout, "Unrecognized command:", comando[0])
		}
	}
}

func imprimirVuelo(vuelo TDAVuelo.Vuelo, condicion string) {
	fecha := vuelo.GetDate().Format(FORMATO)
	if condicion == VUELO_COMPLETO {
		horaInt, _ := strconv.Atoi(vuelo.GetDepartureDelay())

		fmt.Printf("%s %s %s %s %s %d %s %d %s %s\n",
			vuelo.GetFlightNumber(),
			vuelo.GetAirLine(),
			vuelo.GetOrigin(),
			vuelo.GetDestino(),
			vuelo.GetTailNumber(),
			vuelo.GetPriority(),
			fecha,
			horaInt,
			vuelo.GetAirTime(),
			vuelo.GetCancelled(),
		)
	} else if condicion == VUELO_PRIORIDAD_CODIGO {
		fmt.Printf("%d - %s\n",
			vuelo.GetPriority(),
			vuelo.GetFlightNumber(),
		)
	} else {
		fmt.Printf("%s - %s\n",
			fecha,
			vuelo.GetFlightNumber(),
		)
	}
}
