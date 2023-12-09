package configs

import (
	"os"
	"time"
)

type Config struct {
	EchoPort    string
	JWTDuration time.Duration
	JWTSecret   []byte
}

func InitConfig() *Config {
	echoPort := GetEnv("ECHO_PORT", ":1323")
	jwtDuration := GetEnv("JWT_DURATION", "1h")
	jwtSecret := GetEnv("JWT_SECRET", "yFxZ69vilnFJ6FKwGmFHNHkaEcEs")

	duration, err := time.ParseDuration(jwtDuration)
	if err != nil {
		panic(err)
	}

	return &Config{
		EchoPort:    echoPort,
		JWTDuration: duration,
		JWTSecret:   []byte(jwtSecret),
	}
}

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
