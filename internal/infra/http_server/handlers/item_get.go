package http_server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	response "go-service-template/internal/infra/http_server/response"
)

func (c *ItemsController) GetItem(w http.ResponseWriter, req *http.Request) {
	itemId, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		// TODO: replace with TryProcessSlug(w, req)
		response.WriteErrorResponse(w, "not found", http.StatusNotFound)
		return
	}

	ctx, cancel := context.WithTimeout(
		req.Context(),
		time.Duration(c.config.Http.Timeout)*time.Second,
	)
	defer cancel()

	item, err := c.itemsSrv.GetItem(ctx, itemId)
	if err != nil {
		response.WriteErrorResponse(
			w,
			fmt.Sprintf("error while fetching item: %v", err),
			http.StatusNotFound,
		)
		return
	}
	response.WriteResponse(w, ToResponseGetItem(item), http.StatusOK)
}
