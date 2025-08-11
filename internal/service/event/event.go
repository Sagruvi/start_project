package service_event

import (
	"context"
	model_event "start/internal/model/event"
)

type EventUsecase interface {
	GetEvent(ctx context.Context, id int) (model_event.Event, error)
	CreateEvent(ctx context.Context, event model_event.Event) (model_event.Event, error)
	DeleteEvent(ctx context.Context, id int) error
	UpdateEvent(ctx context.Context, event model_event.Event) (model_event.Event, error)
}

type service struct {
	uc EventUsecase
}

func New(uc EventUsecase) *service {
	return &service{uc: uc}
}

func (s *service) GetEvent(ctx context.Context, id int) (model_event.Event, error) {
	return s.uc.GetEvent(ctx, id)
}

func (s *service) CreateEvent(ctx context.Context, event model_event.Event) (model_event.Event, error) {
	return s.uc.CreateEvent(ctx, event)
}

func (s *service) DeleteEvent(ctx context.Context, id int) error {
	return s.uc.DeleteEvent(ctx, id)
}

func (s *service) UpdateEvent(ctx context.Context, event model_event.Event) (model_event.Event, error) {
	return s.uc.UpdateEvent(ctx, event)
}
