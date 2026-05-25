package postgres

import "context"

func (ir *ItemsRepository) CreateItem(
	ctx context.Context,
	content string,
) (int, error) {
	query := "INSERT INTO items(content) VALUES($1) RETURNING id;"
	stmt, err := ir.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var newId int
	if err := stmt.QueryRowContext(ctx, content).Scan(&newId); err != nil {
		return 0, err
	}
	return newId, nil
}
