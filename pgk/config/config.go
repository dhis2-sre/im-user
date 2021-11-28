package config

import (
	"log"
	"os"
	"strconv"
)

func ProvideConfig() Config {
	return Config{
		ServerPort: requireEnv("SERVER_PORT"),
		BasePath:   requireEnv("BASE_PATH"),
		Authentication: authentication{
			Keys: keys{
				PrivateKey: "",
				PublicKey:  "",
			},
			RefreshTokenSecretKey:         "",
			AccessTokenExpirationSeconds:  0,
			RefreshTokenExpirationSeconds: 0,
		},
		Postgresql: postgresql{
			Host:         requireEnv("DATABASE_HOST"),
			Port:         getEnvAsInt("DATABASE_PORT"),
			Username:     requireEnv("DATABASE_USERNAME"),
			Password:     requireEnv("DATABASE_PASSWORD"),
			DatabaseName: requireEnv("DATABASE_NAME"),
		},
		Redis:       redis{},
		AdminUser:   user{},
		DefaultUser: user{},
	}
}

type Config struct {
	ServerPort     string
	BasePath       string
	Authentication authentication
	Postgresql     postgresql
	Redis          redis
	AdminUser      user
	DefaultUser    user
}

type authentication struct {
	Keys                          keys
	RefreshTokenSecretKey         string
	AccessTokenExpirationSeconds  int
	RefreshTokenExpirationSeconds int
}

type keys struct {
	PrivateKey string
	PublicKey  string
}

type postgresql struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

type redis struct {
	Host string
	Port int
}

type user struct {
	Email    string
	Password string
}

func requireEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Can't find environment varialbe: %s\n", key)
	}
	return value
}

func getEnvAsInt(key string) int {
	valueStr := requireEnv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Can't parse value as integer: %s", err.Error())
	}
	return value
}
