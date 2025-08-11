package usecase_event

import (
	"context"
	model_event "start/internal/model/event"
)

type EventRepository interface {
	GetEvent(ctx context.Context, id int) (model_event.Event, error)
	CreateEvent(ctx context.Context, event model_event.Event) (model_event.Event, error)
	DeleteEvent(ctx context.Context, id int) error
	UpdateEvent(ctx context.Context, event model_event.Event) (model_event.Event, error)
}

type EventUsecase struct {
	repo EventRepository
}

func New(repo EventRepository) *EventUsecase {
	return &EventUsecase{repo: repo}
}

func (u *EventUsecase) GetEvent(ctx context.Context, id int) (model_event.Event, error) {
	return u.repo.GetEvent(ctx, id)
}

func (u *EventUsecase) CreateEvent(ctx context.Context, event model_event.Event) (model_event.Event, error) {
	return u.repo.CreateEvent(ctx, event)
}

func (u *EventUsecase) DeleteEvent(ctx context.Context, id int) error {
	return u.repo.DeleteEvent(ctx, id)
}

func (u *EventUsecase) UpdateEvent(ctx context.Context, event model_event.Event) (model_event.Event, error) {
	return u.repo.UpdateEvent(ctx, event)
}
