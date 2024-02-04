package entities

import "context"

type CreateEntityUseCase interface {
	CreateEntity(ctx context.Context, data CreateEntityData) (CreateEntityResult, error)
}
