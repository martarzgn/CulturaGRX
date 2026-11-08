package horarios

import "time"

type HorarioSemanal map[time.Weekday][]FranjaDeApertura
