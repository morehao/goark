package objuser

type PersonBaseInfo struct {
	// AvatarUrl 头像URL
	AvatarUrl string `json:"avatarUrl" form:"avatarUrl"`
	// BirthDate 出生日期
	BirthDate int64 `json:"birthDate" form:"birthDate"`
	// Email 邮箱
	Email string `json:"email" form:"email"`
	// Gender 性别: male-男 female-女 unknown-未知
	Gender string `json:"gender" form:"gender"`
	// Mobile 手机号
	Mobile string `json:"mobile" form:"mobile"`
	// PasswordHash 密码哈希(不存储盐值,盐值在应用层生成)
	PasswordHash string `json:"passwordHash" form:"passwordHash"`
	// RealName 真实姓名
	RealName string `json:"realName" form:"realName"`
	// Remark 备注
	Remark string `json:"remark" form:"remark"`
	// Wechat 微信号
	Wechat string `json:"wechat" form:"wechat"`
}
