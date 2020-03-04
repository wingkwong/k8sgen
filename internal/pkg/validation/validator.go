package validation

const (
	EMAIL          = "email"
	DEPLOYMENTNAME = "deployment_name"
)

var (
	DefaultValidator = defaultValidation{}
	Email            = emailValidation{}
	DeploymentName   = deploymentNameValidation{}
)

type Validation interface {
	Validate(obj interface{}) error
}

func Default(validateType string) Validation {
	switch validateType {
	case EMAIL:
		return Email
	case DEPLOYMENTNAME:
		return DeploymentName
	}
	return DefaultValidator
}
