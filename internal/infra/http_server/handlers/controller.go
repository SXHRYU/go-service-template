package http_server

import (
	"context"

	"go-service-template/internal/configs"
	items "go-service-template/internal/services"
)

// itemsService is the boundary the controller depends on. Keep it narrow:
// only the methods this controller actually calls, expressed against
// service-layer DTOs.
type itemsService interface {
	CreateItem(ctx context.Context, content string) (int, error)
	GetItem(ctx context.Context, itemId int) (*items.Item, error)
}

type ItemsController struct {
	itemsSrv itemsService
	config   *configs.Config
}

func NewItemsController(itemsSrv itemsService, config *configs.Config) *ItemsController {
	return &ItemsController{itemsSrv: itemsSrv, config: config}
}
