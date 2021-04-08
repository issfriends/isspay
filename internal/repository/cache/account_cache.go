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
	err := cache.Client.Set(ctx, tokenKey+":"+msgID, token.AccessToken, time.Duration(token.ExpiredAt)).Err()
	if err != nil {
		return err
	}

	return nil
}
