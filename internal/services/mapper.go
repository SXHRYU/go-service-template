package items

import (
	"go-service-template/internal/repositories"
)

func ToServiceOwnerItems(
	repoDto *repositories.GetItemsByOwnerIdDto,
	page, pages, total int,
) *GetItemsByOwnerIdDto {
	items := make([]Item, len(repoDto.Items))
	for i := range repoDto.Items {
		items[i] = Item(repoDto.Items[i])
	}
	return &GetItemsByOwnerIdDto{
		Items: items,
		Pagination: PagedPagination{
			Page:  page,
			Pages: pages,
			Total: total,
		},
	}
}

func ToServiceItem(item *repositories.Item) *Item {
	return &Item{
		Id:        item.Id,
		Content:   item.Content,
	}
}
