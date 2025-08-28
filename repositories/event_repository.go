package repositories

import (
	"alumni-circle-api/interfaces"
	"alumni-circle-api/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type EventMySQLRepository struct {
	DB *gorm.DB
}

func NewEventMySQLRepository(db *gorm.DB) interfaces.EventRepository {
	return &EventMySQLRepository{
		DB: db,
	}
}

func (repo *EventMySQLRepository) GetAllEvent(ctx context.Context, limit int, offset int, search string) ([]models.Event, int64, error) {
	db := repo.DB.WithContext(ctx)

	query := db.Model(&models.Event{})

	if search != "" {
		query = query.Where("nama_event LIKE ?", "%"+search+"%")
	}

	var totalItems int64
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count alumni: %w", err)
	}

	var event []models.Event
	if err := query.
		Limit(limit).
		Offset(offset).
		Find(&event).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch alumni: %w", err)
	}

	return event, totalItems, nil
}

func (repo *EventMySQLRepository) GetEventByCategory(ctx context.Context, IDKategori int64) ([]models.EventWithCategory, error) {
	var events []models.EventWithCategory

	err := repo.DB.WithContext(ctx).
		Table("event").
		Select("event.id_event, event.id_kategori, event.nama_event, event.tanggal_event, event.lokasi, event.deskripsi, event.gambar, event.latitude, event.longitude, kategori.kategori").
		Joins("INNER JOIN kategori ON event.id_kategori = kategori.id_kategori").
		Where("event.id_kategori = ?", IDKategori).
		Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (repo *EventMySQLRepository) GetEventByID(ctx context.Context, IDEvent int64) (*models.Event, error) {
	event, err := gorm.G[models.Event](repo.DB).Where("id_event = ?", IDEvent).First(ctx)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
