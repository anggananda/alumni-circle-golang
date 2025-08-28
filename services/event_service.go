package services

import (
	"alumni-circle-api/interfaces"
	"alumni-circle-api/models"
	"context"
)

type EventService struct {
	EventRepository interfaces.EventRepository
}

func NewEventService(repo interfaces.EventRepository) *EventService {
	return &EventService{
		EventRepository: repo,
	}
}

func (service *EventService) GetAllEvent(ctx context.Context, limit int, offset int, search string) ([]models.Event, int64, error) {
	return service.EventRepository.GetAllEvent(ctx, limit, offset, search)
}

func (service *EventService) GetEventByCategory(ctx context.Context, IDKategori int64) ([]models.EventWithCategory, error) {
	return service.EventRepository.GetEventByCategory(ctx, IDKategori)
}

func (service *EventService) GetEventByID(ctx context.Context, IDEvent int64) (*models.Event, error) {
	return service.EventRepository.GetEventByID(ctx, IDEvent)
}
