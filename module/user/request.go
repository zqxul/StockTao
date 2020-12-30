package user

// RegisterRequest ==> User register request
type RegisterRequest struct {
	Username string `json:"username" binding:"required"` // 用户名：2-15位
	Password string `json:"password" binding:"required"` // 密码：8-10位
	Email    string `json:"email"`                       // 邮箱
	NickName string `json:"nickName" binding:"required"` // 昵称
}

// LoginRequest ==> User login request
type LoginRequest struct {
	Username   string `json:"username" binding:"required"`   // 用户名：2-15位
	Password   string `json:"password" binding:"required"`   // 密码：8-10位
	RememberMe bool   `json:"rememberMe"`                    // 是否记住我: true-是|false-否(默认)
	VerifyCode string `json:"verifyCode" binding:"required"` // 验证码
}
