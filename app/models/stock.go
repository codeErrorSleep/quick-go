package models

import "quick-go/db"

type Stock struct {
	ID         int64  `json:"id"`
	AppID      string `json:"app_id"`
	SpuID      string `json:"spu_id"`
	SkuID      string `json:"sku_id"`
	SpuType    string `json:"spu_type"`
	ResourceID string `json:"resource_id"`
	SellNum    int64  `json:"sell_num"`
	LeftNum    int64  `json:"left_num"`
	Sales      int64  `json:"sales"`
	IsDeleted  int64  `json:"is_deleted"`
	// SettingStock int    `json:"setting_stock"`
	// CreatedAt string `json:"created_at"`
	// UpdatedAt string `json:"updated_at"`
}

func (stock *Stock) TableName(appID string) string {
	num := appID[len(appID)-1 : len(appID)]
	return "t_stock_" + num
}

func (stock *Stock) GetStockDetail(appID string, spuID string) (stockList []Stock, err error) {
	query := db.LocalMysql.
		Table(stock.TableName(appID)).
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
