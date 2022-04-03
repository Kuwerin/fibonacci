# Proto compiling stage
FROM namely/protoc-all as proto-compiler

WORKDIR /app

COPY proto/ /app/proto/

RUN mkdir -p compiled/fibonaccipb

RUN protoc --proto_path=/app/proto/:/opt/include/ \
    --go_out=/app/compiled/fibonaccipb \
    --go-grpc_out=/app/compiled/fibonaccipb \
    /app/proto/fibonacci_service.proto

# Build stage
FROM golang:1.17-stretch as build-stage


WORKDIR /app

COPY . /app

# Copy compiled files for geo service
COPY --from=proto-compiler /app/compiled/fibonaccipb /app/pkg/transport/grpc/fibonaccipb

ENV CGO_ENABLED=0
RUN go build -o /app/build/svc-fibonacci ./main.go

# Runtime stage
FROM alpine

COPY --from=build-stage /app/build/svc-fibonacci /bin/svc-fibonacci

EXPOSE 5010
EXPOSE 5000

ENTRYPOINT ["/bin/svc-fibonacci"]

