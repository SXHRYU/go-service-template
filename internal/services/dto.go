package items

type PagedPagination struct {
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Total int `json:"total"`
}

// Item is the service-layer representation of the example entity.
// Rename/reshape this to match your real domain object.
type Item struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type GetItemsByOwnerIdDto struct {
	Items      []Item          `json:"items"`
	Pagination PagedPagination `json:"pagination"`
}
