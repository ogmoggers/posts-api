package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	network "social-network-api"
	"social-network-api/pkg/clients/kafka"
	"social-network-api/pkg/handler"
	"social-network-api/pkg/repository"
	"social-network-api/pkg/service"
	"strings"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: .env file not found, using environment variables")
	}

	db, err := repository.NewPostgres(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed when initializing database: %s", err.Error())
	}

	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	if kafkaBrokers == "" {
		kafkaBrokers = "kafka:9092" // Default Kafka broker
	}

	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	if kafkaTopic == "" {
		kafkaTopic = "posts" // Default Kafka topic
	}

	kafkaProducer, err := kafka.NewProducer(strings.Split(kafkaBrokers, ","), kafkaTopic)
	if err != nil {
		log.Printf("Warning: failed to initialize Kafka producer: %s", err.Error())
		kafkaProducer = nil
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, kafkaProducer)
	handlers := handler.NewHandler(services)

	fmt.Println("Starting API")
	srv := new(network.Server)
	port := viper.GetString("port")
	if port == "" {
		port = "8080"
	}
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
