package http_server

import (
	items "go-service-template/internal/services"
)

// not a lot of records expected, so plain page-based pagination is fine
type PagedPagination struct {
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Total int `json:"total"`
}

type Item struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type GetOwnerItemsResponse struct {
	Items      []Item          `json:"items"`
	Pagination PagedPagination `json:"pagination"`
}

type CreateItemResponse struct {
	Id int `json:"id"`
}

func ToResponseOwnerItems(
	serviceDto *items.GetItemsByOwnerIdDto,
) *GetOwnerItemsResponse {
	res := make([]Item, len(serviceDto.Items))
	for i := range serviceDto.Items {
		res[i] = Item(serviceDto.Items[i])
	}
	return &GetOwnerItemsResponse{
		Items:      res,
		Pagination: PagedPagination(serviceDto.Pagination),
	}
}

func ToResponseCreateItem(newItemId int) *CreateItemResponse {
	return &CreateItemResponse{
		Id: newItemId,
	}
}

func ToResponseGetItem(item *items.Item) *Item {
	return &Item{
		Id:      item.Id,
		Content: item.Content,
	}
}
