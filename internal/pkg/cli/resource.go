package cli

// Kind
const (
	clusterRoleName         = "ClusterRole"
	clusterRoleBindingName  = "ClusterRoleBinding"
	configmapName           = "ConfigMap"
	deploymentName          = "Deployment"
	jobName                 = "Job"
	namespaceName           = "Namespace"
	podDisruptionBudgetName = "PodDisruptionBudget"
	priorityClassName       = "PriorityClass"
	quotaName               = "Quota"
	roleName                = "Role"
	roleBindingName         = "RoleBinding"
	secretName              = "Secret"
	serviceName             = "Service"
	serviceAccountName      = "ServiceAccount"
)

// Kind
const (
	clusterRoleCmdName         = "clusterrole"
	clusterRoleBindingCmdName  = "clusterrolebinding"
	configmapCmdName           = "configmap"
	deploymentCmdName          = "deployment"
	jobCmdName                 = "job"
	cmdNamespaceCmdName        = "namespace"
	podDisruptionBudgetCmdName = "poddisruptionbudget"
	priorityClassCmdName       = "priorityclass"
	quotaCmdName               = "quota"
	roleCmdName                = "role"
	roleBindingCmdName         = "rolebinding"
	secretCmdName              = "secret"
	serviceCmdName             = "service"
	serviceAccountCmdName      = "serviceaccount"
)

// Output Format
const (
	jsonName = "json"
	yamlName = "yaml"
	// nameName           = "name"
	// goTemplateName     = "go-template"
	// goTemplateFileName = "go-template-file"
	// templateName       = "template"
	// templateFileName   = "templatefile"
	// jsonPathName       = "jsonpath"
	// jsonPathFileName   = "jsonpath-file"
)

// Secret
const (
	dockerRegistryCmdName = "docker-registry"
	genericCmdName        = "generic"
	tlsCmdName            = "tls"

	// Generic Opts
	genericOpt1 = "Create a new secret with keys for each file in folder"
	genericOpt2 = "Create a new secret with specified keys instead of names on disk"
	genericOpt3 = "Create a new secret with keys"
	genericOpt4 = "Create a new secret using a combination of a file and a literal"
	genericOpt5 = "Create a new secret from an env file"
)

func getKindNames() []string {
	return []string{
		clusterRoleName,
		clusterRoleBindingName,
		configmapName,
		deploymentName,
		jobName,
		namespaceName,
		podDisruptionBudgetName,
		priorityClassName,
		quotaName,
		roleName,
		roleBindingName,
		secretName,
		serviceName,
		serviceAccountName,
	}
}

func getOutputFormats() []string {
	return []string{
		jsonName,
		yamlName,
		// nameName,
		// goTemplateName,
		// goTemplateFileName,
		// templateName,
		// templateFileName,
		// jsonPathName,
		// jsonPathFileName,
	}
}

func getSecretNames() []string {
	return []string{
		dockerRegistryCmdName,
		genericCmdName,
		tlsCmdName,
	}
}

func getSecretGenericOpts() []string {
	return []string{
		genericOpt1,
		genericOpt2,
		genericOpt3,
		genericOpt4,
		genericOpt5,
	}
}
