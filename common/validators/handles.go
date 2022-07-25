package validators

import "regexp"

// Password 检查密码
func Password(d string) Validator {
	return func() error {
		ok, err := regexp.MatchString(passwordRule, d)
		if err != nil {
			return err
		}
		if !ok {
			return passwordError
		}
		return nil
	}
}

// Email 邮箱
func Email(d string) Validator {
	return func() error {
		ok, err := regexp.MatchString(emailRule, d)
		if err != nil {
			return err
		}
		if !ok {
			return emailError
		}
		return nil
	}
}

func Empty(d interface{}) Validator {
	return func() error {
		switch d.(type) {
		case string:
			if d == "" {
				return emptyError
			}
		case int64, int, int32, uint, uint32, uint64:
			if d == 0 {
				return emptyError
			}
		}
		return nil
	}
}
