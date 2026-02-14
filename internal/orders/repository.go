package orders

import (
	"context"
	"sort"
	"sync"
	"time"
)

type Repository interface {
	List(ctx context.Context) ([]Order, error)
}

type InMemoryRepository struct {
	mu     sync.RWMutex
	orders []Order
}

func NewInMemoryRepository() *InMemoryRepository {
	now := time.Now().UTC()

	return &InMemoryRepository{
		orders: []Order{
			{ID: "ord_1001", Customer: "alice", AmountCents: 2599, CreatedAt: now.Add(-2 * time.Hour)},
			{ID: "ord_1002", Customer: "bob", AmountCents: 1099, CreatedAt: now.Add(-1 * time.Hour)},
		},
	}
}

func (r *InMemoryRepository) List(_ context.Context) ([]Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]Order, len(r.orders))
	copy(result, r.orders)

	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt.After(result[j].CreatedAt)
	})

	return result, nil
}
