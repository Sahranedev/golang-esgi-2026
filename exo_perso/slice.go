package main

import "fmt"

func main() {
	// on crée un slice de 3 éléments avec une capacité de 5 DIRECTEMENT dans la mémoire du programme, ca permet ici qu'on est une lgoneur connues déja préparé dans la mémemoire 
	notes := make([]int, 3, 5)

	// on assigne des valeurs au slice
	notes[0] = 12
	notes[1] = 15
	notes[2] = 8

	fmt.Println("slice:", notes)
	fmt.Println("len:", len(notes))
	fmt.Println("cap:", cap(notes))

	// on ajoute un élément au slice avec append, vu qu'on avait prévu des le départ une capacité de 5 et qu'on avait aloué seulement 3 places, on peut ajouter un élément supplémentaire sans réajuster la capacité mémoire du slice 

	fmt.Println("après append:", notes)
	fmt.Println("len:", len(notes))
	fmt.Println("cap:", cap(notes))
}
