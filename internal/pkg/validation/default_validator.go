package validation

type defaultValidation struct {
}

func (m defaultValidation) Validate(obj interface{}) error {
	return nil
}
