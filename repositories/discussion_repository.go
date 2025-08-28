package repositories

import (
	"alumni-circle-api/interfaces"
	"alumni-circle-api/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type DiscussionMySQLRepository struct {
	DB *gorm.DB
}

func NewDiscussionMySQLRepository(db *gorm.DB) interfaces.DiscussionRepository {
	return &DiscussionMySQLRepository{
		DB: db,
	}
}

func (repo *DiscussionMySQLRepository) GetAllDiscussion(ctx context.Context, limit int, offset int, search string) ([]models.ListDiskusi, int64, error) {
	db := repo.DB.WithContext(ctx)

	query := db.Model(&models.Diskusi{})

	if search != "" {
		query = query.Where("subjek_diskusi LIKE ?", "%"+search+"%")
	}

	var totalItems int64
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count alumni: %w", err)
	}

	var diskusi []models.ListDiskusi
	// if err := query.
	// 	Limit(limit).
	// 	Offset(offset).
	// 	Find(&diskusi).Error; err != nil {
	// 	return nil, 0, fmt.Errorf("failed to fetch alumni: %w", err)
	// }

	if err := query.
		Table("diskusi").
		Select("diskusi.*, alumni.foto_profile, alumni.nama_alumni, alumni.email").
		Joins("INNER JOIN alumni ON diskusi.id_alumni = alumni.id_alumni").
		Limit(limit).
		Offset(offset).
		Find(&diskusi).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch alumni: %w", err)
	}

	return diskusi, totalItems, nil
}

func (repo *DiscussionMySQLRepository) GetDiscussionByID(ctx context.Context, IDDiskusi int64) (*models.Diskusi, error) {
	diskusi, err := gorm.G[models.Diskusi](repo.DB).Where("id_diskusi = ?", IDDiskusi).First(ctx)
	if err != nil {
		return nil, err
	}

	return &diskusi, nil
}
