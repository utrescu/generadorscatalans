package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/utrescu/generadorscatalans/oficis"
	"github.com/utrescu/generadorscatalans/persones"
	"github.com/utrescu/generadorscatalans/pobles"
	"github.com/utrescu/generadorscatalans/pirates"
)

func main() {
	// Genera 10 pobles
	fmt.Println(".--------------------")
	fmt.Println("Genera 10 pobles imaginaris")
	fmt.Println("--------------------")
	for i := range 10 {
		fmt.Println(i, pobles.GeneraNomDePoble())
	}

	fmt.Println("--------------------")
	fmt.Println("Genera 10 persones inventades")
	fmt.Println("--------------------")
	// Genera 10 persones
	for j := range 10 {
		var sex persones.Sexe
		que := rand.IntN(2)
		if que == 0 {
			sex = persones.Home
		} else {
			sex = persones.Dona
		}
		individu := persones.GeneraPersona(sex)
		fmt.Println(j, individu.Nom, individu.Cognom, individu.Cognom2, "és", individu.Sexe)
	}
	fmt.Println("--------------------")
	fmt.Println("Selecciona 10 oficis aleatòriament")
	fmt.Println("--------------------")
	for i := range 10 {
		fmt.Println(i, oficis.GeneraOfici())
	}

	fmt.Println("--------------------")
	fmt.Println("Genera 10 noms de pirata català")
	fmt.Println("--------------------")

	for i := range 10 {
		fmt.Println(i, pirates.GenerarNomDePirata())
	}
	fmt.Println("--------------------")
	
}
