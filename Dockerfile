FROM golang:1.18-alpine AS base
RUN go version
ENV GOPATH=/
RUN apk update
RUN apk add postgresql-client
WORKDIR /api
EXPOSE 8080

FROM base AS build
COPY ./ ./
RUN go mod download
RUN chmod +x wait-for-postgres.sh
RUN env GOOS=linux go build -o ./api ./cmd/main.go

FROM build AS final
COPY --from=build ./api ./

CMD ["./api"]