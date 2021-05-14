package rdd

// LoginValue redis存储的登录参数值，键为用户的uid+"_login"
type LoginValue struct {
	Password string   `json:"password"` // 用户的密码
	Mids     []string `json:"mids"`     // 已经登录的设备号
}
