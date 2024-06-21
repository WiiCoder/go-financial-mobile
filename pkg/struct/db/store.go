package db

import "gorm.io/gorm"

// Store 门店
type Store struct {
	gorm.Model
	Code            string // 门店代码
	Name            string // 门店名称
	DocumentAddress string // 法律文书地址
	Court           string // 所属法院
	Head            string // 负责人姓名
	HeadMobile      string // 负责人手机号
	HeadQq          string // 负责人QQ
	HeadWechat      string // 负责人微信
	DistributorId   int    // 渠道ID
	PartnerId       int    // 股东ID
}
