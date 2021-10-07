package repository

import (
	"context"

	"github.com/service_content/database"
	"github.com/service_content/model"
)

type (
	EventStore interface {
		CreateEvent(ctx context.Context, eventParam *model.EventParam) error
	}

	evnUsecase struct {
		ev database.EventStoreRepository
	}
)

func NewEventEntity(evn database.EventStoreRepository) *evnUsecase {
	return &evnUsecase{
		ev: evn,
	}
}

func (env *evnUsecase) CreateEvent(ctx context.Context, evp *model.EventParam) error {
	return env.CreateEvent(ctx, evp)
}
