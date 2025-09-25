package models

type Contact struct {
	ID    int    `json:"id"`
	Nom   string `json:"nom"`
	Email string `json:"email"`
}
