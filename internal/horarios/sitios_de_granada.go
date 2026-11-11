package horarios

type NombreSitio string

type SitiosDeGranada struct {
	horarios map[NombreSitio]Horario
}
