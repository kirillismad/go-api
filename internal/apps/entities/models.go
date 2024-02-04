package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Entity struct {
	ID            int            `gorm:"primaryKey;autoIncrement"`
	UUIDField     uuid.UUID      `gorm:"column:uuid_field;not null;unique;"`
	IntField      int            `gorm:"column:int_field;not null;"`
	FloatField    float64        `gorm:"column:float_field;not null;"`
	DatetimeField time.Time      `gorm:"column:datetime_field;not null;"`
	StringField   string         `gorm:"column:string_field;not null;size:256;"`
	BoolField     bool           `gorm:"column:bool_field;not null;"`
	JsonField     datatypes.JSON `gorm:"column:json_field;"`
}

func (e Entity) TableName() string {
	return "entities"
}
