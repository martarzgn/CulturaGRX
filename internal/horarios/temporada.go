package horarios

import "time"

type DiaDelAnio struct {
	Mes time.Month
	Dia uint8
}

type Temporada struct {
	Desde          DiaDelAnio
	Hasta          DiaDelAnio
	HorarioSemanal HorarioSemanal
}

type Horario []Temporada
