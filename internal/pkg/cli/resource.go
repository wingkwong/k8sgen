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
