package validation

const (
	EMAIL = "email"
)

var (
	Email            = emailValidation{}
	DefaultValidator = defaultValidation{}
)

type Validation interface {
	Validate()
}

func Default(validateType string) Validation {
	switch validateType {
	case EMAIL:
		return Email
	}
	return DefaultValidator
}
