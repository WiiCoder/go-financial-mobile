package db

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// Bill 账单
type Bill struct {
	gorm.Model
	OrderId         int             // 订单号
	Sequence        int             // 期数次序
	RepaymentTime   time.Time       // 计划还款时间
	Principal       decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 计划还款本金
	DaysPastDue     int             // 逾期天数
	InterestPastDue decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 逾期利息
	SalePrincipal   decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 优惠本金
	SaleInterest    decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 优惠利息
	PaidPrincipal   decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 已还本金
	PaidInterest    decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 已还利息
	Status          int             // 还款状态;1 - 待还款、2 - 支付中  、3 -  逾期、4 - 结清
	SettleTime      time.Time       // 结清时间
}

// BillInterest 账单利息记录
type BillInterest struct {
	gorm.Model
	BillId    int             // 账单ID
	Principal decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 计算本金
	Interest  decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 利息
	Rate      decimal.Decimal `gorm:"type:decimal(12,4);not null;default:0"` // 利率
	Days      int             // 计算逾期天数
}

// Contact 联系人
type Contact struct {
	gorm.Model
	Name        string // 姓名
	PhoneNumber string // 手机号
	Relation    int    // 关系： 1 - 父亲、2 - 母亲、3 - 朋友
}
