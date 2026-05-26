FROM golang:1.27.0-alpine3.24 AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd cmd
COPY internal internal
RUN go build -o /bin/app_server cmd/server/main.go

FROM base AS test
CMD [ "go", "test", "./..." ]

FROM base AS ci
RUN go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

FROM alpine:3.24 AS prod
COPY --from=base /bin/app_server /bin/app_server
CMD ["/bin/app_server"]
