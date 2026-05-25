package postgres

import (
	"context"

	"go-service-template/internal/repositories"
)

func (ir *ItemsRepository) GetItem(
	ctx context.Context,
	itemId int,
) (*repositories.Item, error) {
	const query = `SELECT id, content FROM items WHERE id = $1;`
	stmt, err := ir.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res repositories.Item
	if err := stmt.QueryRowContext(ctx, itemId).Scan(&res.Id, &res.Content); err != nil {
		return nil, err
	}

	return &res, nil
}
