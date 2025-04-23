package pirates

import (
	"fmt"
	"math/rand"
	"github.com/utrescu/generadorscatalans/persones"
)

var adjectius = []string{"l'allargassat", "el fort", "el fràgil", "gros", "el ferotge", "l'astut", "el temerari",
	"espectre", "sanguinari", "veloç", "silenciós", "mut", "trencapixes", "matamosques", "l'assassí", "el lladre",
	"la paret", "sabater", "sense cor", "barba", "l'ós", "el tigre", "tres dits", "pixatorta", 
	"el llarg", "el llargarut", "massís", "el moliner", "el robust", "el rodanxó", "en plana", "el pelat", "el pelut", "el girarodó", 
	"l'aspre", "el bru", "el cargolat", "el curt", "el dur", "el pagès", "l'arrugat", "l'escatos", "el greixós", 
	"en resseca persones", "el dispersador", "l'estarrufat", "el cara de cavall", "en calabera", 
	"en trinxador", "el fantasma", "el furiós", "el bastard", "el bord", "catiu", "l'empordanès", "el francès",
	"el gavatx", "amb un ull", "el ros", "el roig", "pellroja", "barbudet", "anima pecadora", "miratort", "martell",
	"teixidor de mar", "Carboner", "el bastard de Cerdans", "el boter", "cabirolet", "matamoros", "morrofes",
	"l'esquerrà", "el pubill", "l'hereu", "llavifés", "dels naps", "lafont", "l'antic", "el rectoret",
	"el de la creu", "el de la ma negra", "ripart", "el tendre", "Trucafort", "Barbafina", "cua de llop",
	"Barrabam", "Moica", "el coix", "Canyadelles", "Pay", "Rocallaura", "Tallaferro", "Perpunter",
}


func GenerarNomDePirata() string {
	adjectiu := adjectius[rand.Intn(len(adjectius))]
	substantiu := persones.GeneraNom(persones.Home)

	cognom := persones.GeneraCognom()
	return fmt.Sprintf("%s '%s' %s", substantiu, adjectiu, cognom)
}
