package models

// 商品sku表
type Sku struct {
	Id                uint64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	SkuId             string `gorm:"column:sku_id" db:"sku_id" json:"sku_id" form:"sku_id"`                                                     //SKU_ID
	SkuBusinessId     string `gorm:"column:sku_business_id" db:"sku_business_id" json:"sku_business_id" form:"sku_business_id"`                 //资源侧的sku id
	AppId             string `gorm:"column:app_id" db:"app_id" json:"app_id" form:"app_id"`                                                     //店铺id
	SpuId             string `gorm:"column:spu_id" db:"spu_id" json:"spu_id" form:"spu_id"`                                                     //商品id
	SkuImg            string `gorm:"column:sku_img" db:"sku_img" json:"sku_img" form:"sku_img"`                                                 //sku默认图片
	SkuName           string `gorm:"column:sku_name" db:"sku_name" json:"sku_name" form:"sku_name"`                                             //sku名字
	SkuDesc           string `gorm:"column:sku_desc" db:"sku_desc" json:"sku_desc" form:"sku_desc"`                                             //规格描述
	SkuPrice          int    `gorm:"column:sku_price" db:"sku_price" json:"sku_price" form:"sku_price"`                                         //价格
	SkuLinePrice      int    `gorm:"column:sku_line_price" db:"sku_line_price" json:"sku_line_price" form:"sku_line_price"`                     //划线价
	SkuMinPurchase    int    `gorm:"column:sku_min_purchase" db:"sku_min_purchase" json:"sku_min_purchase" form:"sku_min_purchase"`             //起购数量(最小购买数量,0为不限制)默认0
	LimitPurchase     int    `gorm:"column:limit_purchase" db:"limit_purchase" json:"limit_purchase" form:"limit_purchase"`                     //限购数量
	LimitPurchaseType int8   `gorm:"column:limit_purchase_type" db:"limit_purchase_type" json:"limit_purchase_type" form:"limit_purchase_type"` //限购类型：0-无限制 1-每个用户购买次数 2-有效期内每个用户购买次数
	PeriodValue       string `gorm:"column:period_value" db:"period_value" json:"period_value" form:"period_value"`                             //有效期值
	PeriodType        int8   `gorm:"column:period_type" db:"period_type" json:"period_type" form:"period_type"`                                 //有效期类型：0-长期有效；1-具体时间前有效；2-有效期时间范围（秒）
	IsFree            int8   `gorm:"column:is_free" db:"is_free" json:"is_free" form:"is_free"`                                                 //是否免费（0-收费 1-免费）
	ValueCount        int8   `gorm:"column:value_count" db:"value_count" json:"value_count" form:"value_count"`                                 //1单规格 2多规格
	OrderStockLimit   int8   `gorm:"column:order_stock_limit" db:"order_stock_limit" json:"order_stock_limit" form:"order_stock_limit"`         //下单是否受库存限制：0 = 否，1 = 是
	HasDistribute     int8   `gorm:"column:has_distribute" db:"has_distribute" json:"has_distribute" form:"has_distribute"`                     //是否参与推广员推广（只针对sku的推广）0-否 1-是
	State             int8   `gorm:"column:state" db:"state" json:"state" form:"state"`                                                         //1上架 0下架
	IsDefault         int8   `gorm:"column:is_default" db:"is_default" json:"is_default" form:"is_default"`                                     //是否默认 0-否1-是
	IsDeleted         int8   `gorm:"column:is_deleted" db:"is_deleted" json:"is_deleted" form:"is_deleted"`                                     //0正常 1已删除
	CreatedAt         int64  `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`                                     //创建时间
	UpdatedAt         int64  `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`                                     //更新时间，有修改自动更新
	SkuSpecCode       string `gorm:"column:sku_spec_code" db:"sku_spec_code" json:"sku_spec_code" form:"sku_spec_code"`                         //商品sku规格编码
	SkuCostPrice      int    `gorm:"column:sku_cost_price" db:"sku_cost_price" json:"sku_cost_price" form:"sku_cost_price"`                     //成本价单位分
	SkuWeight         string `gorm:"column:sku_weight" db:"sku_weight" json:"sku_weight" form:"sku_weight"`                                     //商品sku重量kg为单位
	SkuVolume         string `gorm:"column:sku_volume" db:"sku_volume" json:"sku_volume" form:"sku_volume"`                                     //商品sku体积立方米为单位
}
