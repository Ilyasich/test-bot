# Используем базовый образ Golang
FROM golang:1.20-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы в контейнер
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка приложения
RUN go build -o bot .

# Устанавливаем порт
ENV PORT=8080

# Запускаем приложение
CMD ["./bot"]
