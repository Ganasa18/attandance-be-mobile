package domain

import "time"

type LogGeneralModel struct {
	Id        int
	LogBy     string
	ActionLog string
	CreatedAt time.Time
}
