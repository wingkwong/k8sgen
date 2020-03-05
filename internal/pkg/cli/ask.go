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
	inputNamespacePrompt                = ""

	// select
	inputOutputFormatPrompt            = "Please select an output format:"
	inputSecretCmdNamePrompt           = "Please select the type of secret:"
	inputRequireObjectMetaPrompt       = "Do you want to input Object Meta?"
	inputRequireDeploymentSpecPrompt   = "Do you want to input Deployment Spec?"
	inputRequireDeploymentStatusPrompt = "Do you want to input Deployment Status?"
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
	"Kind":                    {"KindName", "string", inputKindNamePrompt, "", "Select kind name", nil, kindNames, "AskSelect"},
	"DeploymentName":          {"DeploymentName", "string", inputDeploymentNamePrompt, "", "Prompt for deployment name", validateDeploymentName, nil, "AskGet"},
	"Image":                   {"ImageName", "string", inputImageNamePrompt, "", "Prompt for image name", nil, nil, "AskGet"},
	"OutputPath":              {"outputPath", "string", inputOutputPathPrompt, "", "Prompt for output path", nil, nil, "AskGet"},
	"SecretName":              {"secretName", "string", inputSecretNamePrompt, "", "Prompt for secret name", nil, nil, "AskGet"},
	"DockerServerName":        {"dockerServer", "string", inputDockerServerNamePrompt, "", "Prompt for docker server name", nil, nil, "AskGet"},
	"DockerUserName":          {"dockerUserName", "string", inputDockerUserNamePrompt, "", "Prompt for docker user name", nil, nil, "AskGet"},
	"DockerPassword":          {"dockerUserPassword", "string", inputDockerUserPasswordPrompt, "", "Prompt for docker password", nil, nil, "AskGetSecret"},
	"DockerEmail":             {"dockerEmail", "string", inputDockerEmailPrompt, "", "Prompt for docker email", nil, nil, "AskGet"}, // TODO: email validation
	"CertPath":                {"certPath", "string", inputDockerUserNamePrompt, "", "Prompt for cert path", nil, nil, "AskGet"},
	"KeyPath":                 {"keyPath", "string", inputDockerUserNamePrompt, "", "Prompt for key path", nil, nil, "AskGet"},
	"FromEnvFile":             {"fromEnvFile", "string", inputFromEnvFilePrompt, "", "Prompt for env", nil, nil, "AskGet"},
	"Namespace":               {"namespace", "string", inputNamespacePrompt, "", "Prompt for namespace", nil, nil, "AskGet"},
	"RequireObjectMeta":       {"requireObjectMeta", "bool", inputRequireObjectMetaPrompt, "", "Prompt for requireObjectMeta", nil, yesOrNo, "AskSelect"},
	"RequireDeploymentSpec":   {"requireDeploymentSpec", "bool", inputRequireDeploymentSpecPrompt, "", "Prompt for requireDeploymentSpec", nil, yesOrNo, "AskSelect"},
	"RequireDeploymentStatus": {"requireDeploymentStatus", "bool", inputRequireDeploymentStatusPrompt, "", "Prompt for requireDeploymentStatus", nil, yesOrNo, "AskSelect"},
	"OutputFormat":            {"outputFormat", "string", inputOutputFormatPrompt, "", "Prompt for output format", nil, outputFormats, "AskSelect"},
	"SecretCmdName":           {"secretCmdName", "string", inputSecretCmdNamePrompt, "", "Prompt for secret cmd name", nil, secretNames, "AskSelect"},
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
	val, err := o.prompt.SelectOne(inputOutputFormatPrompt, promptHelpMessage, opts)
	if err != nil {
		return fmt.Errorf("%s: %w", promptErrorMessage, err)
	}

	if err := setField(o, name, val, qType); err != nil {
		return fmt.Errorf("%s: %w", promptErrorMessage, err)
	}

	return nil
}

func (o *askOpts) AskGetSecret(name string, qType string, promptMessage string, promptHelpMessage string, promptErrorMessage string) error {
	val, err := o.prompt.GetSecret(inputOutputFormatPrompt, promptHelpMessage)
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
		o.AskGet(q.name, q.qType, q.promptMessage, q.promptHelpMessage, q.promptErrorMessage, q.validation)
	} else if q.funcName == "AskSelect" {
		o.AskSelect(q.name, q.qType, q.promptMessage, q.promptHelpMessage, q.promptErrorMessage, q.opts)
	} else if q.funcName == "AskGetSecret" {
		o.AskGetSecret(q.name, q.qType, q.promptMessage, q.promptHelpMessage, q.promptErrorMessage)
	} else {
		return fmt.Errorf("Unexpected q.funcName. Available options: AskGet, AskSelect, AskGetSecret")
	}

	return nil
}

func (o *askOpts) AskFromFilePath() error {
	if err := o.AskFromFileIteration(); err != nil {
		return err
	}

	for i := 0; i < o.noOfFromFileIteration; i++ {
		fromFile, err := o.prompt.Get(inputFromFilePrompt, "", nil /*no validation*/)
		if err != nil {
			return fmt.Errorf("Prompt for from-file: %w", err)
		}
		o.fromFile = append(o.fromFile, fromFile)
	}

	return nil
}

func (o *askOpts) AskFromLiteral() error {
	if err := o.AskFromLiteralIteration(); err != nil {
		return err
	}

	for i := 0; i < o.noOfFromLiteralIteration; i++ {
		fromLiteral, err := o.prompt.Get(inputFromLiteralPrompt, "", nil /*no validation*/)
		if err != nil {
			return fmt.Errorf("Prompt for from-literal: %w", err)
		}
		o.fromLiteral = append(o.fromLiteral, fromLiteral)
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

// TODO: refak
func (o *askOpts) AskFromFileIteration() error {
	// TODO: int vaildation
	noOfIterationStr, err := o.prompt.Get(inputNoOfFromFileIterationPrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for from-file iteration: %w", err)
	}

	noOfFromFileIteration, err := strconv.Atoi(noOfIterationStr)
	if err != nil {
		return fmt.Errorf("Prompt for from-file iteration: %w", err)
	}
	o.noOfFromFileIteration = noOfFromFileIteration

	return nil
}

// TODO: refak
func (o *askOpts) AskFromLiteralIteration() error {
	// TODO: int vaildation
	noOfFromLiteralIterationStr, err := o.prompt.Get(inputNoOfFromLiteralIterationPrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for from-literal iteration: %w", err)
	}

	noOfFromLiteralIteration, err := strconv.Atoi(noOfFromLiteralIterationStr)
	o.noOfFromLiteralIteration = noOfFromLiteralIteration

	return nil
}
