package main

import "fmt"

const (
	IMCMaigreur  = 18.5
	IMCNormal    = 25.0
	IMCSurpoids  = 30.0
)

const Nom = "Sahrane"

func main() {
	poids := 70.5
	taille := 1.75

	imc := poids / (taille * taille)

	fmt.Printf("%s - IMC : %.2f\n", Nom, imc)

	if imc < IMCMaigreur {
		fmt.Println("Catégorie : Maigreur")
	} else if imc < IMCNormal {
		fmt.Println("Catégorie : Normal")
	} else if imc < IMCSurpoids {
		fmt.Println("Catégorie : Surpoids")
	} else {
		fmt.Println("Catégorie : Obésité")
	}
}
