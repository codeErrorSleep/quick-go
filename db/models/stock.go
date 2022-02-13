package models

import "quick-go/db"

type Stock struct {
	ID         int    `json:"id"`
	AppID      string `json:"app_id"`
	SpuID      string `json:"spu_id"`
	SkuID      string `json:"sku_id"`
	SpuType    string `json:"spu_type"`
	ResourceID string `json:"resource_id"`
	SellNum    int    `json:"sell_num"`
	LeftNum    int    `json:"left_num"`
	Sales      int    `json:"sales"`
	IsDeleted  int    `json:"is_deleted"`
	// SettingStock int    `json:"setting_stock"`
	// CreatedAt string `json:"created_at"`
	// UpdatedAt string `json:"updated_at"`
}

func (stock *Stock) TableName(appId string) string {
	num := appId[len(appId)-1 : len(appId)]
	return "t_stock_" + num
}

func (stock *Stock) GetStockDetail(appId string, spuId string) (stockList []Stock, err error) {
	query := db.LocalMysql.
		Table(stock.TableName(appId)).
		Where("app_id = ?", appId).
		Where("spu_id = ?", spuId).
		Where("is_deleted = ?", 0).
		Select("spu_id,sku_id,left_num")

	err = query.Find(&stockList).Error

	if err != nil {
		return stockList, err
	}
	return

}
