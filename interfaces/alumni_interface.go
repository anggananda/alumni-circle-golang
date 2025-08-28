package interfaces

import (
	"alumni-circle-api/models"
	"context"
)

type AlumniRepository interface {
	Register(ctx context.Context, username, password, email string) error
	GetAllAlumni(ctx context.Context, limit int, offset int, search string) ([]models.Alumni, int64, error)
	GetAlumniByUsername(ctx context.Context, username string) (*models.Alumni, error)
	GetAlumniByID(ctx context.Context, id int64) (*models.Alumni, error)
	DeleteAlumni(ctx context.Context, id int64) error
}
