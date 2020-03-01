package cli

const (
	clusterRoleName         = "ClusterRole"
	clusterRoleBindingName  = "ClusterRoleBinding"
	configmapName           = "Configmap"
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

const (
	jsonName           = "json"
	yamlName           = "yaml"
	// nameName           = "name"
	// goTemplateName     = "go-template"
	// goTemplateFileName = "go-template-file"
	// templateName       = "template"
	// templateFileName   = "templatefile"
	// jsonPathName       = "jsonpath"
	// jsonPathFileName   = "jsonpath-file"
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
