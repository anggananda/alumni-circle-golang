package interfaces

import (
	"alumni-circle-api/models"
	"context"
)

type DiscussionRepository interface {
	GetAllDiscussion(ctx context.Context, limit int, offset int, search string) ([]models.ListDiskusi, int64, error)
  GetDiscussionByID(ctx context.Context, IDDiskusi int64)(*models.Diskusi, error)
}
