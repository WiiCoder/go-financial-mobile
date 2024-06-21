package db

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// File 文件
type File struct {
	gorm.Model
	Name       string // 文件名称
	OriginName string // 原文件名称
	Path       string // 文件路径
	Url        string // 文件链接
	Suffix     string // 文件后缀
	Type       int    // 文件类型，1 - 图片、2 - 文档、3 - 其他
}

// Order 订单
type Order struct {
	gorm.Model
	Code                    string          // 订单编号
	StoreId                 int             // 门店ID
	Mobile                  string          // 手机号
	FullName                string          // 购买人姓名
	IdNumber                string          // 身份证号码
	DocumentAddress         string          // 法律文书地址
	CompleteAddress         string          // 详细住址
	GoodsSkuId              int             // 商品规格ID
	Quantity                int             // 购买数量
	AmountTotal             decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 订单金额
	FirstPayAmount          decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 首付金额
	FirstPayType            int             // 付款方式，1 - 现金、2 - 微信、3 - 支付宝
	FirstPayTime            time.Time       // 首付时间
	InstallmentsNum         int             // 分期数
	InstallmentsStartMonth  time.Time       // 开始还款月份
	RepaymentDay            int             // 每月还款日
	InstallmentsAmount      decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` //  每期还款金额
	InstallmentsAmountTotal decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 总分期金额
	ProcessStatus           int             // 订单处理状态，1 - 待处理，2 - 已处理
	Status                  int             // 订单状态，1 - 正常、2 - 逾期、3 - 结清、4 - 退单
	Salesman                string          // 业务员
	ServiceFee              decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 服务费
	GoodsCost               decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 成本
	Profits                 decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 利润
	Commission              decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 提成
	AdvanceFunds            decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 垫资金额
}

// OrderContact 订单联系人
type OrderContact struct {
	gorm.Model
	OrderId   int // 订单ID
	ContactId int // 联系人ID
}

// OrderFile 订单文件
type OrderFile struct {
	gorm.Model
	OrderId int // 订单ID
	FileId  int // 文件ID
}
