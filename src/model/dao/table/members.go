package table

import (
	"time"
)

type Members struct {
	ID          int64     `sql:"primary_key;column:id" json:"id,omitempty"`         //
	Name        string    `sql:"column:name" json:"name,omitempty"`                 //会员姓名
	Phone       string    `sql:"column:phone" json:"phone,omitempty"`               //手机号
	Amount      int64     `sql:"column:amount" json:"amount,omitempty"`             //会员余额
	CertID      string    `sql:"column:cert_id" json:"cert_id,omitempty"`           //身份证号
	MchName     string    `sql:"column:mch_name" json:"mch_name,omitempty"`         //商家名称
	Status      int64     `sql:"column:status" json:"status,omitempty"`             //状态，1表示正常，2表示删除
	CreatedUser string    `sql:"column:created_user" json:"created_user,omitempty"` //创建人
	CreatedAt   time.Time `sql:"column:created_at" json:"created_at,omitempty"`     //
	UpdatedAt   time.Time `sql:"column:updated_at" json:"updated_at,omitempty"`     //

}

func (Members) TableName() string {
	return "members"
}
