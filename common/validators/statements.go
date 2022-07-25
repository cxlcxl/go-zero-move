package validators

import "errors"

var (
	passwordError = errors.New("密码格式不正确")
	emailError    = errors.New("邮箱格式不正确")
	emptyError    = errors.New("必填参数不可为空")
)

const (
	passwordRule = "^[a-zA-z]+[a-zA-Z0-9@.&!#?,%$]{5,17}$"
	// 配置代码
	configKeyRule = `^[a-zA-Z]+[a-zA-Z0-9_]{0,49}$`
	// 邮箱正则
	emailRule = `^[\w\.]+@\w+\.[a-z]{2,3}(\.[a-z]{2,3})?$`
	// 手机号码正则
	mobileRule = `^[1]\d{10}$`
	// IDCard 身份证号码
	IDCard = `^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`
)
