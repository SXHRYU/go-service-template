package items

import (
	"context"

	items "go-service-template/internal/services"
)

func (s *ItemsService) GetItem(ctx context.Context, itemId int) (*items.Item, error) {
	response, err := s.repo.GetItem(ctx, itemId)
	if err != nil {
		return nil, err // TODO: map to service-level errors
	}
	return items.ToServiceItem(response), nil
}
