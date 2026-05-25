package repositories

// Item is an example entity carried through the repository layer.
// Rename/reshape this (and its fields) to match your real domain object.
type Item struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
}

type GetItemsByOwnerIdDto struct {
	Items []Item `json:"items"`
}
