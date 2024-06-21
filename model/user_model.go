package model

import (
	"time"

	"github.com/google/uuid"
)

type UserReq struct {
	ID           uuid.UUID
	FirstName    string `validate:"required min=2,max=20`
	LastName     string `validate:"required min=1,max=20`
	Password     string `validate:"required min=8`
	Email        string `validate:"email required`
	PhoneNumber  string
	Token        string
	UserType     string `validate:"requried,eq=ADMIN|eq=USER`
	RefreshToken string
	CreatedOn    time.Time
	CreatedBy    string
	UpdatedBy    string
}
