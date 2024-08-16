package config

import (
	"os"
)

type Config struct {
	RabbitMQUser 	string
	RabbitMQHPass 	string
	RabbitMQHost 	string
	RabbitMQPort 	string
	RedisHost    	string
	RedisPort    	string
	RedisPassword	string
	QueueName    	string
	JwtSecret		string

}

func LoadConfig() *Config {
	return &Config{
		RabbitMQUser: 	getEnv("RABBITMQ_UNAME", "guest"),
		RabbitMQHPass: 	getEnv("RABBITMQ_PASS", "guest"),
		RabbitMQHost: 	getEnv("RABBITMQ_HOST", "rabbitmq"),
		RabbitMQPort: 	getEnv("RABBITMQ_PORT", "5672"),
		RedisHost:    	getEnv("REDIS_HOST", "redis"),
		RedisPort:    	getEnv("REDIS_PORT", "6379"),
		RedisPassword:	getEnv("REDIS_PASSWORD", ""),
		QueueName:    	getEnv("QUEUE_NAME", "test-job-queue"),
		JwtSecret:		getEnv("JWT_SECRET", "YOU_LIVE_ONLY_ONCE"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
