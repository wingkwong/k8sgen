package cli

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/wingkwong/k8sgen/third_party/term/prompt"
)

// promptMessage
const (
	// text input
	inputKindNamePrompt                = "What kind of object you want to create?"
	inputDeploymentNamePrompt          = "What deployment you want to name?"
	inputImageNamePrompt               = "What image you want to name to run?"
	inputOutputPathPrompt              = "What directory you want to save?"
	inputSecretNamePrompt              = "What secret you want to name?"
	inputDockerServerNamePrompt        = "What docker server you want to name?"
	inputDockerUserNamePrompt          = "What is the username for Docker registry authentiation?"
	inputDockerUserPasswordPrompt      = "What is the password for Docker registry authentiation?"
	inputDockerEmailPrompt             = "What is the email address for Docker registry?"
	inputCertPathPrompt                = "Where is the cert path?"
	inputKeyPathPrompt                 = "Where is the key path?"
	inputFromFilePrompt                = "Input key names from file: (e.g path/to/bar):"
	inputFromLiteralPrompt             = "Input a key-value pair secret (e.g foo='bar'):"
	inputFromEnvFilePrompt             = "Where is the env file path?"
	inputFromFileIterationPrompt       = "How many from-file iterations for your input?"
	inputFromLiteralIterationPrompt    = "How many from-literal iterations for your input?"
	inputRoleNamePrompt                = "What role you want to name?"
	inputResourcePrompt                = "What resource you want to name?"
	inputResourceNamePrompt            = "What resource name you want to name?"
	inputVerbPrompt                    = "What verb you want to name?"
	inputQuotaPrompt                   = "What quota you want to name?"
	inputHardPrompt                    = "What hard you want to name?"
	inputScopesPrompt                  = "What scopes you want to name?"
	inputPriorityClassPrompt           = "What priority class you want to name?"
	inputValuePrompt                   = "What value you want to have?"
	inputDescriptionPrompt             = "What description you want to have?"
	inputGlobalDefaultPrompt           = "Do you want this PriorityClass to be considered as the default priority?"
	inputPreemptionPolicyPrompt        = "What preemption-policy you want to have?"
	inputNamespacePrompt               = "What namespace you want to name?"
	inputConfigMapNamePrompt           = "What config map you want to name?"
	inputAppenHashPrompt               = "Do you want to append a hash to its name?"
	inputTCPPrompt                     = "Input Port pairs (e.g <port>:<targetPort>)"
	inputClusterIPNamePrompt           = "What cluster IP name you want to name?"
	inputClusterIPPrompt               = "What cluster IP you want to assign?"
	inputExternalNamePrompt            = "What external name you want to name?"
	inputExternalServiceNamePrompt     = "What external service name you want to name?"
	inputLoadbalancerNamePrompt        = "What load balancer name you want to name?"
	inputNodePortNamePrompt            = "What node port you want to use?"
	inputJobNamePrompt                 = "What job you want to name?"
	inputCommandPrompt                 = "What command you want to run?"
	inputFromPrompt                    = "What resource to create a Job from (only cronjob is supported)?"
	inputPodDisruptionBudgetNamePrompt = "What pod disruption budget you wnat name?"
	inputMaxUnavailablePrompt          = "What is the maximum number or percentage of unavailable pods this budget requires?"
	inputMinAvailablePrompt            = "What is the minimum number or percentage of available pods this budget requires?"
	inputSelectorPrompt                = "What label selector to use for this budget?"

	// select
	inputOutputFormatPrompt   = "Please select an output format:"
	inputSecretCmdNamePrompt  = "Please select the type of secret:"
	inputServiceCmdNamePrompt = "Please select the type of service:"
)

