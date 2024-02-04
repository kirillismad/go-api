package entities

import (
	"time"

	"github.com/google/uuid"
)

type CreateEntityData struct {
	UUIDField     uuid.UUID
	IntField      int
	FloatField    float64
	DatetimeField time.Time
	StringField   string
	BoolField     bool
	JsonField     []string
}

type CreateEntityResult struct {
	ID int
}
