package domain

import "time"

type LocationAttandanceModel struct {
	Id           int
	LocationName string
	LocationLtd  string
	UserId       int
	CreatedAt    time.Time
}

type AttandanceModel struct {
	Id        int
	ClockIn   *time.Time
	ClockOut  *time.Time
	WorkHour  *time.Time
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FaceConfigModel struct {
	Id        int
	UserId    int
	FaceUser  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
