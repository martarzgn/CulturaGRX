package horarios

import "time"

type DiaDelAño struct {
	Mes time.Month
	Dia uint8
}

type Temporada struct {
	Desde   DiaDelAño
	Hasta   DiaDelAño
	Semanal HorarioSemanal
}

type Horario []Temporada
