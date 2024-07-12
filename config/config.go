package config

import "github.com/spf13/viper"

type Config struct {
	Postgres    PostgresConfig
	KafkaConfig KafkaConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	URI      string
}

type KafkaConfigWriter struct {
	Brokers []string
	Topic   string
}

type KafkaConfigReader struct {
	Brokers   []string
	Topic     string
	Partition int
}

type KafkaConfig struct {
	KafkaConfigReader KafkaConfigReader
	KafkaConfigWriter KafkaConfigWriter
}

func SetConfig() *Config {
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	v := viper.GetViper()

	viper.AutomaticEnv()

	return &Config{
		Postgres: PostgresConfig{
			Host:     v.GetString("POSTGRES_HOST"),
			Port:     v.GetString("POSTGRES_PORT"),
			Database: v.GetString("POSTGRES_DATABASE"),
			User:     v.GetString("POSTGRES_USER"),
			Password: v.GetString("POSTGRES_PASSWORD"),
			URI:      v.GetString("POSTGRES_URI"),
		},
		KafkaConfig: KafkaConfig{
			KafkaConfigReader: KafkaConfigReader{
				Brokers:   v.GetStringSlice("KAFKA_BROKERS"),
				Topic:     v.GetString("KAFKA_TOPIC"),
				Partition: v.GetInt("KAFKA_PARTITION"),
			},
			KafkaConfigWriter: KafkaConfigWriter{
				Brokers: v.GetStringSlice("KAFKA_BROKERS"),
				Topic:   v.GetString("KAFKA_TOPIC"),
			},
		},
	}
}
