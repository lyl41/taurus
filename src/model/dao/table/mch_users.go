package table

import (
	"time"
)

type MchUsers struct {
	ID        int64     `sql:"primary_key;column:id" json:"id,omitempty"`     //
	Username  string    `sql:"column:username" json:"username,omitempty"`     //用户名
	Password  string    `sql:"column:password" json:"password,omitempty"`     //密码，MD5形式
	Role      int64     `sql:"column:role" json:"role,omitempty"`             //1：是管理员，2：普通用户
	CreatedAt time.Time `sql:"column:created_at" json:"created_at,omitempty"` //
	UpdatedAt time.Time `sql:"column:updated_at" json:"updated_at,omitempty"` //

}

func (MchUsers) TableName() string {
	return "mch_users"
}
