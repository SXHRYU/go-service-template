package http_server

import (
	"net/http"

	"github.com/oaswrap/spec/adapter/httpopenapi"
	"github.com/oaswrap/spec/option"
	handlers "go-service-template/internal/infra/http_server/handlers"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter(controller *handlers.ItemsController) *Router {
	mux := http.NewServeMux()
	r := httpopenapi.NewGenerator(
		mux,
		option.WithTitle("Service API"),
		option.WithVersion("0.0.1"),
	)

	api := r.Group("/api")
	api.HandleFunc("POST /create/", exampleMiddleware(controller.CreateItem)).With(
		option.Summary("Create item"),
		option.Request(new(CreateItemRequest)),
		option.Response(http.StatusCreated, new(handlers.CreateItemResponse)),
		option.Response(http.StatusBadRequest, new(ErrorResponse)),
		option.Response(http.StatusRequestEntityTooLarge, new(ErrorResponse)),
	)
	api.HandleFunc("GET /{id}", controller.GetItem).With(
		option.Summary("Get item"),
		option.Request(new(GetItemRequest)),
		option.Response(http.StatusOK, new(handlers.Item)),
		option.Response(http.StatusNotFound, new(ErrorResponse)),
	)

	return &Router{
		mux: mux,
	}
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(rw, req)
}

type PageLimitPaginated struct {
	Page  int `query:"page"  default:"1"`
	Limit int `query:"limit" default:"20"`
}

type GetItemRequest struct {
	Id int `path:"id" required:"true"`
}

type CreateItemRequest struct {
	Content string `required:"true" json:"content" minLength:"1" maxLength:"10000"`
}

type ErrorResponse struct {
	Error string
}
