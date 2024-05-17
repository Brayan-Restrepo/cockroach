package entities

import "time"

type (
	InsertCockroachDto struct {
		Id        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
		Amount    uint32    `json:"amount"`
		CreatedAt time.Time `json:"createdAt"`
	}

	Cockroach struct {
		Id        uint32    `json:"id"`
		Amount    uint32    `json:"amount"`
		CreatedAt time.Time `json:"createdAt"`
	}

	CockroachPushNotificationDto struct {
		Title        string `json:"title"`
		Amount       uint32 `json:"amount"`
		ReportedTime string `json:"createdAt"`
	}

	InsertAnimalDto struct {
		Id        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
		Name      string    `json:"name"`
		Species   string    `json:"species"`
		Emoji     string    `json:"emoji"`
		CreatedAt time.Time `json:"createdAt"`
	}

	Animals struct {
		Id        uint32    `json:"id"`
		Name      string    `json:"name"`
		Species   string    `json:"species"`
		Emoji     string    `json:"emoji"`
		CreatedAt time.Time `json:"createdAt"`
	}
)
