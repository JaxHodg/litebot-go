FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o github.com/JaxHodg/litebot-go

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build github.com/JaxHodg/litebot-go /litebot-go

# EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/litebot-go"]
