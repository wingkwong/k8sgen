package cli

import (
	"errors"
	"fmt"
	"github.com/wingkwong/k8sgen/third_party/term/prompt"
	"reflect"
	"strconv"
)

const (
	// text input
	inputKindNamePrompt                 = "What kind of object you want to create?"
	inputDeploymentNamePrompt           = "What deployment you want to name?"
	inputImageNamePrompt                = "What image you want to name to run?"
	inputOutputPathPrompt               = "What directory you want to save?"
	inputSecretNamePrompt               = "What secret you want to name?"
	inputDockerServerNamePrompt         = "What docker server you want to name?"
	inputDockerUserNamePrompt           = "What is the username for Docker registry authentiation?"
	inputDockerUserPasswordPrompt       = "What is the password for Docker registry authentiation?"
	inputDockerEmailPrompt              = "What is the email address for Docker registry?"
	inputFromFilePrompt                 = "Input key names from file: (e.g path/to/bar):"
	inputFromLiteralPrompt              = "Input a key-value pair secret (e.g foo='bar'):"
	inputFromEnvFilePrompt              = "Where is the env file path?"
	inputNoOfFromFileIterationPrompt    = "How many from-file iterations for your input?"
	inputNoOfFromLiteralIterationPrompt = "How many from-literal iterations for your input?"

	// select
	inputOutputFormatPrompt  = "Please select an output format:"
	inputSecretCmdNamePrompt = "Please select the type of secret:"
)

type Question struct {
	name               string
	qType              string
	promptMessage      string
	promptHelpMessage  string
	promptErrorMessage string
	validation         prompt.ValidatorFunc
	opts               []string
	funcName           string
}

var questions = map[string]Question{
	"Kind":                     {"KindName", "string", inputKindNamePrompt, "", "Select kind name", nil /*no validation*/, kindNames, "AskSelect"},
	"DeploymentName":           {"DeploymentName", "string", inputDeploymentNamePrompt, "", "Prompt for deployment name", validateDeploymentName, nil, "AskGet"},
	"Image":                    {"ImageName", "string", inputImageNamePrompt, "", "Prompt for image name", nil /*no validation*/, nil, "AskGet"},
	"OutputPath":               {"OutputPath", "string", inputOutputPathPrompt, "", "Prompt for output path", nil /*no validation*/, nil, "AskGet"},
	"SecretName":               {"SecretName", "string", inputSecretNamePrompt, "", "Prompt for secret name", nil /*no validation*/, nil, "AskGet"},
	"DockerServerName":         {"DockerServer", "string", inputDockerServerNamePrompt, "", "Prompt for docker server name", nil /*no validation*/, nil, "AskGet"},
	"DockerUserName":           {"DockerUserName", "string", inputDockerUserNamePrompt, "", "Prompt for docker user name", nil /*no validation*/, nil, "AskGet"},
	"DockerPassword":           {"DockerUserPassword", "string", inputDockerUserPasswordPrompt, "", "Prompt for docker password", nil /*no validation*/, nil, "AskGetSecret"},
	"DockerEmail":              {"DockerEmail", "string", inputDockerEmailPrompt, "", "Prompt for docker email", nil /*no validation*/, nil, "AskGet"}, // TODO: email validation
	"CertPath":                 {"CertPath", "string", inputDockerUserNamePrompt, "", "Prompt for cert path", nil /*no validation*/, nil, "AskGet"},
	"KeyPath":                  {"KeyPath", "string", inputDockerUserNamePrompt, "", "Prompt for key path", nil /*no validation*/, nil, "AskGet"},
	"FromEnvFile":              {"FromEnvFile", "string", inputFromEnvFilePrompt, "", "Prompt for env", nil /*no validation*/, nil, "AskGet"},
	"OutputFormat":             {"OutputFormat", "string", inputOutputFormatPrompt, "", "Prompt for output format", nil /*no validation*/, outputFormats, "AskSelect"},
	"SecretCmdName":            {"SecretCmdName", "string", inputSecretCmdNamePrompt, "", "Prompt for secret cmd name", nil /*no validation*/, secretNames, "AskSelect"},
	"NoOfFromFileIteration":    {"NoOfFromFileIteration", "int", inputNoOfFromFileIterationPrompt, "", "Prompt for from-file iteration", nil /*no validation*/, nil, "AskGet"},
	"NoOfFromLiteralIteration": {"NoOfFromLiteralIteration", "int", inputNoOfFromLiteralIterationPrompt, "", "Prompt for from-literal iteration", nil /*no validation*/, nil, "AskGet"},
}

var yesOrNo = []string{"Yes", "No"}

func newAskOpts(vars askVars) (*askOpts, error) {
	return &askOpts{
		askVars: vars,
	}, nil
}

