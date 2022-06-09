package repo

import (
	"context"
	"quick-go/app/models"

	"gorm.io/gorm"
)

type mysqlSpuRepository struct {
	DB *gorm.DB
}

func NewMysqlSpuRepository(DB *gorm.DB) ISpuRepo {
	return &mysqlSpuRepository{DB}
}

func (m *mysqlSpuRepository) GetSpuDetail(ctx context.Context, appID string, spuID string) (spuDetail models.Spu, err error) {
	query := m.DB.WithContext(ctx).
		Table((&models.Spu{}).TableName(appID)).
		Where("app_id = ?", appID).
		Where("spu_id = ?", spuID).
		Where("is_deleted = ?", 0)

	err = query.First(&spuDetail).Error

	if err != nil {
		return spuDetail, err
	}
	return
}
