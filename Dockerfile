##
## Build
##

FROM golang:1.17-bullseye as build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o ./bin/app main.go

##
## Deploy
##

FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY --from=build /build/bin/app /
COPY --from=build /build/app.yaml /

CMD ["/app"]