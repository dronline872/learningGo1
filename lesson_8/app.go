package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	port       = flag.Int("port", 8080, "Номер порта")
	db_url     = flag.String("db-url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "Ссылка на DB")
	jaeger_url = flag.String("jaeger-url", "http://jaeger:16686", "Ссылка на jaeger")
	sentry_url = flag.String("sentry-url", "http://sentry:9000", "Ссылка на sentry")
)

type App struct {
	Port        int    `json:"port"`
	KafkaBroker string `json:"kafka_broker"`
	SomeAppId   string `json:"some_app_id"`
	SomeAppKey  string `json:"some_app_key"`
}

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

	//Урок 9
	file, err := os.Open("conf.json")
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	app := App{}

	err = json.NewDecoder(file).Decode(&app)
	if err != nil {
		log.Printf("Не могу декодировать json-файл в структуру: %v", err)
	}

	fmt.Println(app.Port)
}
