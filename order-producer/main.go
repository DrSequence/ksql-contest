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

type Order struct {
	OrderID    string      `json:"order_id"`
	UserID     string      `json:"user_id"`
	OrderTime  int64       `json:"order_time"`
	TotalPrice float64     `json:"total_price"`
	Products   []OrderItem `json:"products"`
}

type OrderItem struct {
	ProductID string  `json:"product_id"`
	Category  string  `json:"category"`
	Price     float64 `json:"price"`
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

func generateOrder(userID string, categories []string) Order {
	var order Order
	orderID := uuid.New().String()
	order.OrderID = orderID
	order.UserID = userID
	order.OrderTime = time.Now().Unix()
	order.Products = make([]OrderItem, 0)

	numProducts := rand.Intn(10) + 1

	for i := 0; i < numProducts; i++ {
		productID := uuid.New().String()
		category := getRandomElement(categories)
		price := rand.Float64() * 100.0

		orderItem := OrderItem{
			ProductID: productID,
			Category:  category,
			Price:     price,
		}

		order.Products = append(order.Products, orderItem)
		order.TotalPrice += price
	}

	return order
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

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		userID := getRandomElement(users)
		order := generateOrder(userID, categories)

		orderBytes, err := json.Marshal(order)
		if err != nil {
			log.Printf("Failed to marshal order: %s", err)
			continue
		}

		topic := "shop-orders"
		producer.Produce(&kafka.Message{
			Key:            []byte(userID),
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          orderBytes,
		}, nil)

		log.Println("Order sent successfully")
	}
}
