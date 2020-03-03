package cli

import (
	"github.com/wingkwong/k8sgen/third_party/term/prompt"
)

type prompter interface {
	Get(message, help string, validator prompt.ValidatorFunc, opts ...prompt.GetOption) (string, error)
	GetSecret(message, help string) (string, error)
	SelectOne(message, help string, options []string) (string, error)
	Confirm(message, help string, options ...prompt.ConfirmOption) (bool, error)
}
