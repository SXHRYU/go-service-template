package items

import "context"

func (s *ItemsService) CreateItem(
	ctx context.Context,
	content string,
) (int, error) {
	// TODO: put your business validation / rules here before persisting.
	return s.repo.CreateItem(ctx, content)
}
