package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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
			fmt.Println("Ajouter un contact")
		case 2:
			fmt.Println("Lister les contacts")
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
