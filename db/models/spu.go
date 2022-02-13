package models

import (
	"quick-go/db"
	"time"
)

type Spu struct {
	ID                  uint64    `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null" json:"-"`
	AppID               string    `gorm:"uniqueIndex:unx_app_resource_type;index:idx_app_spu;column:app_id;type:varchar(64);not null" json:"app_id"`    // 店铺id
	SpuID               string    `gorm:"index:idx_app_spu;column:spu_id;type:varchar(64);not null" json:"spu_id"`                                      // 商品id
	SpuType             string    `gorm:"column:spu_type;type:varchar(32);not null;default:''" json:"spu_type"`                                         // 商品类型(关联t_spu_type)
	ResourceID          string    `gorm:"uniqueIndex:unx_app_resource_type;column:resource_id;type:varchar(64);not null" json:"resource_id"`            // 资源id
	ResourceType        int       `gorm:"uniqueIndex:unx_app_resource_type;column:resource_type;type:int(4);not null" json:"resource_type"`             // 资源类型(21-实物商品 、1-图文、2-音频、3-视频、20-电子书、 6-专栏、 8-大专栏 、25-训练营、5-会员、23-超级会员、4-直播、 31-班课、29-线下课、11-活动管理、16-付费打卡、7-小社群、34-练习、41-有价优惠券)
	GoodsSn             string    `gorm:"column:goods_sn;type:varchar(64);not null;default:''" json:"goods_sn"`                                         // 商品编号（商家录入）
	GoodsCategoryID     string    `gorm:"column:goods_category_id;type:varchar(64);not null;default:''" json:"goods_category_id"`                       // 商品分类id
	GoodsName           string    `gorm:"column:goods_name;type:varchar(256);not null;default:''" json:"goods_name"`                                    // 商品名称
	GoodsImg            string    `gorm:"column:goods_img;type:varchar(1024);not null;default:''" json:"goods_img"`                                     // 商品封面图（默认封面图）
	CustomCover         string    `gorm:"column:custom_cover;type:varchar(1024);not null;default:''" json:"custom_cover"`                               // 主图视频自定义封面
	GoodsBriefText      string    `gorm:"column:goods_brief_text;type:mediumtext" json:"goods_brief_text"`                                              // 商品简介
	GoodsDetailText     string    `gorm:"column:goods_detail_text;type:mediumtext" json:"goods_detail_text"`                                            // 商品详情/买点
	SellType            int8      `gorm:"column:sell_type;type:tinyint(4);not null;default:1" json:"sell_type"`                                         // 付费类型：1 = 独立售卖，2 = 关联售卖
	PriceLow            int       `gorm:"column:price_low;type:int(11);not null;default:0" json:"price_low"`                                            // 商品最低价（取自sku中最低价值）
	PriceHigh           int       `gorm:"column:price_high;type:int(11);not null;default:0" json:"price_high"`                                          // 商品高价（取自最低价所在sku对应化线价）
	PriceLine           int       `gorm:"column:price_line;type:int(11);not null;default:0" json:"price_line"`                                          // 划线价(取值price_low所在sku划线价)
	VisitNum            int       `gorm:"column:visit_num;type:int(11);not null;default:0" json:"visit_num"`                                            // 访问量
	GoodsTag            string    `gorm:"column:goods_tag;type:varchar(258);not null;default:''" json:"goods_tag"`                                      // 商品标签
	GoodsTagIsShow      bool      `gorm:"column:goods_tag_is_show;type:tinyint(1);not null;default:1" json:"goods_tag_is_show"`                         // 商品标签是否展示;0:不展示;1:展示
	SaleStatus          bool      `gorm:"column:sale_status;type:tinyint(1);not null;default:1" json:"sale_status"`                                     // 上架状态： 0下架 1上架 2待上架
	IsTimingSale        bool      `gorm:"column:is_timing_sale;type:tinyint(1);not null;default:0" json:"is_timing_sale"`                               // 是否定时上架：1是 0否
	TimingSale          string    `gorm:"column:timing_sale;type:varchar(20);not null;default:''" json:"timing_sale"`                                   // 定时上架的时间
	SaleAt              string    `gorm:"column:sale_at;type:varchar(20);not null;default:''" json:"sale_at"`                                           // 上架的时间
	HasDistribute       bool      `gorm:"column:has_distribute;type:tinyint(1);not null;default:0" json:"has_distribute"`                               // 是否参与推广分销：0 = 否，1 = 是
	VideoURL            string    `gorm:"column:video_url;type:varchar(255);not null;default:''" json:"video_url"`                                      // 视频文件url
	VideoImgURL         string    `gorm:"column:video_img_url;type:varchar(255);not null;default:''" json:"video_img_url"`                              // 视频封面url
	Period              int       `gorm:"column:period;type:int(11);not null;default:-1" json:"period"`                                                 // 有效期（-1 永久，大于等于0是真实有效期）
	IsGoodsPackage      bool      `gorm:"column:is_goods_package;type:tinyint(1);not null;default:0" json:"is_goods_package"`                           // 是否带货 0 否 1 是
	IsDisplay           bool      `gorm:"column:is_display;type:tinyint(1);not null;default:0" json:"is_display"`                                       // 是否显示：0否(隐藏状态) 1是(显示状态)
	IsStopSell          bool      `gorm:"column:is_stop_sell;type:tinyint(1);not null;default:0" json:"is_stop_sell"`                                   // 是否停售：0否 1是
	IsForbid            int8      `gorm:"column:is_forbid;type:tinyint(4);not null;default:0" json:"is_forbid"`                                         // 商品是否被封禁：0 = 否，1 = 是
	IsIgnore            int8      `gorm:"column:is_ignore;type:tinyint(4);not null;default:0" json:"is_ignore"`                                         // 商品是否被忽略：0 = 否，1 = 是
	LimitPurchase       int       `gorm:"column:limit_purchase;type:int(11);not null;default:0" json:"limit_purchase"`                                  // 限购数量
	StockDeductMode     bool      `gorm:"column:stock_deduct_mode;type:tinyint(1);not null;default:0" json:"stock_deduct_mode"`                         // 扣库存方式：0付款减库存 1拍下减库存
	AppraiseNum         int       `gorm:"column:appraise_num;type:int(11);not null;default:0" json:"appraise_num"`                                      // 评价数
	ShowStock           bool      `gorm:"column:show_stock;type:tinyint(1);not null;default:1" json:"show_stock"`                                       // 是否展示库存 0-不展示 1-展示
	IsBest              bool      `gorm:"column:is_best;type:tinyint(1);not null;default:0" json:"is_best"`                                             // 是否精品 0否 1是
	IsHot               bool      `gorm:"column:is_hot;type:tinyint(1);not null;default:0" json:"is_hot"`                                               // 是否热销产品 0否 1是
	IsNew               bool      `gorm:"column:is_new;type:tinyint(1);not null;default:0" json:"is_new"`                                               // 是否新品   0否 1是
	IsRecom             bool      `gorm:"column:is_recom;type:tinyint(1);not null;default:0" json:"is_recom"`                                           // 是否推荐   0否 1是
	DistributionPattern int8      `gorm:"column:distribution_pattern;type:tinyint(4);not null;default:1" json:"distribution_pattern"`                   // 配送方式
	Freight             int       `gorm:"column:freight;type:int(11);not null;default:0" json:"freight"`                                                // 运费（单位：分）
	IsUniformFreight    int8      `gorm:"column:is_uniform_freight;type:tinyint(1);not null;default:1" json:"is_uniform_freight"`                       // 是否是统一运费 1:统一运费 2:运费模板
	FreightTemplateID   int       `gorm:"column:freight_template_id;type:int(11);not null;default:0" json:"freight_template_id"`                        // 运费模板ID
	ImgURLCompressed    string    `gorm:"column:img_url_compressed;type:varchar(255);not null;default:''" json:"img_url_compressed"`                    // 压缩后的列表配图url
	IsDeleted           int8      `gorm:"column:is_deleted;type:tinyint(4);not null;default:0" json:"is_deleted"`                                       // 0正常 1已删除
	IsPassword          bool      `gorm:"column:is_password;type:tinyint(1);not null;default:0" json:"is_password"`                                     // 是否加密（0-不加密 1--加密）
	IsFree              bool      `gorm:"column:is_free;type:tinyint(1);not null;default:0" json:"is_free"`                                             // 是否免费（0-收费 1-免费）
	CanSoldStart        time.Time `gorm:"column:can_sold_start;type:timestamp" json:"can_sold_start"`                                                   // 可售开始时间
	CanSoldEnd          time.Time `gorm:"column:can_sold_end;type:timestamp" json:"can_sold_end"`                                                       // 可售结束时间
	SellMode            bool      `gorm:"column:sell_mode;type:tinyint(1);not null;default:1" json:"sell_mode"`                                         // 售卖类型（1-自营 2-内容市场）
	IsPublic            bool      `gorm:"column:is_public;type:tinyint(1);not null;default:1" json:"is_public"`                                         // 是否公开售卖(0-不公开 1-公开)
	IsSingle            bool      `gorm:"column:is_single;type:tinyint(1);not null;default:1" json:"is_single"`                                         // 是否为单品（0-否  1-是）
	CreatedAt           time.Time `gorm:"index:idx_created_at;column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`   // 创建时间
	UpdatedAt           time.Time `gorm:"index:idx_updated_at;column:updated_at;type:timestamp;not null;default:0000-00-00 00:00:00" json:"updated_at"` // 更新时间
	WxGoodsCategoryID   string    `gorm:"column:wx_goods_category_id;type:varchar(64);not null;default:''" json:"wx_goods_category_id"`                 // 微信商品分类id
}

func (spu *Spu) TableName(appId string) string {
	num := appId[len(appId)-1 : len(appId)]
	return "t_spu_" + num
}

func (spu *Spu) GetSpuDetail(appId string, spuId string) (spuDetail Spu, err error) {
	query := db.LocalMysql.
		Table(spu.TableName(appId)).
		Where("app_id = ?", appId).
		Where("spu_id = ?", spuId).
		Where("is_deleted = ?", 0)

	err = query.First(&spuDetail).Error

	if err != nil {
		return spuDetail, err
	}
	return
}
