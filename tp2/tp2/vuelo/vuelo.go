package vuelo

import "time"

type Vuelo interface {
	GetFlightNumber() string
	GetAirLine() string
	GetOrigin() string
	GetDestino() string
	GetConexion() Conexion
	GetTailNumber() string
	GetPriority() int
	GetDate() time.Time
	GetDepartureDelay() string
	GetAirTime() string
	GetCancelled() string
}
