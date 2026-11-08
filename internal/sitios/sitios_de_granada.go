package sitios

import "github.com/santiaguiyo/CulturaGRX/internal/horarios"

type NombreSitio string

type SitiosDeGranada struct {
	horarios map[NombreSitio]horarios.Horario
}
