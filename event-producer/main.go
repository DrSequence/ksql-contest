package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

const topic = "product_view_events"

type Event struct {
	ProductID string `json:"product_id"`
	Timestamp int64  `json:"timestamp"`
	Category  string `json:"category"`
	App       string `json:"app"`
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
}

func readCSVFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.Read()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func getRandomElement(slice []string) string {
	index := rand.Intn(len(slice))
	return slice[index]
}

func main() {
	categories, err := readCSVFile("../mock_data/categories.csv")
	if err != nil {
		log.Fatalf("Failed to read categories file: %s", err)
	}

	users, err := readCSVFile("../mock_data/uuids.csv")
	if err != nil {
		log.Fatalf("Failed to read users file: %s", err)
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9093"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer producer.Close()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		user := getRandomElement(users)

		event := Event{
			ProductID: uuid.New().String(),
			Timestamp: time.Now().Unix(),
			Category:  getRandomElement(categories),
			App:       "GoShop",
			SessionID: uuid.New().String(),
			UserID:    user,
		}

		eventBytes, err := json.Marshal(event)
		if err != nil {
			log.Printf("Failed to marshal event: %s", err)
			continue
		}

		topic := topic
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(user),
			Value:          eventBytes,
		}, nil)

		log.Println("Message sent successfully")
	}
}