func setField(v interface{}, name string, value string, fvType string) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return errors.New("v must be pointer to struct")
	}
	rv = rv.Elem()
	fv := rv.FieldByName(name)
	if !fv.IsValid() {
		return fmt.Errorf("not a field name: %s", name)
	}

	if !fv.CanSet() {
		return fmt.Errorf("cannot set field %s", name)
	}

	if fvType == "string" {
		if fv.Kind() != reflect.String {
			return fmt.Errorf("%s is not a string field", name)
		}
		fv.SetString(value)
	} else if fvType == "int" {
		if fv.Kind() != reflect.Int {
			return fmt.Errorf("%s is not an int field", name)
		}
		val, _ := strconv.ParseInt(value, 10, 64)
		fv.SetInt(val)
	} else if fvType == "bool" {
		if fv.Kind() != reflect.Bool {
			return fmt.Errorf("%s is not a bool field", name)
		}
		var boolVal bool
		if value == "Yes" {
			boolVal = true
		} else if value == "No" {
			boolVal = false
		} else {
			return fmt.Errorf("%s is not either Yes or No", value)
		}

		fv.SetBool(boolVal)
	} else {
		return fmt.Errorf("%s is not either string, int or bool", fvType)
	}

	return nil
}

func (o *askOpts) AskGet(name string, qType string, promptMessage string, promptHelpMessage string, promptErrorMessage string, validation prompt.ValidatorFunc) error {
	val, err := o.prompt.Get(promptMessage, promptHelpMessage, validation)
	if err != nil {
		return fmt.Errorf("%s: %w", promptErrorMessage, err)
	}

	if err := setField(o, name, val, qType); err != nil {
		return fmt.Errorf("%s: %w", promptErrorMessage, err)
	}

	return nil
}

func (o *askOpts) AskSelect(name string, qType string, promptMessage string, promptHelpMessage string, promptErrorMessage string, opts []string) error {
	val, err := o.prompt.SelectOne(promptMessage, promptHelpMessage, opts)

	if err != nil {
		return fmt.Errorf("%s: %w", promptErrorMessage, err)
	}

	if err := setField(o, name, val, qType); err != nil {
		return fmt.Errorf("%s: %w", promptErrorMessage, err)
	}

	return nil
}

func (o *askOpts) AskGetSecret(name string, qType string, promptMessage string, promptHelpMessage string, promptErrorMessage string) error {
	val, err := o.prompt.GetSecret(promptMessage, promptHelpMessage)
	if err != nil {
		return fmt.Errorf("%s: %w", promptErrorMessage, err)
	}

	if err := setField(o, name, val, qType); err != nil {
		return fmt.Errorf("%s: %w", promptErrorMessage, err)
	}

	return nil
}

func (o *askOpts) Ask(key string) error {
	q, exists := questions[key]

	if !exists {
		return fmt.Errorf("key %s not exists in questions", key)
	}

	if q.funcName == "AskGet" {
		if err := o.AskGet(q.name, q.qType, q.promptMessage, q.promptHelpMessage, q.promptErrorMessage, q.validation); err != nil {
			return err
		}
	} else if q.funcName == "AskSelect" {
		if err := o.AskSelect(q.name, q.qType, q.promptMessage, q.promptHelpMessage, q.promptErrorMessage, q.opts); err != nil {
			return err
		}
	} else if q.funcName == "AskGetSecret" {
		if err := o.AskGetSecret(q.name, q.qType, q.promptMessage, q.promptHelpMessage, q.promptErrorMessage); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Unexpected q.funcName. Available options: AskGet, AskSelect, AskGetSecret")
	}

	return nil
}

func (o *askOpts) AskFromFilePath() error {
	if err := o.Ask("NoOfFromFileIteration"); err != nil {
		return err
	}

	for i := 0; i < o.NoOfFromFileIteration; i++ {
		fromFile, err := o.prompt.Get(inputFromFilePrompt, "", nil /*no validation*/)
		if err != nil {
			return fmt.Errorf("Prompt for from-file: %w", err)
		}
		o.FromFile = append(o.FromFile, fromFile)
	}

	return nil
}

func (o *askOpts) AskFromLiteral() error {
	if err := o.Ask("NoOfFromLiteralIteration"); err != nil {
		return err
	}

	for i := 0; i < o.NoOfFromLiteralIteration; i++ {
		fromLiteral, err := o.prompt.Get(inputFromLiteralPrompt, "", nil /*no validation*/)
		if err != nil {
			return fmt.Errorf("Prompt for from-literal: %w", err)
		}
		o.FromLiteral = append(o.FromLiteral, fromLiteral)
	}

	return nil
}

func (o *askOpts) AskOutputInfo() error {

	if err := o.Ask("OutputFormat"); err != nil {
		return err
	}

	if err := o.Ask("OutputPath"); err != nil {
		return err
	}

	return nil
}
