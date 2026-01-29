package operaciones

type Operacion interface {
	//Simbolo devuelve el simbolo de la operacion
	Simbolo() Simbolo

	//Precedencia devuelve el Precedencia de la operacion
	Precedencia() Precedencia

	//Asociatividad devuelve la Asociativa de la operacion
	Asociatividad() Asociativa
}

type Asociativa string

const (
	Izquierda Asociativa = "I"
	Derecha   Asociativa = "D"
)

type Precedencia int

const (
	PrecSumaResta      Precedencia = 2
	PrecMultiplicacion Precedencia = 3
	PrecDivision       Precedencia = 3
	PrecPotencia       Precedencia = 4
)

type Simbolo string

const (
	Suma           Simbolo = "+"
	Resta          Simbolo = "-"
	Multiplicacion Simbolo = "*"
	Division       Simbolo = "/"
	Potencia       Simbolo = "^"
)
