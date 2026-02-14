package orders

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) List(ctx context.Context) ([]Order, error) {
	return s.repo.List(ctx)
}
