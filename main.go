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
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")

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
			fmt.Print("ID du contact à supprimer : ")
			idInput, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idInput))
			if err != nil {
				fmt.Println("Erreur : ID invalide.")
				continue
			}

			if _, ok := contacts[id]; ok {
				delete(contacts, id)
				fmt.Println("Contact avec ID", id, "supprimé.")
			} else {
				fmt.Println("Aucun contact trouvé avec cet ID.")
			}

		case 4:
			fmt.Print("ID du contact à mettre à jour : ")
			idInput, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idInput))
			if err != nil {
				fmt.Println("Erreur : ID invalide.")
				continue
			}

			if contact, ok := contacts[id]; ok {
				fmt.Printf("Nom actuel : %s | Nouveau nom : ", contact.Nom)
				newNom, _ := reader.ReadString('\n')
				newNom = strings.TrimSpace(newNom)

				fmt.Printf("Email actuel : %s | Nouvel email : ", contact.Email)
				newEmail, _ := reader.ReadString('\n')
				newEmail = strings.TrimSpace(newEmail)

				if newNom != "" {
					contact.Nom = newNom
				}
				if newEmail != "" {
					contact.Email = newEmail
				}

				contacts[id] = contact
				fmt.Println("Contact mis à jour.")
			} else {
				fmt.Println("Aucun contact trouvé avec cet ID.")
			}

		case 5:
			fmt.Println("Quitter")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
