package dto

import (
	"github.com/google/uuid"
	"time"
)

type RefreshTokenData struct {
	SS        string
	Id        uuid.UUID
	ExpiresIn time.Duration
}
