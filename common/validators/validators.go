package validators

import (
	chinese "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	vzh "github.com/go-playground/validator/v10/translations/zh"
)

// New 创建验证器
func New() *validator.Validate {
	zh := chinese.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")
	v := validator.New()
	for s, f := range vs {
		_ = v.RegisterValidation(s, f)
	}
	_ = vzh.RegisterDefaultTranslations(v, trans)
	return v
}
