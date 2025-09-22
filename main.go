package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID    int
	Nom   string
	Email string
}

func main() {
	contacts := make(map[int]Contact)
	nextID := 1

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Mini CRM ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Quitter")

		fmt.Print("Votre choix : ")
		input, _ := reader.ReadString('\n')
		choix, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Erreur : veuillez entrer un nombre.")
			continue
		}

		switch choix {
		case 1:
			fmt.Print("Nom du contact : ")
			nom, _ := reader.ReadString('\n')
			nom = strings.TrimSpace(nom)

			fmt.Print("Email du contact : ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			contacts[nextID] = Contact{ID: nextID, Nom: nom, Email: email}
			fmt.Println("Contact ajouté avec ID", nextID)
			nextID++

		case 2:
			fmt.Println("Liste des contacts :")
			for _, c := range contacts {
				fmt.Println(c.ID, c.Nom, c.Email)
			}

		case 3:
			fmt.Println("Supprimer un contact")

		case 4:
			fmt.Println("Quitter")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
