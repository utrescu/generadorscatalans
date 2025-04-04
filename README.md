# Generador de noms i pobles catalans

Mòduls en llenguatge GO per generar de forma aleatòria tant noms de pobles com noms de persones

També pot generar un ofici aleatòri per la persona que s'ha creat

Els faig servir per preparar pràctiques de Bases de dades o APIS. Els generadors em permeten omplir automàticament grans quantitats de dades que poden semblar reals.

## persones

Pot generar noms aleatoris de persones amb cognoms o no.

### GeneraNom(s Sexe)

Retorna un nom aleatòri d'home o de dona en funció del Sexe especificat

```golang
fmt.Println(persones.GeneraNom(persones.Home))
```

### GeneraCognom()

Retorna un cognom aleatòri 

```
fmt.Println(persones.GeneraCognom())
```

### GeneraPersona(s Sexe)

Retorna el nom i cognoms d'una persona del sexe especificat en una estructura 'Persona'

```golang
for j := range 10 {
    var sex persones.Sexe
    que := rand.IntN(2)
    if que == 0 {
        sex = persones.Home
    } else {
        sex = persones.Dona
    }
    individu := persones.GeneraPersona(sex)
    fmt.Println(j, individu.Nom, individu.Cognom, individu.Cognom2, individu.Sexe)
}
```

Genera una llista de persones:

```log
0 Laia Ibarra Llabot Dona
1 Maria Lluisa Medrano Petit Dona
2 Damià Piqueras Llorens Home
3 Charly Llorens Casajuana Home
4 Diego Oveja Guixà Home
5 Max Julià Pi Home
6 Marta Moragues Aubert Dona
7 Montse Mach Llofriu Dona
8 Pepa Jordana Mata Dona
9 Magda Hom Matute Dona
```

## pobles

El mòdul genera un nom de poble inventat de forma aleatòria. 

### GeneraNomDePoble()

Crea un nom de poble a partir d'un patró semblant al que s'ha fet servir per donar nom als pobles catalans

Per exemple aquest codi:

```
for i := range 10 {
	fmt.Println(i, pobles.GeneraNomDePoble())
}
```

Imprimirà per pantalla una llista com aquesta: (cada cop serà diferent)

```log
0 Santa Marta de la Font
1 Bellllargaruda de l'Església
2 Santa Pilar del Punt
3 Bellescatosa
4 Sant Jou del Matoll
5 Santa Rosalia
6 Vilallisa de la Soca
7 Sant Francesc
8 Sant Josep de Mar
9 Santa Isabel de les Manies
```

Que surti un nom de poble real és pura casualitat.

### Oficis

Generador d'oficis de forma aleatòria

```golang
for i:= range 10 {
    fmt.Println(i, oficis.GeneraOfici())
}
```

Dóna una cosa com aquesta:

```log
0 Paleta
1 Físic
2 Locutor
3 Transportista
4 Picapedrer
5 Carboner
6 Actor
7 Economista
8 Intel·lectual
9 Flequer
```

## TODO

- Permetre entrar les dades que fan servir els paquets des de fitxers.
- Afegir nous generadors