// promptHelpMessage
const (
	inputKindNamePromptHelpMessage                = ""
	inputDeploymentNamePromptHelpMessage          = ""
	inputImageNamePromptHelpMessage               = ""
	inputOutputPathPromptHelpMessage              = ""
	inputSecretNamePromptHelpMessage              = ""
	inputDockerServerNamePromptHelpMessage        = ""
	inputDockerUserNamePromptHelpMessage          = ""
	inputDockerUserPasswordPromptHelpMessage      = ""
	inputDockerEmailPromptHelpMessage             = ""
	inputCertPathPromptHelpMessage                = ""
	inputKeyPathPromptHelpMessage                 = ""
	inputFromFilePromptHelpMessage                = ""
	inputFromLiteralPromptHelpMessage             = ""
	inputFromEnvFilePromptHelpMessage             = ""
	inputFromFileIterationPromptHelpMessage       = ""
	inputFromLiteralIterationPromptHelpMessage    = ""
	inputRoleNamePromptHelpMessage                = ""
	inputResourcePromptHelpMessage                = ""
	inputResourceNamePromptHelpMessage            = ""
	inputVerbPromptHelpMessage                    = ""
	inputQuotaPromptHelpMessage                   = ""
	inputHardPromptHelpMessage                    = ""
	inputScopesPromptHelpMessage                  = ""
	inputPriorityClassPromptHelpMessage           = ""
	inputValuePromptHelpMessage                   = ""
	inputDescriptionPromptHelpMessage             = ""
	inputGlobalDefaultPromptHelpMessage           = ""
	inputPreemptionPolicyPromptHelpMessage        = ""
	inputNamespacePromptHelpMessage               = ""
	inputConfigMapNamePromptHelpMessage           = ""
	inputAppenHashPromptHelpMessage               = ""
	inputOutputFormatPromptHelpMessage            = ""
	inputSecretCmdNamePromptHelpMessage           = ""
	inputServiceCmdNamePromptHelpMessage          = ""
	inputTCPPromptHelpMessage                     = ""
	inputClusterIPNamePromptHelpMessage           = ""
	inputClusterIPPromptHelpMessage               = ""
	inputExternalNamePromptHelpMessage            = ""
	inputExternalServiceNamePromptHelpMessage     = ""
	inputLoadbalancerNamePromptHelpMessage        = ""
	inputNodePortNamePromptHelpMessage            = ""
	inputJobNamePromptHelpMessage                 = "Name of Job"
	inputCommandPromptHelpMessage                 = "Command to-be-run"
	inputFromPromptHelpMessage                    = "The name of the resource to create a Job from (only cronjob is supported)."
	inputPodDisruptionBudgetNamePromptHelpMessage = "Name of Pod Disruption Budget"
	inputMaxUnavailablePromptHelpMessage          = "The maximum number or percentage of unavailable pods this budget requires"
	inputMinAvailablePromptHelpMessage            = "The minimum number or percentage of available pods this budget requires"
	inputSelectorPromptHelpMessage                = "A label selector to use for this budget. Only equality-based selector requirements are supported"
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
	"Kind":                    {"KindName", "string", inputKindNamePrompt, inputKindNamePromptHelpMessage, "Select kind name", nil /*no validation*/, kindNames, "AskSelect"},
	"DeploymentName":          {"DeploymentName", "string", inputDeploymentNamePrompt, inputDeploymentNamePromptHelpMessage, "Prompt for deployment name", validateDeploymentName, nil, "AskGet"},
	"Image":                   {"ImageName", "string", inputImageNamePrompt, inputImageNamePromptHelpMessage, "Prompt for image name", nil /*no validation*/, nil, "AskGet"},
	"OutputPath":              {"OutputPath", "string", inputOutputPathPrompt, inputOutputPathPromptHelpMessage, "Prompt for output path", nil /*no validation*/, nil, "AskGet"},
	"SecretName":              {"SecretName", "string", inputSecretNamePrompt, inputSecretNamePromptHelpMessage, "Prompt for secret name", nil /*no validation*/, nil, "AskGet"},
	"DockerServerName":        {"DockerServer", "string", inputDockerServerNamePrompt, inputDockerServerNamePromptHelpMessage, "Prompt for docker server name", nil /*no validation*/, nil, "AskGet"},
	"DockerUserName":          {"DockerUserName", "string", inputDockerUserNamePrompt, inputDockerUserNamePromptHelpMessage, "Prompt for docker user name", nil /*no validation*/, nil, "AskGet"},
	"DockerPassword":          {"DockerUserPassword", "string", inputDockerUserPasswordPrompt, inputDockerUserPasswordPromptHelpMessage, "Prompt for docker password", nil /*no validation*/, nil, "AskGetSecret"},
	"DockerEmail":             {"DockerEmail", "string", inputDockerEmailPrompt, inputDockerEmailPromptHelpMessage, "Prompt for docker email", nil /*no validation*/, nil, "AskGet"}, // TODO: email validation
	"CertPath":                {"CertPath", "string", inputCertPathPrompt, inputCertPathPromptHelpMessage, "Prompt for cert path", nil /*no validation*/, nil, "AskGet"},
	"KeyPath":                 {"KeyPath", "string", inputKeyPathPrompt, inputKeyPathPromptHelpMessage, "Prompt for key path", nil /*no validation*/, nil, "AskGet"},
	"FromEnvFile":             {"FromEnvFile", "string", inputFromEnvFilePrompt, inputFromEnvFilePromptHelpMessage, "Prompt for env", nil /*no validation*/, nil, "AskGet"},
	"OutputFormat":            {"OutputFormat", "string", inputOutputFormatPrompt, inputOutputFormatPromptHelpMessage, "Prompt for output format", nil /*no validation*/, outputFormats, "AskSelect"},
	"SecretCmdName":           {"SecretCmdName", "string", inputSecretCmdNamePrompt, inputSecretCmdNamePromptHelpMessage, "Prompt for secret cmd name", nil /*no validation*/, secretNames, "AskSelect"},
	"RoleName":                {"RoleName", "string", inputRoleNamePrompt, inputRoleNamePromptHelpMessage, "Prompt for role name", nil /*no validation*/, nil, "AskGet"},
	"Resource":                {"Resource", "string", inputResourcePrompt, inputResourcePromptHelpMessage, "Prompt for resource", nil /*no validation*/, nil, "AskGet"},
	"ResourceName":            {"ResourceName", "string", inputResourceNamePrompt, inputResourceNamePromptHelpMessage, "Prompt for resource name", nil /*no validation*/, nil, "AskGet"},
	"Verb":                    {"Verb", "string", inputVerbPrompt, inputVerbPromptHelpMessage, "Prompt for verb", nil /*no validation*/, nil, "AskGet"},
	"QuotaName":               {"QuotaName", "string", inputQuotaPrompt, inputQuotaPromptHelpMessage, "Prompt for quota name", nil /*no validation*/, nil, "AskGet"},
	"Hard":                    {"Hard", "string", inputHardPrompt, inputHardPromptHelpMessage, "Prompt for hard", nil /*no validation*/, nil, "AskGet"},
	"Scopes":                  {"Scopes", "string", inputScopesPrompt, inputScopesPromptHelpMessage, "Prompt for scopes", nil /*no validation*/, nil, "AskGet"},
	"PriorityClassName":       {"PriorityClassName", "string", inputPriorityClassPrompt, inputPriorityClassPromptHelpMessage, "Prompt for priority class", nil /*no validation*/, nil, "AskGet"},
	"Value":                   {"Value", "int", inputValuePrompt, inputValuePromptHelpMessage, "Prompt for value", nil /*no validation*/, nil, "AskGet"},
	"Description":             {"Description", "string", inputDescriptionPrompt, inputDescriptionPromptHelpMessage, "Prompt for description", nil /*no validation*/, nil, "AskGet"},
	"GlobalDefault":           {"GlobalDefault", "bool", inputGlobalDefaultPrompt, inputGlobalDefaultPromptHelpMessage, "Prompt for global default", nil /*no validation*/, yesOrNo, "AskSelect"},
	"AppendHash":              {"AppendHash", "bool", inputAppenHashPrompt, inputAppenHashPromptHelpMessage, "Prompt for append hash", nil /*no validation*/, yesOrNo, "AskSelect"},
	"PreemptionPolicy":        {"PreemptionPolicy", "string", inputPreemptionPolicyPrompt, inputPreemptionPolicyPromptHelpMessage, "Prompt for preemption policy", nil /*no validation*/, nil, "AskGet"},
	"NamespaceName":           {"NamespaceName", "string", inputNamespacePrompt, inputNamespacePromptHelpMessage, "Prompt for namespace", nil /*no validation*/, nil, "AskGet"},
	"ConfigMapName":           {"ConfigMapName", "string", inputConfigMapNamePrompt, inputConfigMapNamePromptHelpMessage, "Prompt for config map", nil /*no validation*/, nil, "AskGet"},
	"ServiceCmdName":          {"ServiceCmdName", "string", inputServiceCmdNamePrompt, inputServiceCmdNamePromptHelpMessage, "Prompt for service cmd name", nil /*no validation*/, seviceNames, "AskSelect"},
	"TCP":                     {"TCP", "string", inputTCPPrompt, inputTCPPromptHelpMessage, "Prompt for tcp", nil /*no validation*/, nil, "AskGet"},
	"ClusterIPName":           {"ClusterIPName", "string", inputClusterIPNamePrompt, inputClusterIPNamePromptHelpMessage, "Prompt for cluster ip name", nil /*no validation*/, nil, "AskGet"},
	"ClusterIP":               {"ClusterIP", "string", inputClusterIPPrompt, inputClusterIPPromptHelpMessage, "Prompt for cluster ip", nil /*no validation*/, nil, "AskGet"},
	"ExternalName":            {"ExternalName", "string", inputExternalNamePrompt, inputExternalNamePromptHelpMessage, "Prompt for external name", nil /*no validation*/, nil, "AskGet"},
	"ExternalServiceName":     {"ExternalServiceName", "string", inputExternalServiceNamePrompt, inputExternalServiceNamePromptHelpMessage, "Prompt for external service name", nil /*no validation*/, nil, "AskGet"},
	"LoadbalancerName":        {"LoadbalancerName", "string", inputLoadbalancerNamePrompt, inputLoadbalancerNamePromptHelpMessage, "Prompt for load balancer name", nil /*no validation*/, nil, "AskGet"},
	"NodePortName":            {"NodePortName", "string", inputNodePortNamePrompt, inputNodePortNamePromptHelpMessage, "Prompt for node port name", nil /*no validation*/, nil, "AskGet"},
	"JobName":                 {"JobName", "string", inputJobNamePrompt, inputJobNamePromptHelpMessage, "Prompt for job name", nil /*no validation*/, nil, "AskGet"},
	"Command":                 {"Command", "string", inputCommandPrompt, inputCommandPromptHelpMessage, "Prompt for command", nil /*no validation*/, nil, "AskGet"},
	"FromResource":            {"FromResource", "string", inputFromPrompt, inputFromPromptHelpMessage, "Prompt for from-resource", nil /*no validation*/, nil, "AskGet"},
	"PodDisruptionBudgetName": {"PodDisruptionBudgetName", "string", inputPodDisruptionBudgetNamePrompt, inputPodDisruptionBudgetNamePromptHelpMessage, "Prompt for node port name", nil /*no validation*/, nil, "AskGet"},
	"MaxUnavailable":          {"MaxUnavailable", "string", inputMaxUnavailablePrompt, inputMaxUnavailablePromptHelpMessage, "Prompt for node port name", nil /*no validation*/, nil, "AskGet"},
	"MinAvailable":            {"MinAvailable", "string", inputMinAvailablePrompt, inputMinAvailablePromptHelpMessage, "Prompt for node port name", nil /*no validation*/, nil, "AskGet"},
	"Selector":                {"Selector", "string", inputSelectorPrompt, inputSelectorPromptHelpMessage, "Prompt for node port name", nil /*no validation*/, nil, "AskGet"},
	// Iteration
	"FromFileIteration":    {"Iterator", "int", inputFromFileIterationPrompt, inputFromFileIterationPromptHelpMessage, "Prompt for from-file iteration", nil /*no validation*/, nil, "AskGet"},
	"FromLiteralIteration": {"Iterator", "int", inputFromLiteralIterationPrompt, inputFromLiteralIterationPromptHelpMessage, "Prompt for from-literal iteration", nil /*no validation*/, nil, "AskGet"},
	// Array
	"FromFile":    {"FromFile", "array", inputFromFilePrompt, inputFromFilePromptHelpMessage, "Prompt for from-file", nil /*no validation*/, nil, "AskGet"},
	"FromLiteral": {"FromLiteral", "array", inputFromLiteralPrompt, inputFromLiteralPromptHelpMessage, "Prompt for from-literal", nil /*no validation*/, nil, "AskGet"},
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
	} else if fvType == "array" {
		//TODO: check type
		fv.Set(reflect.Append(fv, reflect.ValueOf(value)))
	} else {
		return fmt.Errorf("%s is not either string, int or bool", fvType)
	}

	return nil
}

func (o *askOpts) getValueByFieldName(v interface{}, name string) (string, error) {
	rv := reflect.ValueOf(v)
	fv := reflect.Indirect(rv).FieldByName(name)
	return fv.String(), nil
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

func (o *askOpts) AskWithIterator(key string) error {
	_, exists := questions[key]

	if !exists {
		return fmt.Errorf("key %s not exists in questions", key)
	}

	iteratorKey := key + "Iteration"
	if err := o.Ask(iteratorKey); err != nil {
		return err
	}

	for i := 0; i < o.Iterator; i++ {
		if err := o.Ask(key); err != nil {
			return err
		}
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

func (o *askOpts) AskOutputInfo() error {

	if err := o.Ask("OutputFormat"); err != nil {
		return err
	}

	if err := o.Ask("OutputPath"); err != nil {
		return err
	}

	return nil
}
