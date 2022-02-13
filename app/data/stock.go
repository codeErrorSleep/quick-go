package data

type StockDBEntry struct {
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
