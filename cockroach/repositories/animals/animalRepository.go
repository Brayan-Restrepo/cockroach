package repositories

import (
	"cockroach/cockroach/entities"
)

type AnimalRepository interface {
	InsertAnimalData(in *entities.InsertAnimalDto) error
}
