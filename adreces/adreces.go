package adreces

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type TipusAdreca int

const (
	Casa TipusAdreca = iota
	Pis
	Terreny
)

func (t TipusAdreca) String() string {
	switch t {
	case Casa:
		return "Casa"
	case Pis:
		return "Pis"
	case Terreny:
		return "Terreny"
	}
	return "Desconegut"
}

type Adreca struct {
	Tipus     TipusAdreca
	TipusVia  string
	NomVia    string
	Numero    int
	NumeroPis int
	Porta     string
}

func (a Adreca) String() string {
	base := fmt.Sprintf("%s %s, %d", a.TipusVia, a.NomVia, a.Numero)
	if a.Tipus == Casa || a.Tipus == Terreny {
		return base
	}

	return fmt.Sprintf("%s, %d %s", base, a.NumeroPis, a.Porta)
}

type Genere int

const (
	Masculi Genere = iota
	Femeni
	FemeniPlural
	MasculiPlural
	Toponim
)

type ParaulaVia struct {
	Text   string
	Genere Genere
}

var tipusVia = []string{
	"Carrer",
	"Avinguda",
	"Placa",
	"Carretera",
	"Ronda",
	"Passeig",
	"Travessera",
	"Camí",
}

var paraulesVia = []ParaulaVia{
	{Text: "Mar", Genere: Masculi},
	{Text: "Font", Genere: Femeni},
	{Text: "Pineda", Genere: Femeni},
	{Text: "Roca", Genere: Femeni},
	{Text: "Castell", Genere: Masculi},
	{Text: "Pedra", Genere: Femeni},
	{Text: "Romani", Genere: Masculi},
	{Text: "Olivera", Genere: Femeni},
	{Text: "Muntanya", Genere: Femeni},
	{Text: "Mercat", Genere: Masculi},
	{Text: "Esglesia", Genere: Femeni},
	{Text: "Conca", Genere: Femeni},
	{Text: "Pi", Genere: Masculi},
	{Text: "Ametller", Genere: Masculi},
	{Text: "Rosella", Genere: Femeni},
	{Text: "Falguera", Genere: Femeni},
	{Text: "Mirador", Genere: Masculi},
	{Text: "Sardana", Genere: Femeni},
	{Text: "Raval", Genere: Masculi},
	{Text: "Portal", Genere: Masculi},
	{Text: "Horta", Genere: Femeni},
	{Text: "Bosc", Genere: Masculi},
	{Text: "Gavina", Genere: Femeni},
	{Text: "Clot", Genere: Masculi},
	{Text: "Pedra blanca", Genere: Femeni},
	{Text: "Escudella", Genere: Femeni},
	{Text: "Escudillers", Genere: Toponim},
	{Text: "Falgueres", Genere: FemeniPlural},
	{Text: "Pinedes", Genere: FemeniPlural},
	{Text: "Pardals", Genere: MasculiPlural},
	{Text: "Ajuntament", Genere: Masculi},
	{Text: "Torre", Genere: Femeni},
	{Text: "Centre", Genere: Masculi},
	{Text: "3 cantons", Genere: MasculiPlural},
	{Text: "Presó", Genere: Femeni},
	{Text: "Major", Genere: Toponim},
	{Text: "Església", Genere: Femeni},
	{Text: "Pau", Genere: Femeni},
	{Text: "Llibertat", Genere: Femeni},
	{Text: "Pau Casals", Genere: Toponim},
	{Text: "Pau Picasso", Genere: Toponim},
	{Text: "Jacint Verdaguer", Genere: Toponim},
	{Text: "Ramon Llull", Genere: Toponim},
	{Text: "Francesc Macià", Genere: Toponim},
	{Text: "Merce Rodoreda", Genere: Toponim},
	{Text: "Badaboc", Genere: Masculi},
	{Text: "Àvia", Genere: Femeni},
	{Text: "Àngel", Genere: Masculi},
	{Text: "Rossinyol", Genere: Masculi},
	{Text: "Lleó", Genere: Masculi},
	{Text: "Lleona", Genere: Femeni},
	{Text: "Roser", Genere: Masculi},
	{Text: "Canigó", Genere: Masculi},
	{Text: "Valentí Almirall", Genere: Toponim},
	{Text: "Gran", Genere: Toponim},
	{Text: "General Moragues", Genere: Toponim},
	{Text: "Almogavers", Genere: MasculiPlural},
	{Text: "Molí de can Ferrer", Genere: Masculi},
	{Text: "Arquitecte borratxo", Genere: Masculi},

}

var portes = []string{
	"A",
	"B",
	"C",
	"D",
	"1a",
	"2a",
	"3a",
}

func GeneraNomVia() string {
	paraula1 := paraulesVia[rand.IntN(len(paraulesVia))]
	nom := composaNomVia(paraula1)
	
	if paraula1.Genere != Toponim && rand.IntN(100) < 15 {
		paraula2 := paraulesVia[rand.IntN(len(paraulesVia))]
		if paraula2.Text != paraula1.Text && paraula2.Genere != Toponim {
			return fmt.Sprintf("%s i %s", nom, composaNomVia(paraula2))
		}
	}

	return nom
}

func composaNomVia(paraula ParaulaVia) string {
	prefix := prefixPerParaula(paraula)
	if strings.HasSuffix(prefix, "'") {
		return prefix + paraula.Text
	}

	return prefix + " " + paraula.Text
}

func prefixPerParaula(paraula ParaulaVia) string {
	if (paraula.Genere == Toponim) {
		return ""
	}
	if (paraula.Genere == FemeniPlural) {
		return "de les"
	}
	if (paraula.Genere == MasculiPlural) {
		return "dels"
	}

	if comencaAmbVocalOH(paraula.Text) {
		return "de l'"
	}

	if paraula.Genere == Femeni {
		return "de la"
	}

	return "del"
}

func comencaAmbVocalOH(text string) bool {
	runes := []rune(strings.TrimSpace(strings.ToLower(text)))
	if len(runes) == 0 {
		return false
	}

	switch runes[0] {
	case 'a', 'e', 'i', 'o', 'u', 'h', 'à', 'á', 'è', 'é', 'í', 'ï', 'ò', 'ó', 'ú', 'ü':
		return true
	default:
		return false
	}
}

func GeneraNumero() int {
	return rand.IntN(250) + 1
}

func GeneraAdreca(t TipusAdreca) Adreca {

	via := tipusVia[rand.IntN(len(tipusVia))]
	a := Adreca{
		Tipus:    t,
		TipusVia: via,
		NomVia:   GeneraNomVia(),
		Numero:   GeneraNumero(),
	}

	if t == Pis {
		a.NumeroPis = rand.IntN(10) + 1
		a.Porta = portes[rand.IntN(len(portes))]
	}

	return a
}

func GeneraAdrecaAleatoria() Adreca {
	if rand.IntN(3) == 0 {
		return GeneraAdreca(Casa)
	}

	if rand.IntN(3) == 1 {
		return GeneraAdreca(Terreny)
	}

	return GeneraAdreca(Pis)
}
