package cli

type GlobalOpts struct {
	// Global Prompter
	prompt prompter
	// Output Format: json / yaml
	OutputFormat string
	// Output Directory where a resource file is saved
	OutputPath string
}

type askVars struct {
	*GlobalOpts
	KindName          string
	Iterator          int
	intPlaceholder    int
	stringPlaceholder string
	boolPlaceholder   bool
	// Cmd Opts
	DeploymentCmdOpts
	SecretCmdOpts
	RoleCmdOpts
	QuotaCmdOpts
	PriorityClassCmtOpts
	NamespaceCmtOpts
	ConfigMapCmtOpts
	ServiceCmdOpts
	JobCmtOpts
	PodDisruptionBudget
	ServiceAccountCmdOpts
	FileOpts
}

type askOpts struct {
	askVars
}

type ValidatorFunc func(interface{})

type getSelectOptsFn func()

type DeploymentCmdOpts struct {
	// Name of the Deployment
	DeploymentName string
	// Name of the image
	ImageName string
}

type SecretCmdOpts struct {
	// Type of Secret
	SecretCmdName string
	// Name of Secret to-be-created
	SecretName string
	// Server location for Docker registry
	DockerServer string
	// Username for Docker registry authentication
	DockerUserName string
	// Password for Docker registry authentication
	DockerUserPassword string
	// Email for Docker registry
	DockerEmail string
	// Secret cert path
	CertPath string
	// Secret key path
	KeyPath string
}

type RoleCmdOpts struct {
	// Name Of Role
	RoleName string
	// Resource that the rule applies to
	Resource string
	// Resource in the white list that the rule applies to, repeat this flag for multiple items
	ResourceName string
	// Verb that applies to the resources contained in the rule
	Verb string
}

type QuotaCmdOpts struct {
	// Name of Quota
	QuotaName string
	// A comma-delimited set of resource=quantity pairs that define a hard limit.
	Hard string
	// A comma-delimited set of quota scopes that must all match each object tracked by the quota.
	Scopes string
}

type PriorityClassCmtOpts struct {
	// Name of Priority Class
	PriorityClassName string
	// The value of this priority class.
	Value int
	// Description is an arbitrary string that usually provides guidelines on when this priority class should be used
	Description string
	// Global-default specifies whether this PriorityClass should be considered as the default priority.
	GlobalDefault bool
	// Preemption-policy is the policy for preempting pods with lower priority
	PreemptionPolicy string
}

type NamespaceCmtOpts struct {
	// Name Of Namespace
	NamespaceName string
}

type ConfigMapCmtOpts struct {
	// Name Of Config Map
	ConfigMapName string
}

type ServiceCmdOpts struct {
	// Name of Service CMD Name
	ServiceCmdName string
	// Port pairs can be specified as '<port>:<targetPort>'.
	TCP string

	// clusterip
	// -------------
	// Name of Cluster IP
	ClusterIPName string
	// Assign your own ClusterIP or set to 'None' for a 'headless' service (no loadbalancing).
	ClusterIP string

	// externalname
	// -------------
	// Name of ExternalName
	ExternalName string
	// External name of service
	ExternalServiceName string

	// loadbalancer
	// -------------
	// Name of Load Balancer
	LoadbalancerName string

	// nodeport
	// -------------
	// Name of Node Port
	NodePortName string
}

type JobCmtOpts struct {
	// Name of Job
	JobName string
	// Command to-be-run
	Command string
	// The name of the resource to create a Job from (only cronjob is supported).
	FromResource string
}

type PodDisruptionBudget struct {
	// Name of Pod Disruption Budget
	PodDisruptionBudgetName string
	// The maximum number or percentage of unavailable pods this budget requires.
	MaxUnavailable string
	// The minimum number or percentage of available pods this budget requires.
	MinAvailable string
	// A label selector to use for this budget. Only equality-based selector requirements are supported.
	Selector string
}

type ServiceAccountCmdOpts struct {
	// Name of Service Account
	ServiceAccountName string
}

type FileOpts struct {
	// Key files can be specified using their file path, in which case a default name will be given to them
	// or optionally with a name and file path, in which case the given name will be used
	// Specifying a directory will iterate each named file in the directory that is a valid secret key
	FromFile []string
	// Specify a key and literal value to insert (i.e. mykey=somevalue)
	FromLiteral []string
	// Specify the path to a file to read lines of key=val pairs (i.e. a Docker .env file).
	FromEnvFile string
	// Number of iteration for the same question for fromFile
	FromFileIteration int
	// Number of iteration for the same question for fromLiteral
	FromLiteralIteration int
	// Append a hash to its name
	AppendHash bool
}
