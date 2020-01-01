package common

type Token struct {
	Username    string `json:"username"`
	Role        int    `json:"role"`
	MchId       int    `json:"mch_id"`
	ExpiredTime int64  `json:"expired_time"`
}
