package operaciones

type operacion struct {
	simbolo       Simbolo
	precedencia   Precedencia
	asociatividad Asociativa
}

// CrearOperacion Crea una operación para el simbolo
func CrearOperacion(simbolo Simbolo) Operacion {
	switch simbolo {
	case Suma:
		return operacion{Suma, PrecSumaResta, Izquierda}
	case Resta:
		return operacion{Resta, PrecSumaResta, Izquierda}
	case Multiplicacion:
		return operacion{Multiplicacion, PrecMultiplicacion, Izquierda}
	case Division:
		return operacion{Division, PrecDivision, Izquierda}
	case Potencia:
		return operacion{Potencia, PrecPotencia, Derecha}
	}

	panic("símbolo de operación inválido: " + simbolo)

}

func (op operacion) Simbolo() Simbolo {
	return op.simbolo
}

func (op operacion) Precedencia() Precedencia {
	return op.precedencia
}

func (op operacion) Asociatividad() Asociativa {
	return op.asociatividad
}
