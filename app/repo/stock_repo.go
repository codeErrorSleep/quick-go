package repo

import (
	"context"
	"quick-go/app/models"
	"quick-go/global"
)

type mysqlStockRepository struct {
	data *global.Data
}

func NewMysqlStockRepository(data *global.Data) IStockRepo {
	return &mysqlStockRepository{data}
}

func (m *mysqlStockRepository) GetStockDetail(ctx context.Context, appID string, spuID string) (stockList []models.Stock, err error) {
	query := m.data.LocalMysql.
		Table((&models.Stock{}).TableName(appID)).
		Where("app_id = ?", appID).
		Where("spu_id = ?", spuID).
		Where("is_deleted = ?", 0).
		Select("spu_id,sku_id,left_num,sell_num")

	err = query.Find(&stockList).Error

	if err != nil {
		return stockList, err
	}
	return

}
