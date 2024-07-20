# goLang-message-service

### Технологии
* go@1.22.5: https://go.dev/doc/install
```bash
go mod init message-service
go get github.com/gin-gonic/gin
go get github.com/jackc/pgx/v4
go get github.com/joho/godotenv
go get github.com/segmentio/kafka-go
```
* postgresql@16.3: https://www.postgresql.org/download/
* kafka@3.7.1: https://kafka.apache.org/downloads.html
______________________________________________________

### хранения конфигурации: .env
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
KAFKA_BROKER=localhost:9092
KAFKA_TOPIC=messages