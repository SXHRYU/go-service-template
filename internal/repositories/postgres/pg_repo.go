package postgres

type ItemsRepository struct {
	db *DB
}

func NewItemsRepository(db *DB) *ItemsRepository {
	return &ItemsRepository{
		db: db,
	}
}
