# 1. Шаг сборки
FROM golang:1.23-alpine AS builder

# Установка обновлений и зависимостей для сборки приложения
RUN apk update && apk add --no-cache git

ADD . /tiny-url

WORKDIR /tiny-url

RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ./build/tiny ./cmd

# 2. Шаг запуска (минимальный образ)
FROM alpine:latest

# Установка необходимых для работы зависимостей (например, сертификатов для TLS)
RUN apk --no-cache add ca-certificates

# Установка рабочей директории внутри контейнера
WORKDIR /

# Копирование собранного приложения из контейнера сборки
COPY --from=builder /tiny-url/build/tiny .

# Копирование файлов конфигурации, если они есть
COPY  --from=builder /tiny-url/conf.yaml .

# Открытие порта, если приложение работает на определенном порту
EXPOSE 8080

# Команда запуска приложения
CMD ["./tiny"]
