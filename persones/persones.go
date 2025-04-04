package persones

import "math/rand/v2"

var homes = []string{
	"Pere", "Manel", "Marti", "Joan", "Pep", "Toni", "Sam", "Lluis", "Cesc", "Joan", "Josep", "Frederic", "Filomenu",
	"Enric", "Jou", "Bernat", "Francesc", "Paco", "Zeus", "Miquel", "Joan Josep",
	"Met", "Homer", "Ot", "Iu", "Max", "Pol", "Pitu", "Jean", "Òscar", "Jesus", "Alfons",
	"Sergi", "David", "Daniel", "Josep Antoni", "Joanot", "Cristian", "Diego", "Didac", "Rufo", "Quim", "Joaquim", "Xavier",
	"Baldiri", "Víctor", "Ernest", "Carles", "Charly", "Gerard", "Alfred", "Damià", "Robert", "Ernest",
	"Yander", "Rubén",
}

var dones = []string{
	"Anna", "Maria", "Mercè", "Angela", "Maria Lluisa", "Bernarda", "Helena", "Antonia", "Aina", "Queralt",
	"Felicitat", "Fe", "Mònica", "Mireia", "Marta", "Magda", "Susanna", "Sònia", "Pepa", "Candela", "Francesca",
	"Montserrat", "Agnès", "Pilar", "Magda", "Teresa", "Filomena", "Soraya", "Magdalena", "Hilda",
	"Martina", "Berta", "Ramona", "Angela", "Rosalia", "Montse", "Laia", "Frederica", "Pati", "Patricia",
	"Victòria", "Maria Isabel", "Isabel",
}

var cognoms = []string{
	"Pi", "Marti", "Giménez", "Perez", "González", "López", "Vilamoll", "Puig",
	"Serra", "Ruiz", "Roure", "Pujol", "Boix", "Bru", "Brau", "Soler", "Sonet", "Marcu",
	"Terrats", "Pou", "Fernàndez", "Rodríguez", "Gasull", "Aubert", "Cardona", "Casajuana",
	"Codina", "Gasull", "Gual", "Llach", "Maymó", "Bonastre", "Bonaterra", "Armengol", "Alemany",
	"Guasch", "Pitarch", "Raga", "Ricart", "Sabater", "Solé", "Taberner", "Tió", "Vives", "Agulla",
	"Alba", "Albí", "Albiol", "Alcoberro", "Guardiola", "Tubau", "Guvern", "Guevara", "Guilella", "Guimerà",
	"Guirao", "Guixà", "Gutierrez", "Hernández", "Hom", "Hosta", "Ubach", "Huguet", "Ibarra", "Icart",
	"Iglesias", "Ingla", "Iranzo", "Isanta", "Isern", "Jaques", "Jara", "Jiménez", "Jodar", "Jorba", "Jordana",
	"Jorquera", "Jornet", "Jové", "Juanola", "Julià", "Juncosa", "Junyent", "Just", "Juvany", "Laborda", "Lacalle",
	"Lara", "Leal", "Latre", "León", "Linares", "Lladó", "Llabot", "Llansana", "Llanet", "Llatser", "Llauradó",
	"Llenas", "Llesera", "Lloberes", "Llobet", "Llofriu", "Llorens", "Lloris", "Llort", "Lluch", "Llucià", "Llorente",
	"Luque", "Mach", "Macià", "Macrané", "Magrinyà", "Majós", "Maldonado", "Malé", "Malla", "Mallafré", "Mata", "Maluquer",
	"Mancho", "Manent", "Mans", "Manyé", "Manzano", "Lozano", "Marata", "Marcet", "March", "Marcó", "Marcos", "Margalef",
	"Maginet", "Marín", "Marlés", "Marot", "Marqués", "Marsà", "Marsany", "Martell", "Maruny", "Masachs", "Masana",
	"Mascarell", "Masclans", "Masdefiol", "Masdevall", "Masllorens", "Masmitjà", "Masolivé", "Masvidal", "Matamala",
	"Matarodona", "Matute", "Maurell", "Mayans", "Mayné", "Medina", "Medrano", "Mercadal", "Mercadé", "Messegué",
	"Mínguez", "Miravet", "Mitjavila", "Molas", "Monné", "Mont", "Monells", "Montià",
	"Montroig", "Montió", "Montràs", "Morillo", "Moragues", "Moron", "Morral", "Moya", "Mula", "Mundó", "Munné",
	"Murga", "Navajas", "Negre", "Nogué", "Not", "Oliva", "Oveja", "Olivella", "Oliveras", "Ordeig", "Orobitg", "Ortega",
	"Ortiz", "Padilla", "Padrosa", "Pairó", "Palacios", "Palau", "Paladell", "Palau", "Pallarol", "Pallejà",
	"Pellicer", "Palomera", "Palop", "Paraire", "Páramo", "Paradís", "Parer", "Paret", "Parra",
	"Pascal", "Pastoret", "Paula", "Pedreño", "Pedrol", "Pelayo", "Pelegrí", "Peret", "Pericas", "Pesarrodona", "Petit", "Petrus",
	"Rey", "Pey", "Pinyas", "Piñol", "Brunyol", "Picanyol", "Pia", "Rubio", "Piferrer", "Pijoan", "Pinot", "Piqueras",
	"Platja", "Pont", "Vinals", "Visent", "Viu", "Reig",
}


type Persona struct {
	Nom     string
	Cognom  string
	Cognom2 string
	Sexe    string
}

func (p Persona) GetNomSencer() string {
	return p.Nom + " " + p.Cognom + " " + p.Cognom2
}

func (p Persona) IsHome() bool {
	return p.Sexe == "H"
}

func (p Persona) IsDona() bool {
	return p.Sexe == "D"
}

func GeneraNom(s Sexe) string {
	if s == Home {
		return homes[rand.IntN(len(homes))]
	}
	return dones[rand.IntN(len(dones))]
}

func GeneraCognom() string {
	return cognoms[rand.IntN(len(cognoms))]
}

func GeneraPersona(s Sexe ) Persona {
	return Persona{
		Nom:     GeneraNom(s),
		Cognom:  GeneraCognom(),
		Cognom2: GeneraCognom(),
		Sexe:    s.String(),
	}
}

type Sexe int

const (
	Home Sexe = iota
	Dona  
)


func (ss Sexe) String() string {
	switch ss {
	case Home: 
		return "Home"
	case Dona:
		return "Dona"
	}
	return "Desconegut"
}


