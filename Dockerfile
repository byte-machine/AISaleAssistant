# Используем официальный образ Go как базовый
FROM golang:1.23.3-alpine as builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем go.mod и go.sum в контейнер для кэширования зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod tidy

# Копируем исходный код проекта в контейнер
COPY . .

# Копируем .env файл в контейнер (если он есть в проекте)
COPY .env .env

# Копируем папку с файлами (например, для работы с файлами)
COPY files files

# Собираем бинарный файл Go
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Используем более легкий образ для финальной сборки
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/  

# Копируем скомпилированный бинарник из предыдущего шага
COPY --from=builder /app/main .

# Копируем .env файл, если нужно использовать его в контейнере
COPY --from=builder /app/.env .env

# Копируем файлы из каталога "files" в контейнер
COPY --from=builder /app/files /app/files

# Открываем порт, который будет использовать сервис
EXPOSE 8081

# Запускаем приложение
CMD ["./main"]
