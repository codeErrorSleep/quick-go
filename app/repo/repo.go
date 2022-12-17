package repo

import (
	"context"
	"quick-go/app/models"
)

//go:generate mockgen -source repo.go -destination=../../test/mocks/mock_repo.go -package=mocks

// spu
type ISpuRepo interface {
	GetSpuDetail(ctx context.Context, appID string, spuID string) (spuDetail models.Spu, err error)
}

// stock
type IStockRepo interface {
	GetStockDetail(ctx context.Context, appID string, spuID string) (stockList []models.Stock, err error)
}
