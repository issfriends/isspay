package cache

import (
	"context"
	"time"

	"github.com/issfriends/isspay/internal/pkg/crypto"
	"github.com/vx416/gox/cache"
)

const (
	tokenKey = "token"
)

type AccountCache struct {
	*cache.RedisClient
}

func (cache AccountCache) CacheTokenWithMsgID(ctx context.Context, token *crypto.Token, msgID string) error {
	err := cache.Client.Set(ctx, tokenKey+":"+msgID, token.AccessToken, 7*24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (cache AccountCache) GetTokenByMsgID(ctx context.Context, msgID string) (*crypto.Token, error) {
	tokenStr, err := cache.Client.Get(ctx, tokenKey+":"+msgID).Result()
	if err != nil {
		return nil, err
	}
	expiredAt, err := cache.Client.TTL(ctx, tokenKey+":"+msgID).Result()
	if err != nil {
		return nil, err
	}

	return &crypto.Token{
		AccessToken: tokenStr,
		ExpiredAt:   int64(expiredAt),
	}, nil
}
