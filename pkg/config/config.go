package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

func ProvideConfig() Config {
	return Config{
		BasePath: requireEnv("BASE_PATH"),
		Groups:   getGroups(),
		Authentication: authentication{
			Keys: keys{
				PrivateKey: requireEnv("PRIVATE_KEY"),
				PublicKey:  requireEnv("PUBLIC_KEY"),
			},
			RefreshTokenSecretKey:         requireEnv("REFRESH_TOKEN_SECRET_KEY"),
			AccessTokenExpirationSeconds:  requireEnvAsInt("ACCESS_TOKEN_EXPIRATION_IN_SECONDS"),
			RefreshTokenExpirationSeconds: requireEnvAsInt("REFRESH_TOKEN_EXPIRATION_IN_SECONDS"),
		},
		Postgresql: postgresql{
			Host:         requireEnv("DATABASE_HOST"),
			Port:         requireEnvAsInt("DATABASE_PORT"),
			Username:     requireEnv("DATABASE_USERNAME"),
			Password:     requireEnv("DATABASE_PASSWORD"),
			DatabaseName: requireEnv("DATABASE_NAME"),
		},
		Redis: redis{
			Host: requireEnv("REDIS_HOST"),
			Port: requireEnvAsInt("REDIS_PORT"),
		},
		AdminUser: user{
			Email:    requireEnv("ADMIN_USER_EMAIL"),
			Password: requireEnv("ADMIN_USER_PASSWORD"),
		},
		ServiceUsers: getServiceUsers(),
		DefaultUser:  user{},
	}
}

func getGroups() []group {
	groupNames := requireEnvAsArray("GROUP_NAMES")
	groupHostnames := requireEnvAsArray("GROUP_HOSTNAMES")

	if len(groupNames) != len(groupHostnames) {
		log.Fatalln("len(GROUP_NAMES) != len(GROUP_HOSTNAMES)")
	}

	groups := make([]group, len(groupNames))
	for i := 0; i < len(groupNames); i++ {
		groups[i].Name = groupNames[i]
		groups[i].Hostname = groupHostnames[i]
	}

	return groups
}

func getServiceUsers() []user {
	userEmails := requireEnvAsArray("SERVICE_USER_EMAILS")
	userPasswords := requireEnvAsArray("SERVICE_USER_PASSWORDS")

	if len(userEmails) != len(userPasswords) {
		log.Fatalln("len(SERVICE_USER_EMAILS) != len(SERVICE_USER_PASSWORDS)")
	}

	users := make([]user, len(userEmails))
	for i := 0; i < len(userEmails); i++ {
		users[i].Email = userEmails[i]
		users[i].Password = userPasswords[i]
	}

	return users
}

type Config struct {
	BasePath       string
	Groups         []group
	Authentication authentication
	Postgresql     postgresql
	Redis          redis
	AdminUser      user
	ServiceUsers   []user
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

func (k keys) GetPrivateKey() (*rsa.PrivateKey, error) {
	decode, _ := pem.Decode([]byte(k.PrivateKey))
	if decode == nil {
		return nil, errors.New("failed to decode private key")
	}

	// Openssl generates keys formatted as PKCS8 but terraform tls_private_key is producing PKCS1
	// So if parsing of PKCS8 fails we try PKCS1
	privateKey, err := x509.ParsePKCS8PrivateKey(decode.Bytes)
	if err != nil {
		if err.Error() == "x509: failed to parse private key (use ParsePKCS1PrivateKey instead for this key format)" {
			log.Println("Trying to parse PKCS1 format...")
			privateKey, err = x509.ParsePKCS1PrivateKey(decode.Bytes)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
		log.Println("Successfully parsed private key")
	}

	return privateKey.(*rsa.PrivateKey), nil
}

func (k keys) GetPublicKey() (*rsa.PublicKey, error) {
	decode, _ := pem.Decode([]byte(k.PublicKey))
	if decode == nil {
		return nil, errors.New("failed to decode public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(decode.Bytes)
	if err != nil {
		return nil, err
	}

	return publicKey.(*rsa.PublicKey), nil
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

type group struct {
	Name     string
	Hostname string
}

func requireEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Can't find environment varialbe: %s\n", key)
	}
	return value
}

func requireEnvAsArray(key string) []string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Can't find environment varialbe: %s\n", key)
	}
	return strings.Split(value, ",")
}

func requireEnvAsInt(key string) int {
	valueStr := requireEnv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Can't parse value as integer: %s", err.Error())
	}
	return value
}
