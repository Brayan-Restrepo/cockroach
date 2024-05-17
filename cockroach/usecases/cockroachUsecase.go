package usecases

import (
	"cockroach/cockroach/models"
)

type CockroachUsecase interface {
	CockroachDataProcessing(in *models.AddCockroachData) error
}
