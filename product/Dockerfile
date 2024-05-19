FROM golang:1.21.4-alpine3.18 AS builder
WORKDIR /src
COPY . .
RUN go  build -v -ldflags="-w -s" -o server main.go

FROM golang:1.21.4-alpine3.18
WORKDIR /app
RUN apk add --no-cache tzdata && \
   apk --no-cache add curl && \
   apk add busybox && \
   addgroup -g 1000 -S appgroup && \
   adduser -u 1000 -S appuser -G appgroup
ENV TZ=Asia/Almaty
COPY internal/migrations /app/internal/migrations
COPY --from=builder /src/server .
COPY --from=builder /src/.env.example .
RUN chmod +x /app/server && chown -R appuser:appgroup /app

USER appuser:appgroup
