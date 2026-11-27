package horarios

type HorarioDelSitio struct {
	temporadas   []Temporada
	diasCerrados map[Fecha]struct{}
}
