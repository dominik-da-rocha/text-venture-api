# build backend
FROM  golang:bullseye AS gobuild

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY ./main.go ./
COPY ./cmd/    ./cmd
COPY ./config/ ./config

RUN go build -o ./text-venture-service main.go

# Deploy
FROM debian:bullseye-slim
WORKDIR /opt/text-venture
COPY --from=gobuild /app/text-venture-service ./
COPY --from=gobuild /app/config               ./config
EXPOSE 8081

ENTRYPOINT ["/opt/text-venture/text-venture-service"]

