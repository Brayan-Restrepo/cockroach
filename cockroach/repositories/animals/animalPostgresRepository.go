package repositories

import (
	"cockroach/cockroach/entities"
	"cockroach/database"
	"github.com/labstack/gommon/log"
)

type animalPostgresRepository struct {
	db database.Database
}

func NewAnimalPostgresRepository(db database.Database) AnimalRepository {
	return &animalPostgresRepository{db: db}
}

func (r *animalPostgresRepository) InsertAnimalData(in *entities.InsertAnimalDto) error {
	data := &entities.Animals{
		Name:    in.Name,
		Species: in.Species,
		Emoji:   in.Emoji,
	}

	result := r.db.GetDb().Create(data)

	if result.Error != nil {
		log.Errorf("InsertAnimalData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertAnimalData: %v", result.RowsAffected)
	return nil
}
