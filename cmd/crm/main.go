package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gocrm/internal/models"
	"gocrm/internal/storage"
)

var file = "internal/contacts.json"

func main() {
	contacts, _ := storage.ReadFile(file)

	nextID := 1
	for _, c := range contacts {
		if c.ID >= nextID {
			nextID = c.ID + 1
		}
	}

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

			newContact := models.Contact{ID: nextID, Nom: nom, Email: email}
			contacts = append(contacts, newContact)
			storage.WriteFile(file, contacts)

			fmt.Println("Contact ajouté avec ID", nextID)
			nextID++

		case 2:
			fmt.Println("Liste des contacts :")
			for _, c := range contacts {
				fmt.Println(c.ID, c.Nom, c.Email)
				storage.ReadFile(file)
				fmt.Println("Les contacts", contacts)
			}

		case 3:
			fmt.Print("ID du contact à supprimer : ")
			idInput, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idInput))
			if err != nil {
				fmt.Println("Erreur : ID invalide.")
				continue
			}

			found := false
			for i, c := range contacts {
				if c.ID == id {
					contacts = append(contacts[:i], contacts[i+1:]...)
					storage.WriteFile(file, contacts)
					fmt.Println("Contact avec ID", id, "supprimé.")
					found = true
					break
				}
			}
			if !found {
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

			updated := false
			for i, c := range contacts {
				if c.ID == id {
					fmt.Printf("Nom actuel : %s | Nouveau nom : ", c.Nom)
					newNom, _ := reader.ReadString('\n')
					newNom = strings.TrimSpace(newNom)

					fmt.Printf("Email actuel : %s | Nouvel email : ", c.Email)
					newEmail, _ := reader.ReadString('\n')
					newEmail = strings.TrimSpace(newEmail)

					if newNom != "" {
						c.Nom = newNom
					}
					if newEmail != "" {
						c.Email = newEmail
					}
					contacts[i] = c
					storage.WriteFile(file, contacts)
					fmt.Println("Contact mis à jour.")
					updated = true
					break
				}
				if !updated {
					fmt.Println("Aucun contact trouvé avec cet ID.")
				}
			}

		case 5:
			fmt.Println("Quitter")
			return

		default:
			fmt.Println("Choix invalide.")
		}
	}
}
