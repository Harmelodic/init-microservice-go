package main

import (
	"github.com/Harmelodic/init-microservice-go/internal/account"
	"github.com/Harmelodic/init-microservice-go/internal/caching"
	"github.com/google/uuid"
)

func main() {
	cacheManager := caching.NewCacheManager[account.Account](caching.NewFifoCache[account.Account](100))

	cacheManager.GetOrAddFromFunc("thing", func(key string) account.Account {
		return account.Account{
			Id:    uuid.UUID{},
			Alias: key,
		}
	})
}
