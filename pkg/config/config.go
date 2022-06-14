package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func New() (Config, error) {
	basePath, err := requireEnv("BASE_PATH")
	if err != nil {
		return Config{}, err
	}

	groups, err := newGroups()
	if err != nil {
		return Config{}, err
	}

	auth, err := newAuthentication()
	if err != nil {
		return Config{}, err
	}

	pg, err := newPostgresql()
	if err != nil {
		return Config{}, err
	}

	rd, err := newRedis()
	if err != nil {
		return Config{}, err
	}

	admin, err := newAdminUser()
	if err != nil {
		return Config{}, err
	}

	serviceUsers, err := newServiceUsers()
	if err != nil {
		return Config{}, err
	}

	return Config{
		BasePath:       basePath,
		Groups:         groups,
		Authentication: auth,
		Postgresql:     pg,
		Redis:          rd,
		AdminUser:      admin,
		ServiceUsers:   serviceUsers,
		DefaultUser:    user{},
	}, nil
}

type group struct {
	Name     string
	Hostname string
}

func newGroups() ([]group, error) {
	groupNames, err := requireEnvAsArray("GROUP_NAMES")
	if err != nil {
		return nil, err
	}
	groupHostnames, err := requireEnvAsArray("GROUP_HOSTNAMES")
	if err != nil {
		return nil, err
	}

	if len(groupNames) != len(groupHostnames) {
		return nil, errors.New("len(GROUP_NAMES) != len(GROUP_HOSTNAMES)")
	}

	groups := make([]group, len(groupNames))
	for i := 0; i < len(groupNames); i++ {
		groups[i].Name = groupNames[i]
		groups[i].Hostname = groupHostnames[i]
	}

	return groups, nil
}

type user struct {
	Email    string
	Password string
}

func newAdminUser() (user, error) {
	email, err := requireEnv("ADMIN_USER_EMAIL")
	if err != nil {
		return user{}, err
	}
	pw, err := requireEnv("ADMIN_USER_PASSWORD")
	if err != nil {
		return user{}, err
	}

	return user{
		Email:    email,
		Password: pw,
	}, nil
}

func newServiceUsers() ([]user, error) {
	emails, err := requireEnvAsArray("SERVICE_USER_EMAILS")
	if err != nil {
		return nil, err
	}
	pws, err := requireEnvAsArray("SERVICE_USER_PASSWORDS")
	if err != nil {
		return nil, err
	}

	if len(emails) != len(pws) {
		return nil, errors.New("len(SERVICE_USER_EMAILS) != len(SERVICE_USER_PASSWORDS)")
	}

	users := make([]user, len(emails))
	for i := 0; i < len(emails); i++ {
		users[i].Email = emails[i]
		users[i].Password = pws[i]
	}

	return users, nil
}

type authentication struct {
	Keys                          keys
	RefreshTokenSecretKey         string
	AccessTokenExpirationSeconds  int
	RefreshTokenExpirationSeconds int
}

func newAuthentication() (authentication, error) {
	privateKey, err := requireEnv("PRIVATE_KEY")
	if err != nil {
		return authentication{}, err
	}
	publicKey, err := requireEnv("PUBLIC_KEY")
	if err != nil {
		return authentication{}, err
	}

	refreshTokenSecretKey, err := requireEnv("REFRESH_TOKEN_SECRET_KEY")
	if err != nil {
		return authentication{}, err
	}
	accessTokenExpirationSeconds, err := requireEnvAsInt("ACCESS_TOKEN_EXPIRATION_IN_SECONDS")
	if err != nil {
		return authentication{}, err
	}
	refreshTokenExpirationSeconds, err := requireEnvAsInt("REFRESH_TOKEN_EXPIRATION_IN_SECONDS")
	if err != nil {
		return authentication{}, err
	}

	return authentication{
		Keys: keys{
			PrivateKey: privateKey,
			PublicKey:  publicKey,
		},
		RefreshTokenSecretKey:         refreshTokenSecretKey,
		AccessTokenExpirationSeconds:  accessTokenExpirationSeconds,
		RefreshTokenExpirationSeconds: refreshTokenExpirationSeconds,
	}, nil
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

func newPostgresql() (postgresql, error) {
	host, err := requireEnv("DATABASE_HOST")
	if err != nil {
		return postgresql{}, err
	}
	port, err := requireEnvAsInt("DATABASE_PORT")
	if err != nil {
		return postgresql{}, err
	}
	usrname, err := requireEnv("DATABASE_USERNAME")
	if err != nil {
		return postgresql{}, err
	}
	pw, err := requireEnv("DATABASE_PASSWORD")
	if err != nil {
		return postgresql{}, err
	}
	name, err := requireEnv("DATABASE_NAME")
	if err != nil {
		return postgresql{}, err
	}

	return postgresql{
		Host:         host,
		Port:         port,
		Username:     usrname,
		Password:     pw,
		DatabaseName: name,
	}, nil
}

type redis struct {
	Host string
	Port int
}

func newRedis() (redis, error) {
	host, err := requireEnv("REDIS_HOST")
	if err != nil {
		return redis{}, err
	}
	port, err := requireEnvAsInt("REDIS_PORT")
	if err != nil {
		return redis{}, err
	}

	return redis{
		Host: host,
		Port: port,
	}, nil
}

func requireEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("can't find environment variable: %s", key)
	}

	return value, nil
}

func requireEnvAsArray(key string) ([]string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return nil, fmt.Errorf("can't find environment variable: %s", key)
	}

	return strings.Split(value, ","), nil
}

func requireEnvAsInt(key string) (int, error) {
	valueStr, err := requireEnv(key)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("can't parse value as integer: %v", err)
	}

	return value, nil
}
