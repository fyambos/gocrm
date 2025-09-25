package storage

import (
	"encoding/json"
	"gocrm/internal/models"
	"log"
	"os"
)

func ReadFile(file string) ([]models.Contact, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var contacts []models.Contact
	if err := json.Unmarshal(data, &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}

func WriteFile(file string, contacts []models.Contact) error {
	data, _ := json.MarshalIndent(contacts, "", "  ")
	return os.WriteFile(file, data, 0644)
}
