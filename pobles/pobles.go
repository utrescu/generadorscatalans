package pobles

import (
	"math/rand/v2"
	"strings"

	"github.com/utrescu/generadorscatalans/persones"
)

const PROVABILITATDEPOSTNOM = 75

var prepend = []string{"Sant ", "Sant ", "Sant ", "Sant ", "Sant ", "Santa ", "Santa ",
	"Santa ", "Santa ", "Santa ", "Vila", "Vila", "Vila", "Vila", "Vila",
	"Bell", "Castell", "Vall", "Vall",
}

var adjectius = []string{"allargassada", "fort", "forta", "fràgil", "gros", "grossa",
	"llarg", "llarga", "llargarut", "llargaruda", "massís", "massissa", "robust", "robusta", "rodanxó",
	"rodanxona", "plana", "pelat", "pelada", "pelut", "peluda", "rodó", "rodona", "aspre", "bru", "bruna",
	"cargolada", "curt", "curta", "dur", "dura", "fi", "fina", "flonja", "llanosa", "llis", "llisa", "suau",
	"vellut", "arrugada", "colrada", "dura", "dur", "escatosa", "greixosa", "greixós", "llisa", "llis",
	"rasposa", "resseca", "rugosa", "setinada", "suau", "dispersa", "dispers", "esclarissada", "espessa",
	"espes", "estarrufada", "estarrufat", "guarnida", "lluent", "fort", "forta", "cavall", "quieta",
}

var postname = []string{"de Mar", "de Baix", "de Dalt", "del Camí", "de la Plana",
	"de la Fresca", "de la Font", "de la Selva", "del Clot", "del Sapastre", "de l'Església", "de la Roca",
	"de la Fosca", "del Clar", "de la Soca", "del Matoll", "de l'Esbarzer", "de la Conca", "del Rosco",
	"de la Costa", "del Punt", "de les Manies",
}

func GeneraNomDePoble() string {

	nom := ""

	tePrepend := rand.IntN(len(prepend))
	preNom := prepend[tePrepend]
	if strings.HasPrefix(preNom, "Sant") {
		if strings.HasPrefix(preNom, "Santa") {
			nom = persones.GeneraNom(persones.Dona)
		} else {
			nom = persones.GeneraNom(persones.Home)
		}

	} else {
		nom = adjectius[rand.IntN(len(adjectius))]
	}

	if rand.IntN(100) < PROVABILITATDEPOSTNOM {
		return preNom + nom + " " + postname[rand.IntN(len(postname))]
	}

	return preNom + nom

}
