package table

import (
	"time"
)

type Merchant struct {
	ID            int64     `sql:"primary_key;column:id" json:"id,omitempty"`             //
	MchName       string    `sql:"column:mch_name" json:"mch_name,omitempty"`             //商家名称
	Amount        int64     `sql:"column:amount" json:"amount,omitempty"`                 //商家余额，钱的单位都是分
	Intro         string    `sql:"column:intro" json:"intro,omitempty"`                   //备注信息
	WarningAmount int64     `sql:"column:warning_amount" json:"warning_amount,omitempty"` //预警金额
	CreatedAt     time.Time `sql:"column:created_at" json:"created_at,omitempty"`         //
	UpdatedAt     time.Time `sql:"column:updated_at" json:"updated_at,omitempty"`         //

}

func (Merchant) TableName() string {
	return "merchant"
}
