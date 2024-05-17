package usecases

import (
	"cockroach/cockroach/entities"
	"cockroach/cockroach/models"
	repositories2 "cockroach/cockroach/repositories"
	"time"
)

type cockroachUsecaseImpl struct {
	cockroachRepository repositories2.CockroachRepository
	cockroachMessaging  repositories2.CockroachMessaging
}

func NewCockroachUsecaseImpl(
	cockroachRepository repositories2.CockroachRepository,
	cockroachMessaging repositories2.CockroachMessaging,
) CockroachUsecase {
	return &cockroachUsecaseImpl{
		cockroachRepository: cockroachRepository,
		cockroachMessaging:  cockroachMessaging,
	}
}

func (u *cockroachUsecaseImpl) CockroachDataProcessing(in *models.AddCockroachData) error {
	insertCockroachData := &entities.InsertCockroachDto{
		Amount: in.Amount,
	}

	if err := u.cockroachRepository.InsertCockroachData(insertCockroachData); err != nil {
		return err
	}

	pushCockroachData := &entities.CockroachPushNotificationDto{
		Title:        "Cockroach Detected 🪳 !!!",
		Amount:       in.Amount,
		ReportedTime: time.Now().Local().Format("2006-01-02 15:04:05"),
	}

	if err := u.cockroachMessaging.PushNotification(pushCockroachData); err != nil {
		return err
	}

	return nil
}
