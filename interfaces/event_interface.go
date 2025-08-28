package interfaces

import (
	"alumni-circle-api/models"
	"context"
)

type EventRepository interface {
	GetAllEvent(ctx context.Context, limit int, offset int, search string) ([]models.Event, int64, error)
	GetEventByCategory(ctx context.Context, IDKategori int64) ([]models.EventWithCategory, error)
	GetEventByID(ctx context.Context, IDEvent int64) (*models.Event, error)
}
