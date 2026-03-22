package adreces

import (
	"strings"
	"testing"
)

func TestPrefixPerParaulaSegonsRegles(t *testing.T) {
	if p := prefixPerParaula(ParaulaVia{Text: "Raval", Genere: Masculi}); p != "del" {
		t.Fatalf("prefix Raval = %s, volia del", p)
	}

	if p := prefixPerParaula(ParaulaVia{Text: "Font", Genere: Femeni}); p != "de la" {
		t.Fatalf("prefix Font = %s, volia de la", p)
	}

	if p := prefixPerParaula(ParaulaVia{Text: "Horta", Genere: Femeni}); p != "de l'" {
		t.Fatalf("prefix Horta = %s, volia de l'", p)
	}

	if p := prefixPerParaula(ParaulaVia{Text: "Esglesia", Genere: Femeni}); p != "de l'" {
		t.Fatalf("prefix Esglesia = %s, volia de l'", p)
	}
}

func TestComposaNomViaSenseEspaiEnApostrof(t *testing.T) {
	nom := composaNomVia(ParaulaVia{Text: "Olivera", Genere: Femeni})
	if nom != "de l'Olivera" {
		t.Fatalf("nom = %s, volia de l'Olivera", nom)
	}

	nom2 := composaNomVia(ParaulaVia{Text: "Castell", Genere: Masculi})
	if nom2 != "del Castell" {
		t.Fatalf("nom = %s, volia del Castell", nom2)
	}
}

func TestComencaAmbVocalOH(t *testing.T) {
	if !comencaAmbVocalOH("Hotel") {
		t.Fatalf("Hotel hauria de comptar com vocal/H")
	}

	if !comencaAmbVocalOH("Olivera") {
		t.Fatalf("Olivera hauria de comptar com vocal/H")
	}

	if comencaAmbVocalOH("Raval") {
		t.Fatalf("Raval no hauria de comptar com vocal/H")
	}
}

func TestGeneraAdrecaCasa(t *testing.T) {
	a := GeneraAdreca(Casa)

	if a.Tipus != Casa {
		t.Fatalf("tipus = %s, volia Casa", a.Tipus)
	}

	if a.NumeroPis != 0 {
		t.Fatalf("NumeroPis = %d, volia 0", a.NumeroPis)
	}

	if a.Porta != "" {
		t.Fatalf("Porta = %s, volia buit", a.Porta)
	}
}

func TestGeneraAdrecaPis(t *testing.T) {
	a := GeneraAdreca(Pis)

	if a.Tipus != Pis {
		t.Fatalf("tipus = %s, volia Pis", a.Tipus)
	}

	if a.NumeroPis < 1 || a.NumeroPis > 10 {
		t.Fatalf("NumeroPis = %d, fora de rang", a.NumeroPis)
	}

	if a.Porta == "" {
		t.Fatalf("Porta buida, volia una porta")
	}
}

func TestStringFormatCasaIPis(t *testing.T) {
	casa := Adreca{Tipus: Casa, TipusVia: "Carrer", NomVia: "de la Font", Numero: 12}
	if casa.String() != "Carrer de la Font, 12" {
		t.Fatalf("String casa = %s", casa.String())
	}

	pis := Adreca{Tipus: Pis, TipusVia: "Avinguda", NomVia: "del Mar", Numero: 7, NumeroPis: 3, Porta: "B"}
	if pis.String() != "Avinguda del Mar, 7, 3 B" {
		t.Fatalf("String pis = %s", pis.String())
	}

	if !strings.Contains(pis.String(), "3 B") {
		t.Fatalf("el string del pis no inclou pis i porta: %s", pis.String())
	}
}
