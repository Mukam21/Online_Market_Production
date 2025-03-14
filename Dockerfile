# Используем официальный образ Go в качестве базового
FROM golang:1.24.1-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

COPY . .

# Устанавливаем зависимости проекта (если есть go.mod)
RUN go mod tidy

# Сборка приложения
RUN go build -o online_market cmd/main.go

# Используем более легкий образ для финальной сборки
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированный бинарный файл из предыдущего этапа
COPY --from=builder /app/online_market .

# Открываем порт для приложения
EXPOSE 8080

# Команда для запуска приложения
CMD ["./online_market"]
