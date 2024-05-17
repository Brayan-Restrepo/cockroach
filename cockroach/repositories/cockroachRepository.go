package repositories

import (
	"cockroach/cockroach/entities"
)

type CockroachRepository interface {
	InsertCockroachData(in *entities.InsertCockroachDto) error
}
