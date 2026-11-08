package horarios

type FranjaDeApertura struct {
	Apertura      HoraDelDia
	Cierre        HoraDelDia
	UltimaEntrada *HoraDelDia
}
