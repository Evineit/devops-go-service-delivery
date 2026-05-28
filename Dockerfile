FROM golang:1.26.1-alpine AS builder
WORKDIR /src

# Cache go modules (copy only go.mod; go.sum may be absent)
COPY go.mod ./
RUN go mod download || true

# copy rest of repo
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/app ./cmd/api

FROM alpine:3.18
RUN addgroup -S app && adduser -S -G app app
COPY --from=builder /app/bin/app /usr/local/bin/app
ENV PORT=8080
EXPOSE 8080
USER app
ENTRYPOINT ["/usr/local/bin/app"]
