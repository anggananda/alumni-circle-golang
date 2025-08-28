package repositories

import (
	"alumni-circle-api/interfaces"
	"alumni-circle-api/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type AlumniMySQLRepository struct {
	DB *gorm.DB
}

func NewAlumniMySQLRepository(db *gorm.DB) interfaces.AlumniRepository {
	return &AlumniMySQLRepository{
		DB: db,
	}
}

func (repo *AlumniMySQLRepository) Register(ctx context.Context, username, password, email string) error {
	alumni := models.Alumni{
		Username:    username,
		Password:    password,
		Email:       email,
		Roles:       "umum",
		FotoProfile: "defualt.png",
	}

	if err := gorm.G[models.Alumni](repo.DB).Create(ctx, &alumni); err != nil {
		return err
	}

	return nil
}

func (repo *AlumniMySQLRepository) GetAllAlumni(ctx context.Context, limit int, offset int, search string) ([]models.Alumni, int64, error) {
	db := repo.DB.WithContext(ctx)

	query := db.Model(&models.Alumni{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	var totalItems int64
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count alumni: %w", err)
	}

	var alumni []models.Alumni
	if err := query.
		Limit(limit).
		Offset(offset).
		Find(&alumni).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch alumni: %w", err)
	}

	return alumni, totalItems, nil
}

func (repo *AlumniMySQLRepository) GetAlumniByUsername(ctx context.Context, username string) (*models.Alumni, error) {
	alumi, err := gorm.G[models.Alumni](repo.DB).Where("username = ?", username).First(ctx)

	if err != nil {
		return nil, err
	}

	return &alumi, nil
}

func (repo *AlumniMySQLRepository) GetAlumniByID(ctx context.Context, IDAlumni int64) (*models.Alumni, error) {
	alumni, err := gorm.G[models.Alumni](repo.DB).Where("id_alumni = ?", IDAlumni).First(ctx)
	if err != nil {
		return nil, err
	}

	return &alumni, nil
}

func (repo *AlumniMySQLRepository) DeleteAlumni(ctx context.Context, IDAlumni int64) error {
	_, err := gorm.G[models.Alumni](repo.DB).Where("id_alumni = ?", IDAlumni).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
