package main

import (
	"fmt"
	"time"
)

func main() {
	// je met ma date de naissance
	birthDate := time.Date(1997, time.April, 16, 0, 0, 0, 0, time.Local)
	// je récupère la date actuelle
	now := time.Now()
	// je récupère l'année de la date actuelle et je vais soustraite l'année  de ma date e naissance
	age := now.Year() - birthDate.Year()	
	fmt.Println("J'ai", age, "ans")
}
