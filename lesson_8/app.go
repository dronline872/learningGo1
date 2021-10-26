package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	port       = flag.Int("port", 8080, "Номер порта")
	db_url     = flag.String("db-url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "Ссылка на DB")
	jaeger_url = flag.String("jaeger-url", "http://jaeger:16686", "Ссылка на jaeger")
	sentry_url = flag.String("sentry-url", "http://sentry:9000", "Ссылка на sentry")
)

func main() {
	os.Setenv("kafka_broker", "kafka:9092")
	os.Setenv("some_app_id", "testid")
	os.Setenv("some_app_key", "testkey")

	flag.Parse()
	fmt.Println("port: ", *port)
	fmt.Println("db-url: ", *db_url)
	fmt.Println("jaeger-url: ", *jaeger_url)
	fmt.Println("sentry-url: ", *sentry_url)
	fmt.Println("kafka_broker: ", os.Getenv("kafka_broker"))
	fmt.Println("some_app_id: ", os.Getenv("some_app_id"))
	fmt.Println("some_app_key: ", os.Getenv("some_app_key"))
}
