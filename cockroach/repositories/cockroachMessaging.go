package repositories

import (
	"cockroach/cockroach/entities"
)

type CockroachMessaging interface {
	PushNotification(m *entities.CockroachPushNotificationDto) error
}
