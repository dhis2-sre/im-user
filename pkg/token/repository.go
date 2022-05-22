package token

import (
	"fmt"
	"log"
	"time"

	"github.com/dhis2-sre/im-user/internal/apperror"
	"github.com/go-redis/redis"
)

type Repository interface {
	SetRefreshToken(userId uint, tokenId string, expiresIn time.Duration) error
	DeleteRefreshToken(userId uint, previousTokenId string) error
	DeleteRefreshTokens(userId uint) error
}

func ProvideTokenRepository(redisClient *redis.Client) Repository {
	return &redisTokenRepository{
		Redis: redisClient,
	}
}

type redisTokenRepository struct {
	Redis *redis.Client
}

func (r redisTokenRepository) SetRefreshToken(userId uint, tokenId string, expiresIn time.Duration) error {
	key := fmt.Sprintf("%d:%s", userId, tokenId)
	if err := r.Redis.Set(key, 0, expiresIn).Err(); err != nil {
		message := fmt.Sprintf("Could not SET refresh token to redis for userId/tokenId: %d/%s: %s", userId, tokenId, err)
		return apperror.NewInternal(message)
	}
	return nil
}

func (r redisTokenRepository) DeleteRefreshToken(userId uint, previousTokenId string) error {
	key := fmt.Sprintf("%d:%s", userId, previousTokenId)

	result := r.Redis.Del(key)

	if err := result.Err(); err != nil {
		message := fmt.Sprintf("Could not delete refresh token to redis for userId/tokenId: %d/%s: %s", userId, previousTokenId, err)
		return apperror.NewInternal(message)
	}

	if result.Val() < 1 {
		log.Printf("Refresh token to redis for userId/tokenId: %d/%s does not exist\n", userId, previousTokenId)
		return apperror.NewUnauthorized("Invalid refresh token")
	}

	return nil
}

func (r redisTokenRepository) DeleteRefreshTokens(userId uint) error {
	pattern := fmt.Sprintf("%d*", userId)

	iterator := r.Redis.Scan(0, pattern, 5).Iterator()
	failCount := 0

	for iterator.Next() {
		if err := r.Redis.Del(iterator.Val()).Err(); err != nil {
			log.Printf("Failed to delete refresh token: %s\n", iterator.Val())
			failCount++
		}
	}

	if err := iterator.Err(); err != nil {
		message := fmt.Sprintf("Failed to delete refresh token: %s\n", iterator.Val())
		return apperror.NewInternal(message)
	}

	if failCount > 0 {
		message := fmt.Sprintf("Failed to delete refresh token: %s", iterator.Val())
		return apperror.NewInternal(message)
	}

	return nil
}
