package main

import "fmt"

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

// on déclare une fonction qui retourne le nom complet de la personne
func (p Personne) NomComplet() string {
	return fmt.Sprintf("%s %s", p.Prenom, p.Nom)
}

// on déclare une fonction qui retourne la présentation de la personne
func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans, %s", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}
// on déclare une fonction qui retourne la formatage de l'adresse de la personne
func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

// on déclare une structure qui embbedded les structures Personne et Adresse (du coup on peut appeler ça une structure imbriquée ou composées, ona  vu en cours que Go fonctionne par composition)
type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}
// on déclare une fonction qui retourne la fiche de l'employé
func (e Employe) FicheEmploye() string {
	return fmt.Sprintf("%s, %s, %.2f, %s", e.Presentation(), e.Poste, e.Salaire, e.Adresse.Format())
}

// on déclare une fonction qui augmente le salaire de l'employé, on utilise le recepteur pointeur pour pouvoir modifier la variable réeelle et pas sur une copie de l'instance 
func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire += e.Salaire * pct / 100
}
// structure Etudiant qui embbedded la structure Personne
type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "TB"
	case e.Moyenne >= 14:
		return "B"
	case e.Moyenne >= 12:
		return "AB"
	default:
		return "P"
	}
}

func main() {
	employe1 := Employe{
		Personne: Personne{Prenom: "Totoprenom", Nom: "Totonom", Age: 99, Email: "toto@science-u.fr"},
		Adresse:  Adresse{Rue: "SAns soucis", Ville: "Paris", CodePostal: "69003"},
		Poste:    "Dev",
		Salaire:  41000,
	}

	employe2 := Employe{
		Personne: Personne{Prenom: "Tataprenom", Nom: "Tatanom", Age:37, Email: "tata@science-u.fr"},
		Adresse:  Adresse{Rue: "Avec des soucis", Ville: "Paris", CodePostal: "69003"},
		Poste:    "DevOps",
		Salaire:  55000,
	}

	etudiant1 := Etudiant{
		Personne: Personne{Prenom: "Sahrane", Nom: "Guassemi", Age: 29, Email: "sahranego@myges.fr.com"},
		Promo:    "IW M2",
		Moyenne:  18,
	}

	etudiant2 := Etudiant{
		Personne: Personne{Prenom: "Plus", Nom: "Dinspiration", Age: 24, Email: "plus-dinspiration@myges.fr"},
		Promo:    "IW M2",
		Moyenne:  12,
	}

	fmt.Println(employe1.FicheEmploye())
	fmt.Println(employe2.FicheEmploye())
	fmt.Println(etudiant1.Presentation(), etudiant1.Promo, etudiant1.Moyenne, etudiant1.MentionObtenue())
	fmt.Println(etudiant2.Presentation(), etudiant2.Promo, etudiant2.Moyenne, etudiant2.MentionObtenue())
}
