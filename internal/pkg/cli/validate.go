package cli

import (
	"errors"
	"fmt"
	"regexp"
)
var (
	errValueEmpty        = errors.New("value must not be empty")
	errValueTooLong      = errors.New("value must not exceed 255 characters")
	errValueNotAString   = errors.New("value must be a string")
	errValueBadFormat = errors.New("value must be a valid deployment name")
)


func validateDeploymentName(val interface{}) error {
	if err := basicNameValidation(val); err != nil {
		return fmt.Errorf("deployment name %v is invalid: %w", val, err)
	}

	// TODO: check regex
	valid, err := regexp.MatchString(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*`, val.(string))
	
	if err != nil  {
		return errValueBadFormat
	}
	
	if !valid {
		return errValueBadFormat
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
