package domain

import "time"

type OtpModel struct {
	Id        int
	OtpNumber string
	CreatedAt time.Time
	UpdatedAt time.Time
}
