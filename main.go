package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/utescu/generadorscatalans/persones"
	"github.com/utescu/generadorscatalans/pobles"
)

func main() {
	// Genera 10 pobles
	for i := range 10 {
		fmt.Println(i, pobles.GeneraNomDePoble())
	}

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
		fmt.Println(j, individu.Nom, individu.Cognom, individu.Cognom2, individu.Sexe)
	}

	fmt.Println("--------------")
	fmt.Println(persones.GeneraNom(persones.Home))
}
