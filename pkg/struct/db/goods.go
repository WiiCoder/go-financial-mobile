package db

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Goods 商品
type Goods struct {
	gorm.Model
	DistributorId int    // 渠道ID
	Name          string // 商品名称
	PublishStatus int    // 上架状态，1 - 下架、2 - 上架
}

// GoodsSku 商品规格
type GoodsSku struct {
	gorm.Model
	GoodsId    int             // 商品ID
	Name       string          // 规格名称
	DiskSize   int             // 容量大小
	GoodsColor string          // 商品颜色
	Price      decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 商品单价
}
