package services

import (
	"alumni-circle-api/interfaces"
	"alumni-circle-api/models"
	"context"
)

type DiscussionService struct {
	DiscussionRepository interfaces.DiscussionRepository
}

func NewDiscussionService(repo interfaces.DiscussionRepository) *DiscussionService {
	return &DiscussionService{
		DiscussionRepository: repo,
	}
}

func (service *DiscussionService) GetAllDiscussion(ctx context.Context, limit int, offset int, search string) ([]models.ListDiskusi, int64, error) {
	return service.DiscussionRepository.GetAllDiscussion(ctx, limit, offset, search)
}

func (service *DiscussionService) GetDiscussionByID(ctx context.Context, IDDiskusi int64) (*models.Diskusi, error) {
	return service.DiscussionRepository.GetDiscussionByID(ctx, IDDiskusi)
}
