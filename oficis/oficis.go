package oficis

import "math/rand/v2"

var oficis = []string{
	"Ablegat",
	"Actor",
	"Administrador de sistemes",
	"Adornista",
	"Adroguer",
	"Adroguer",
	"Advocat",
	"Agent cívic",
	"Agent forestal",
	"Agent literari",
	"Agrimensor",
	"Alt càrrec",
	"Antiquari",
	"Antropòleg",
	"Aparcacotxes",
	"Apoderat",
	"Apotecari",
	"Aprenent",
	"Apuntador de teatre",
	"Àrbitre",
	"Argenter",
	"Armador",
	"Arquitecte",
	"Assessor financer",
	"Assessor personal",
	"Astronauta",
	"Ataconador",
	"Au-pair",
	"Autor",
	"Aviador de combat",
	"Ballarí",
	"Barber",
	"Bàrman",
	"Bibliotecari",
	"Biotecnoleg",
	"Bomber",
	"Brocanter",
	"Bugadera",
	"Calceter",
	"Calciner",
	"Calculadora humana",
	"Calderer",
	"Camarlenc",
	"Cambrer",
	"Candeler",
	"Candeler",
	"Cap de projecte",
	"Carboner",
	"Carnisser",
	"Càrrec de confiança",
	"Carter",
	"Cercoler",
	"Científic",
	"Cirurgia",
	"Companyia de Fusellers Guardaboscos Reials",
	"Conductor",
	"Contaire",
	"Coper",
	"Coraller",
	"Corder de viola",
	"Corrector de textos",
	"Corredor de finances",
	"Corretger",
	"Courer",
	"Criminòleg",
	"Crític",
	"Crupier",
	"Culleraire",
	"Dama de companyia",
	"Decorador",
	"Degà",
	"Delineant",
	"Detectiu",
	"Dida",
	"Docent",
	"Documentalista audiovisual",
	"Dramaturg",
	"Drover",
	"Economista",
	"Electricista",
	"Empresari",
	"Enginyer",
	"Enllustrador",
	"Entrenador",
	"Entrenament personal",
	"Escloper",
	"Escombraire",
	"Esmolet",
	"Espectacle de carrer",
	"Espectroscopista",
	"Esportista",
	"Estàtua vivent",
	"Fanaler",
	"Farmacèutic",
	"Fiscal",
	"Físic",
	"Fisioterapeuta",
	"Flequer",
	"Foto fixa",
	"Fotògraf",
	"Funcionari",
	"Fuster",
	"Futbolista",
	"Gallinaire",
	"Garimpeiro",
	"Genet",
	"Gestor de comunitats en línia",
	"Gestor d'informació",
	"Grum",
	"Guaita forestal",
	"Guardabosc",
	"Guarda rural",
	"Guia acompanyant",
	"Historiador",
	"Hortolà",
	"Infermer",
	"Intel·lectual",
	"Inventor",
	"Investigador privat",
	"Jutge",
	"Lingüista",
	"Llenyataire",
	"Locutor",
	"Mahout",
	"Maître",
	"Mangaka",
	"Manicur",
	"Maquillador",
	"Mariner",
	"Matemàtic",
	"Medievalista",
	"Mestre de cases",
	"Mestre d'escola",
	"Mestre d'obres",
	"Mestressa de casa",
	"Metge",
	"Minaire",
	"Missatger",
	"Modista",
	"Moliner",
	"Notari",
	"Odontòleg",
	"Orfebreria",
	"Pagès",
	"Pagès de remença",
	"Paisatgista",
	"Paleta",
	"Paredador",
	"Pastor",
	"Pelador de suro",
	"Periodista",
	"Perruquer",
	"Pescador",
	"Picapedrer",
	"Pintor de quadres",
	"Pintor",
	"Planxadora",
	"Polític",
	"Portabandera",
	"Portaveu",
	"Porter de cadena",
	"Porter",
	"Pràctic",
	"Presentador",
	"Presentador de notícies",
	"Procurador",
	"Productor de televisió",
	"Programador",
	"Psicòleg",
	"Recaptador d'impostos",
	"Recepcionista",
	"Rellotger",
	"Repartidor de pizzes",
	"Revisor",
	"Riberer",
	"Roadie",
	"Secretari",
	"Serralleria",
	"Soldat",
	"Sommelier",
	"Spin doctor",
	"Supervisor",
	"Tècnic",
	"Tècnic d'Operació de Línies Automàtiques",
	"Timoner",
	"Traginer",
	"Transportista",
	"Treballador domèstic",
	"Treballador independent",
	"Trencador",
	"Tutor",
	"Ullerer",
	"Venedor",
	"Vigilant de seguretat",
	"Visual image developer",
	"Xofer",
}

func GeneraOfici() string {
	return oficis[rand.IntN(len(oficis))]
}
