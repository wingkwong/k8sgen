package typecheck

import (
	"sync"

	"gopkg.in/go-playground/validator.v8"
)

type TypesValidator interface {
	Struct(interface{}) error
	String(interface{}) error
}

type defaultValidator struct {
	once sync.Once
	v    *validator.Validate
}
