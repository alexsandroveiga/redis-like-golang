package storage

import (
	"context"
	"sync"
	"time"

	"github.com/alexsandroveiga/redis-like-golang/internal/domain/entity"
	"github.com/alexsandroveiga/redis-like-golang/internal/domain/repository"
)

type Store struct {
	data        map[string]*entity.Item
	mu          sync.RWMutex
	stopCleanup chan struct{}
}

func NewStore() repository.KeyValueRepository {
	return &Store{
		data:        make(map[string]*entity.Item),
		stopCleanup: make(chan struct{}),
	}
}

func (s *Store) Set(ctx context.Context, key string, value string) {
	if ctx.Err() != nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = &entity.Item{Value: value, ExpiresAt: nil}
}

func (s *Store) Get(ctx context.Context, key string) (string, bool) {
	if ctx.Err() != nil {
		return "", false
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, exists := s.data[key]
	if !exists {
		return "", false
	}
	if item.IsExpired(time.Now().Unix()) {
		return "", false
	}
	return item.Value, true
}

func (s *Store) Del(ctx context.Context, key string) int {
	if ctx.Err() != nil {
		return 0
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.data[key]; exists {
		delete(s.data, key)
		return 1
	}
	return 0
}

func (s *Store) Expire(ctx context.Context, key string, durationInSeconds int) bool {
	if ctx.Err() != nil {
		return false
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	item, exists := s.data[key]
	if !exists {
		return false
	}
	expiresAt := time.Now().Unix() + int64(durationInSeconds)
	item.ExpiresAt = &expiresAt
	return true
}

func (s *Store) TTL(ctx context.Context, key string) int64 {
	if ctx.Err() != nil {
		return -1
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, exists := s.data[key]
	if !exists {
		return -1
	}
	if item.ExpiresAt == nil {
		return -1
	}
	now := time.Now().Unix()
	remaining := *item.ExpiresAt - now
	if remaining <= 0 {
		return -1
	}
	return remaining
}

func (s *Store) Persist(ctx context.Context, key string) bool {
	panic("unimplemented")
}

func (s *Store) Keys(ctx context.Context, pattern string) []string {
	panic("unimplemented")
}

func (s *Store) Exists(ctx context.Context, key string) bool {
	panic("unimplemented")
}
func (s *Store) Size(ctx context.Context) int {
	panic("unimplemented")
}

func (s *Store) StartCleanup(intervalInMs int64) {
	panic("unimplemented")
}

func (s *Store) StopCleanup() {
	panic("unimplemented")
}
