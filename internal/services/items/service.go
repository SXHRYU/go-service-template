package items

import (
	"context"

	"go-service-template/internal/repositories"
)

// repository is the boundary the service depends on. Keep it narrow: only
// the methods this service actually calls, expressed against repository DTOs.
type repository interface {
	CreateItem(ctx context.Context, content string) (int, error)
	GetItem(ctx context.Context, itemId int) (*repositories.Item, error)
}

type ItemsService struct {
	repo repository
}

func NewItemsService(repo repository) *ItemsService {
	return &ItemsService{repo: repo}
}
