package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var vs = map[string]func(l validator.FieldLevel) bool{
	"password": password,
}

// 检查密码
func password(l validator.FieldLevel) bool {
	pass := l.Field().String()
	ok, err := regexp.MatchString(passwordRule, pass)
	if err != nil {
		return false
	}
	return ok
}
