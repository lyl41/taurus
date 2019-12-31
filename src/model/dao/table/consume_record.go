package table

import (
	"time"
)

type ConsumeRecord struct {
	ID          int64     `sql:"primary_key;column:id" json:"id,omitempty"`         //
	UserID      int64     `sql:"column:user_id" json:"user_id,omitempty"`           //用户id
	Phone       string    `sql:"column:phone" json:"phone,omitempty"`               //手机号
	Amount      int64     `sql:"column:amount" json:"amount,omitempty"`             //消费金额
	MchName     string    `sql:"column:mch_name" json:"mch_name,omitempty"`         //商户名称
	CreatedUser string    `sql:"column:created_user" json:"created_user,omitempty"` //创建人
	CreatedAt   time.Time `sql:"column:created_at" json:"created_at,omitempty"`     //
	UpdatedAt   time.Time `sql:"column:updated_at" json:"updated_at,omitempty"`     //

}

func (ConsumeRecord) TableName() string {
	return "consume_record"
}
