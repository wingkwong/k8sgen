package cli

import (
	"errors"
	"fmt"
	"strconv"
    "reflect"
	"github.com/wingkwong/k8sgen/third_party/term/prompt"
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

type Question struct{
	name string
	promptMessage string
	promptHelpMessage string
	promptErrorMessage string
	validation prompt.ValidatorFunc
	opts []string
	funcName string
}

var questions = map[string]Question{
	"Kind": {"KindName", inputKindNamePrompt, "", "Select kind name", nil, kindNames, "AskSelect"},
	"DeploymentName": {"DeploymentName", inputDeploymentNamePrompt, "", "Prompt for deployment name", validateDeploymentName, nil, "AskGet"},
}



func getYesNoSelectOpts() []string {
	return []string{
		"Yes",
		"No",
	}
}

func getQuestionKey(key string) {
	
}

func newAskOpts(vars askVars) (*askOpts, error) {
	return &askOpts{
		askVars: vars,
	}, nil
}

func (o *askOpts) AskGet(name string, promptMessage string, promptHelpMessage string, promptErrorMessage string, validation prompt.ValidatorFunc ) error {
	val, err := o.prompt.Get(promptMessage, promptHelpMessage, validation)
	if err != nil {
		return fmt.Errorf("%s: %w", promptErrorMessage, err)
	}

	reflect.ValueOf(&o).Elem().Set(reflect.ValueOf(val))
	// (*o)[name] = val

	return nil
}

func (o *askOpts) AskSelect(name string, promptMessage string, promptHelpMessage string, promptErrorMessage string, opts []string) error {
	val, err := o.prompt.SelectOne(inputOutputFormatPrompt, promptHelpMessage, opts)
	if err != nil {
		return fmt.Errorf("%s: %w", promptErrorMessage, err)
	}
	
	// o.KindName = val
	

	return nil
}

func (o *askOpts) AskGetSecret(name string, promptMessage string, promptHelpMessage string, promptErrorMessage string) error {
	val, err := o.prompt.GetSecret(inputOutputFormatPrompt, promptHelpMessage)
	if err != nil {
		return fmt.Errorf("%s: %w", promptErrorMessage, err)
	}
	reflect.ValueOf(&o).Elem().FieldByName(name).SetString(val)

	return nil
}

func (o *askOpts) Ask(key string) error{
	q, exists := questions[key]

	if !exists { 
		return fmt.Errorf("key not exists in questions")
	}

	if q.funcName == "AskGet" {
		o.AskGet(q.name, q.promptMessage, q.promptHelpMessage, q.promptErrorMessage, q.validation)
	} else if q.funcName == "AskSelect" {
		o.AskSelect(q.name, q.promptMessage, q.promptHelpMessage, q.promptErrorMessage, q.opts)
	} else if q.funcName == "AskGetSecret" {
		o.AskGetSecret(q.name, q.promptMessage, q.promptHelpMessage, q.promptErrorMessage)
	} else {
		return fmt.Errorf("Unexpected q.funcName. Available options: AskGet, AskSelect, AskGetSecret")
	}

	return nil
}

func (o *askOpts) askKindName() error {
	if o.KindName != "" {
		return nil
	}

	names := getKindNames()

	if len(names) == 0 {
		return errors.New("No object is found")
	}

	selectedKindName, err := o.prompt.SelectOne(inputKindNamePrompt, "", names)
	if err != nil {
		return fmt.Errorf("Select kind name: %w", err)
	}
	o.KindName = selectedKindName
	return nil
}

func (o *askOpts) AskDeploymentName() error {
	deploymentName, err := o.prompt.Get(inputDeploymentNamePrompt, "", validateDeploymentName)

	if err != nil {
		return fmt.Errorf("Prompt for deployment name: %w", err)
	}

	o.deploymentName = deploymentName

	return nil
}

func (o *askOpts) AskImageName() error {
	imageName, err := o.prompt.Get(inputImageNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for image name: %w", err)
	}
	o.imageName = imageName

	return nil
}

func (o *askOpts) AskOutputFormat() error {
	formats := getOutputFormats()
	outputFormat, err := o.prompt.SelectOne(inputOutputFormatPrompt, "", formats)
	if err != nil {
		return fmt.Errorf("Prompt for output format: %w", err)
	}
	o.outputFormat = outputFormat

	return nil
}

func (o *askOpts) AskOutputPath() error {
	outputPath, err := o.prompt.Get(inputOutputPathPrompt, "", nil)
	if err != nil {
		return fmt.Errorf("Prompt for output path: %w", err)
	}

	// if err := VerifyDirectory(outputPath); err != nil {
	// 	return fmt.Errorf("Failed to verify directory: %w", err)
	// }

	o.outputPath = outputPath
	return nil
}

func (o *askOpts) AskSecretCmdName() error {
	secrets := getSecretNames()
	secretCmdName, err := o.prompt.SelectOne(inputSecretCmdNamePrompt, `
docker-registry Create a secret for use with a Docker registry
generic         Create a secret from a local file, directory or literal value
tls             Create a TLS secret`, secrets)

	if err != nil {
		return fmt.Errorf("Prompt for secret cmd name: %w", err)
	}

	o.secretCmdName = secretCmdName

	return nil
}

func (o *askOpts) AskSecretName() error {
	secretName, err := o.prompt.Get(inputSecretNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for secret name: %w", err)
	}
	o.secretName = secretName

	return nil
}

func (o *askOpts) AskDockerServerName() error {
	dockerServer, err := o.prompt.Get(inputDockerServerNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for docker server name: %w", err)
	}
	o.dockerServer = dockerServer

	return nil
}

func (o *askOpts) AskDockerUserName() error {
	dockerUserName, err := o.prompt.Get(inputDockerUserNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for docker username: %w", err)
	}
	o.dockerUserName = dockerUserName

	return nil
}

func (o *askOpts) AskDockerUserPassword() error {
	dockerUserPassword, err := o.prompt.GetSecret(inputDockerUserPasswordPrompt, "")
	if err != nil {
		return fmt.Errorf("Prompt for docker password: %w", err)
	}
	o.dockerUserPassword = dockerUserPassword

	return nil
}

func (o *askOpts) AskDockerEmail() error {
	// TODO: email validation
	dockerEmail, err := o.prompt.Get(inputDockerEmailPrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for docker email: %w", err)
	}
	o.dockerEmail = dockerEmail

	return nil
}

func (o *askOpts) AskCertPath() error {
	certPath, err := o.prompt.Get(inputDockerUserNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for cert path: %w", err)
	}
	o.certPath = certPath

	return nil
}

func (o *askOpts) AskKeyPath() error {
	keyPath, err := o.prompt.Get(inputDockerUserNamePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for key path: %w", err)
	}
	o.keyPath = keyPath

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

func (o *askOpts) AskFromEnv() error {
	fromEnvFile, err := o.prompt.Get(inputFromEnvFilePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for env: %w", err)
	}
	o.fromEnvFile = fromEnvFile

	return nil
}

func (o *askOpts) AskOutputInfo() error {
	if err := o.AskOutputFormat(); err != nil {
		return err
	}

	if err := o.AskOutputPath(); err != nil {
		return err
	}

	return nil
}

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

// ask for building k8s spec

func (o *askOpts) AskRequireObjectMeta() error {
	opts := getYesNoSelectOpts()
	requireObjectMeta, err := o.prompt.SelectOne(inputRequireObjectMetaPrompt, "" /* TODO: may add help text*/, opts)
	if err != nil {
		return fmt.Errorf("Prompt for AskRequireObjectMeta: %w", err)
	}

	if requireObjectMeta == "Yes" {
		o.requireObjectMeta = true
	} else if requireObjectMeta == "No" {
		o.requireObjectMeta = false
	} else {
		return fmt.Errorf("Unexpected requireObjectMeta")
	}

	return nil
}

func (o *askOpts) AskRequireDeploymentSpec() error {
	opts := getYesNoSelectOpts()
	requireDeploymentSpec, err := o.prompt.SelectOne(inputRequireDeploymentSpecPrompt, "" /* TODO: may add help text*/, opts)
	if err != nil {
		return fmt.Errorf("Prompt for AskRequireDeploymentSpec: %w", err)
	}

	if requireDeploymentSpec == "Yes" {
		o.requireDeploymentSpec = true
	} else if requireDeploymentSpec == "No" {
		o.requireDeploymentSpec = false
	} else {
		return fmt.Errorf("Unexpected requireDeploymentSpec")
	}

	return nil
}

func (o *askOpts) AskRequireDeploymentStatus() error {
	opts := getYesNoSelectOpts()
	requireDeploymentStatus, err := o.prompt.SelectOne(inputRequireDeploymentStatusPrompt, "" /* TODO: may add help text*/, opts)
	if err != nil {
		return fmt.Errorf("Prompt for AskRequireDeploymentStatus: %w", err)
	}

	if requireDeploymentStatus == "Yes" {
		o.requireDeploymentStatus = true
	} else if requireDeploymentStatus == "No" {
		o.requireDeploymentStatus = false
	} else {
		return fmt.Errorf("Unexpected requireDeploymentStatus")
	}

	return nil
}

func (o *askOpts) AskNamespace() error {
	namespace, err := o.prompt.Get(inputNamespacePrompt, "", nil /*no validation*/)
	if err != nil {
		return fmt.Errorf("Prompt for env: %w", err)
	}
	o.Namespace = namespace

	return nil
}
