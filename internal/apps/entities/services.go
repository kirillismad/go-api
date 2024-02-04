package entities

import (
	"context"
	"errors"
	"go-api/internal/apps"

	"gorm.io/gorm"
)

type EntityService struct {
	db *gorm.DB
}

func NewEntityService(db *gorm.DB) EntityService {
	return EntityService{
		db: db,
	}
}

func (s EntityService) CreateEntity(ctx context.Context, data CreateEntityData) (CreateEntityResult, error) {
	var entity Entity
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// validate
		var cnt int64
		stmt := tx.Model(new(Entity)).Where("uuid_field = ?", data.UUIDField)
		if err := stmt.Count(&cnt).Error; err != nil {
			return errors.Join(apps.ErrDataSource, err)
		}
		if cnt != 0 {
			return apps.ErrUnique
		}
		// create
		entity = Entity{
			UUIDField:     data.UUIDField,
			IntField:      data.IntField,
			FloatField:    data.FloatField,
			DatetimeField: data.DatetimeField,
			StringField:   data.StringField,
			BoolField:     data.BoolField,
		}
		if err := tx.Create(&entity).Error; err != nil {
			return errors.Join(apps.ErrDataSource, err)
		}
		return nil
	})

	// handle error
	if err != nil {
		return CreateEntityResult{}, err
	}
	// map result
	return CreateEntityResult{ID: entity.ID}, nil
}
