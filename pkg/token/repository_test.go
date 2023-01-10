package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_redisTokenRepository_deleteRefreshToken_Happy(t *testing.T) {
	id := uint(1)
	previousTokenId := ""
	key := fmt.Sprintf("%d:%s", id, previousTokenId)

	r := &redisMock{}
	r.
		On("Del", []string{key}).
		Return(redis.NewIntCmd())

	repository := NewRepository(r)
	err := repository.deleteRefreshToken(id, "")

	// TODO: How do I set the val property of *redis.IntCmd ?
	//	require.NoError(t, err)
	assert.Equal(t, "invalid refresh token", err.Error())

	r.AssertExpectations(t)
}

func Test_redisTokenRepository_setRefreshToken_Happy(t *testing.T) {
	id := uint(1)
	tokenId := "some-uuid"
	var d time.Duration = 1000000000
	key := fmt.Sprintf("%d:%s", id, tokenId)

	r := &redisMock{}
	r.
		On("Set", key, 0, d).
		Return(redis.NewStatusCmd())

	repository := NewRepository(r)
	err := repository.setRefreshToken(id, tokenId, d)
	require.NoError(t, err)

	r.AssertExpectations(t)
}

type redisMock struct{ mock.Mock }

func (r *redisMock) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	called := r.Called(key, value, expiration)
	statusCmd, ok := called.Get(0).(*redis.StatusCmd)
	if ok {
		return statusCmd
	} else {
		return nil
	}
}

func (r *redisMock) Del(key ...string) *redis.IntCmd {
	called := r.Called(key)
	return called.Get(0).(*redis.IntCmd)
}

func (r *redisMock) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	//TODO implement me
	panic("implement me")
}
