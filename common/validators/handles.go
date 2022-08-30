package validators

import (
	"business/common/vars"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var vs = map[string]func(l validator.FieldLevel) bool{
	"password":      password,
	"campaign_type": campaignType,
	"product_type":  productType,
}

// 检查密码
func password(l validator.FieldLevel) bool {
	v := l.Field().String()
	ok, err := regexp.MatchString(passwordRule, v)
	if err != nil {
		return false
	}
	return ok
}

// 计划类型格式检测
func campaignType(l validator.FieldLevel) bool {
	v := l.Field().String()
	if _, ok := vars.CampaignType[v]; !ok {
		return false
	}
	return true
}

// 计划/创意/任务 名称格式检测
func productType(l validator.FieldLevel) bool {
	v := l.Field().String()
	if _, ok := vars.ProductType[v]; !ok {
		return false
	}
	return true
}
