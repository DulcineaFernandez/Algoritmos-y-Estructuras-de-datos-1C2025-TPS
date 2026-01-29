package vuelo

import (
	"strconv"
	"time"
)

type vuelo struct {
	flightNumber    string
	airline         string
	conexion        Conexion
	tailNumber      string
	priority        string
	date            time.Time
	deperatureDelay string
	airTime         string
	cancelled       string
}

type Conexion struct {
	originAirport      string
	destinationAirport string
}

func CrearVuelo(infoVuelo []string) Vuelo {
	conexionNueva := Conexion{
		originAirport:      infoVuelo[2],
		destinationAirport: infoVuelo[3],
	}

	layout := "2006-01-02T15:04:05"
	fechaNueva, _ := time.Parse(layout, infoVuelo[6])

	vueloNuevo := &vuelo{
		flightNumber:    infoVuelo[0],
		airline:         infoVuelo[1],
		conexion:        conexionNueva,
		tailNumber:      infoVuelo[4],
		priority:        infoVuelo[5],
		date:            fechaNueva,
		deperatureDelay: infoVuelo[7],
		airTime:         infoVuelo[8],
		cancelled:       infoVuelo[9],
	}
	return vueloNuevo
}

func (vuelo *vuelo) GetFlightNumber() string {
	return vuelo.flightNumber
}

func (vuelo *vuelo) GetAirLine() string {
	return vuelo.airline
}

func (vuelo *vuelo) GetOrigin() string {
	return vuelo.conexion.originAirport
}

func (vuelo *vuelo) GetDestino() string {
	return vuelo.conexion.destinationAirport
}

func (vuelo *vuelo) GetConexion() Conexion {
	return vuelo.conexion
}

func (vuelo *vuelo) GetTailNumber() string {
	return vuelo.tailNumber
}

func (vuelo *vuelo) GetPriority() int {
	priority, err := strconv.Atoi(vuelo.priority)
	if err != nil {
		panic("Invalid priority")
	}
	return priority
}

func (vuelo *vuelo) GetDate() time.Time {
	return vuelo.date
}

func (vuelo *vuelo) GetDepartureDelay() string {
	return vuelo.deperatureDelay
}

func (vuelo *vuelo) GetAirTime() string {
	return vuelo.airTime
}

func (vuelo *vuelo) GetCancelled() string {
	return vuelo.cancelled
}

func CrearConexion(origen string, destino string) *Conexion {
	conexionNueva := &Conexion{
		originAirport:      origen,
		destinationAirport: destino,
	}

	return conexionNueva
}
