package cli

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	dns1123LabelFmt       = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"
	dns1123LabelMaxLength = 63
)

var (
	errValueEmpty      = errors.New("value must not be empty")
	errValueTooLong    = errors.New("value must not exceed 255 characters")
	errValueNotAString = errors.New("value must be a string")
	errValueDNS1123    = errors.New("a DNS-1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character")
	dns1123LabelRegexp = regexp.MustCompile("^" + dns1123LabelFmt + "$")
)

func validateDeploymentName(val interface{}) error {
	if err := basicNameValidation(val); err != nil {
		return fmt.Errorf("deployment name %v is invalid: %w", val, err)
	}

	s, _ := val.(string)

	if len(s) > dns1123LabelMaxLength {
		return errors.New(fmt.Sprintf("must be no more than %d characters", dns1123LabelMaxLength))
	}
	if !dns1123LabelRegexp.MatchString(s) {
		return errValueDNS1123
	}

	return nil
}

func basicNameValidation(val interface{}) error {
	s, ok := val.(string)
	if !ok {
		return errValueNotAString
	}
	if s == "" {
		return errValueEmpty
	}
	if len(s) > 255 {
		return errValueTooLong
	}
	return nil
}
