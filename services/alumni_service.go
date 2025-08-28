package services

import (
	"alumni-circle-api/interfaces"
	"alumni-circle-api/models"
	"context"
)

type AlumniService struct {
	AlumniRepository interfaces.AlumniRepository
}

func NewAlumniService(repo interfaces.AlumniRepository) *AlumniService {
	return &AlumniService{AlumniRepository: repo}
}

func (service *AlumniService) Register(ctx context.Context, username, password, email string) error {
	return service.AlumniRepository.Register(ctx, username, password, email)
}

func (service *AlumniService) GetAllAlumni(ctx context.Context, limit int, offset int, search string) ([]models.Alumni, int64, error) {
	return service.AlumniRepository.GetAllAlumni(ctx, limit, offset, search)
}

func (service *AlumniService) GetAlumniByUsername(ctx context.Context, username string) (*models.Alumni, error) {
	return service.AlumniRepository.GetAlumniByUsername(ctx, username)
}

func (service *AlumniService) GetAlumniByID(ctx context.Context, IDAlumni int64) (*models.Alumni, error) {
	return service.AlumniRepository.GetAlumniByID(ctx, IDAlumni)
}

func (service *AlumniService) DeleteAlumni(ctx context.Context, IDAlumni int64) error {
	return service.AlumniRepository.DeleteAlumni(ctx, IDAlumni)
}
