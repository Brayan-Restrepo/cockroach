package models

type AddAnimalData struct {
	Name    string `json:"name"`
	Species string `json:"species"`
	Emoji   string `json:"emoji"`
}
