package validators

// Validator 定义验证器类型
type Validator func() error

// New 创建验证器
func New(vs ...Validator) error {
	if len(vs) == 0 {
		return nil
	}
	for _, v := range vs {
		if err := v(); err != nil {
			return err
		}
	}
	return nil
}
