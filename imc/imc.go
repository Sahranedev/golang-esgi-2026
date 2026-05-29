package main

import "fmt"
/* On déclare des constantes en dehors du scope de la function pour les valeurs de l'IMC */
const (
	IMCMaigreur  = 18.5
	IMCNormal    = 25.0
	IMCSurpoids  = 30.0
)

const Nom = "Sahrane"

func main() {
	// on peut utiliser := pour déclarer une variable et assigner une valeur à la fois car on est a L'INTEIEUR d'une fonction
	poids := 70.5
	taille := 1.75

	imc := poids / (taille * taille)

	fmt.Printf("%s - IMC : %.2f\n", Nom, imc)
		// ici j'ai fait le choix d'un if / else if classique car on a 3 cas fixe possibile, j'ai voulu rester simple et clair
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
