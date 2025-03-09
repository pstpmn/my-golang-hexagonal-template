package conf

import (
	"os"
)

type (
	Config struct {
		App          App   `mapstructure:"App"`
		Mongo        Mongo `mapstructure:"Mongo"`
		OrderRedis   Redis `mapstructure:"Redis"`
		AccountRedis Redis `mapstructure:"Redis"`
		Kafka        Kafka `mapstructure:"Kafka"`
	}

	App struct {
		Name string `mapstructure:"name"`
		Host string `mapstructure:"host"`
	}

	Mongo struct {
		Uri    string `mapstructure:"Uri"`
		DbName string `mapstructure:"dbName"`
	}

	Redis struct {
		Uri string `mapstructure:"Uri"`
	}

	Kafka struct {
		BrokerHost    []string
		User          string
		Pass          string
		PORT          string
		Topic         string
		Reset         string
		Mechanism     string
		Protocol      string
		Base64Cert    string
		ProducerTopic string
		ConsumerTopic string
	}
)

func LoadConfigFromEnv() *Config {
	return &Config{
		App: App{
			Name: os.Getenv("APP_NAME"),
			Host: os.Getenv("HOST"),
		},
		Mongo: Mongo{
			Uri:    os.Getenv("MONGO_URI"),
			DbName: os.Getenv("MONGO_DB_NAME"),
		},
		OrderRedis: Redis{
			Uri: os.Getenv("ORDER_REDIS_URI"),
		},
		AccountRedis: Redis{
			Uri: os.Getenv("ACCOUNT_REDIS_URI"),
		},
		Kafka: Kafka{
			BrokerHost:    []string{os.Getenv("KAFKA_HOST")},
			User:          os.Getenv("KAFKA_USER"),
			Pass:          os.Getenv("KAFKA_PASS"),
			PORT:          os.Getenv("KAFKA_PORT"),
			Reset:         os.Getenv("KAFKA_RESET"),
			Mechanism:     os.Getenv("KAFKA_MECHANISM"),
			Protocol:      os.Getenv("KAFKA_PROTOCOL"),
			Base64Cert:    os.Getenv("KAFKA_BASE64CERT"),
			ProducerTopic: os.Getenv("KAFKA_PRODUCER_TOPIC"),
			ConsumerTopic: os.Getenv("KAFKA_CONSUMER_TOPIC"),
		},
	}
}
