package storage

import "fmt"

type Contact struct {
	Name  string
	Email string
	Phone string
}

type Storer interface {
	Add(contact *Contact) error
	GetAll() ([]Contact, error)
	GetById(id int) (*Contact, error)
	Update(id int, newEmail string) error
	Delete(id int) error
}

var ErrContactNotFound = func(id int) error {
	return fmt.Errorf("contact avec l'ID %d introuvable", id)
}
