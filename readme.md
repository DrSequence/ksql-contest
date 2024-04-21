## Microservices Architecture with Go and Java

This repository contains three microservices developed using Go and Java languages, integrated with Apache Kafka for event processing and messaging.

### Services

1. **Event Generator (Go)**
   - This service is written in Go and is responsible for generating "product view" events and sending them to a Kafka topic.
   - The events include information about the viewed product, category, application, and session ID.

2. **Order Creator (Go)**
   - Also implemented in Go, this service creates orders based on user product views.
   - It subscribes to product view events from Kafka, then aggregates multiple products into a single order.

3. **Purchase Finder (Java)**
   - Developed in Java using Kafka Streams, this service analyzes view events and searches for purchases related to product views.
   - It processes product view events and order data from Kafka, identifies cases where a user viewed a specific product category before making a purchase, and sends relevant messages to a new Kafka topic.

### Components

- **Apache Kafka**
  - Used for asynchronous message processing and event communication between microservices.
  - Services send and receive messages through Kafka topics to ensure scalability and fault tolerance.

### Dependencies

- **Go**
  - The Go services use standard library and third-party libraries like `confluentinc/confluent-kafka-go` for Kafka interaction.

- **Java**
  - The Java service utilizes `spring-kafka` for Kafka integration and `kafka-streams` for real-time data processing.

### How to Run

1. **Install Apache Kafka**
   - Ensure Apache Kafka is installed and running, including a Kafka broker and ZooKeeper.

2. **Configure and Run the Services**
   - For each service (Event Generator, Order Creator, Purchase Finder), follow these steps:
     - Install any necessary dependencies if not already installed.
     - Configure the configuration files (e.g., Kafka connection parameters).
     - Start each service using instructions provided in the respective README files within the service directories.

### Additional Resources

- [Apache Kafka Documentation](https://kafka.apache.org/documentation/)
- [Confluent Kafka Go Documentation](https://docs.confluent.io/platform/current/clients/confluent-kafka-go/html/index.html)
- [Spring Kafka Documentation](https://docs.spring.io/spring-kafka/docs/current/reference/html/index.html)
- [Kafka Streams Documentation](https://kafka.apache.org/documentation/streams/)
