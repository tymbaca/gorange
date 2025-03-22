FROM golang:1.23.6-alpine3.21 AS build
WORKDIR /
COPY . .
RUN GOBIN=/ go install github.com/go-delve/delve/cmd/dlv@latest
RUN go build -gcflags "all=-N -l" -o ./app .

FROM alpine
COPY --from=build /dlv /dlv
COPY --from=build /app /app
ENTRYPOINT [ "/dlv" ]
