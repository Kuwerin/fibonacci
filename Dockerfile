FROM golang:latest as build

WORKDIR /go/src/app

COPY . .
# RUN go test -v -race -timeout 30s ./...

RUN mkdir .bin
RUN  go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go


FROM alpine:latest as go-bin

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=build /go/src/app/.bin/app .
COPY --from=build /go/src/app/configs/* /root/configs/
