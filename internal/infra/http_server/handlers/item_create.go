package http_server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	response "go-service-template/internal/infra/http_server/response"
)

func (c *ItemsController) CreateItem(rw http.ResponseWriter, req *http.Request) {
	// TODO: implement using structs and a validation library
	const size = 10 << 10

	body := http.MaxBytesReader(rw, req.Body, size)
	defer body.Close()

	var r struct{ Content string }

	if err := json.NewDecoder(body).Decode(&r); err != nil {
		if _, ok := errors.AsType[*http.MaxBytesError](err); ok {
			response.WriteErrorResponse(
				rw,
				fmt.Sprintf("`content` size is bigger than allowed (%d KiB)", size>>10),
				http.StatusRequestEntityTooLarge,
			)
			return
		}
		response.WriteErrorResponse(
			rw,
			fmt.Sprintf("failed to decode body: %v", err),
			http.StatusBadRequest,
		)
		return
	}

	ctx, cancel := context.WithTimeout(
		req.Context(),
		time.Duration(c.config.Http.Timeout)*time.Second,
	)
	defer cancel()

	itemId, err := c.itemsSrv.CreateItem(ctx, r.Content)
	if err != nil {
		response.WriteErrorResponse(
			rw,
			fmt.Sprintf("error while creating item: %v", err),
			http.StatusBadRequest,
		)
		return
	}

	response.WriteResponse(rw, ToResponseCreateItem(itemId), http.StatusCreated)
}
