FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY functions ./
COPY manager ./
COPY modules ./
COPY state ./
COPY vendor ./

RUN go build -o /litebot-go

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /litebot-go /litebot-go

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/litebot-go"]