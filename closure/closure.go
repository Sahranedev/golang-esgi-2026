package main

import (
	"errors"
	"fmt"
)

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division par zéro")
		}
		return a / b, nil
	default:
		return 0, errors.New("opérateur inconnu")
	}
}

func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 { return a + b }
	case "-":
		return func(a, b float64) float64 { return a - b }
	case "*":
		return func(a, b float64) float64 { return a * b }
	case "/":
		return func(a, b float64) float64 { return a / b }
	default:
		return func(a, b float64) float64 { return 0 }
	}
}

func main() {
	for {
		var a, b float64
		var op string

		fmt.Print("Dans la zone vide du terminal il faut rentrer deux nombres et un opétateur pour taper quit pour quitter")
		fmt.Scan(&a, &b, &op)
		// on break la boucle si on tape quit (demandé dans la consigne)
		if op == "quit" {
			break
		}

		result, err := operer(a, b, op)
		if err != nil {
			fmt.Println("erreur")
		} else {
			fmt.Println(result)
		}
	}
}
