package main

import "fmt"

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return fmt.Sprintf("%s %s", p.Prenom, p.Nom)
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans, %s", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf("%s, %s, %.2f, %s", e.Presentation(), e.Poste, e.Salaire, e.Adresse.Format())
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire += e.Salaire * pct / 100
}

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
		Personne: Personne{Prenom: "Alice", Nom: "Martin", Age: 32, Email: "alice@entreprise.fr"},
		Adresse:  Adresse{Rue: "12 rue de la Paix", Ville: "Paris", CodePostal: "75002"},
		Poste:    "Developpeuse",
		Salaire:  45000,
	}

	employe2 := Employe{
		Personne: Personne{Prenom: "Bob", Nom: "Dupont", Age: 45, Email: "bob@entreprise.fr"},
		Adresse:  Adresse{Rue: "8 avenue Victor Hugo", Ville: "Lyon", CodePostal: "69003"},
		Poste:    "Chef de projet",
		Salaire:  52000,
	}

	etudiant1 := Etudiant{
		Personne: Personne{Prenom: "Sahrane", Nom: "Guassemi", Age: 29, Email: "sahranego@gmail.com"},
		Promo:    "IW M2",
		Moyenne:  14,
	}

	etudiant2 := Etudiant{
		Personne: Personne{Prenom: "Clara", Nom: "Bernard", Age: 20, Email: "clara@esgi.fr"},
		Promo:    "IW M2",
		Moyenne:  12,
	}

	fmt.Println(employe1.FicheEmploye())
	fmt.Println(employe2.FicheEmploye())
	fmt.Println(etudiant1.Presentation(), etudiant1.Promo, etudiant1.Moyenne, etudiant1.MentionObtenue())
	fmt.Println(etudiant2.Presentation(), etudiant2.Promo, etudiant2.Moyenne, etudiant2.MentionObtenue())
}
