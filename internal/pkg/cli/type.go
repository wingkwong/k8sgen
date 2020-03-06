package cli

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GlobalOpts struct {
	// Global Prompter
	prompt prompter
	// Output Format: json / yaml
	outputFormat string
	// Output Directory where a resource file is saved
	outputPath string
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
	// Create Opts
	DeploymentOpts
	//External
	k8s
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
	// Key files can be specified using their file path, in which case a default name will be given to them
	// or optionally with a name and file path, in which case the given name will be used
	// Specifying a directory will iterate each named file in the directory that is a valid secret key
	FromFile []string
	// Append a hash of the secret to its name
	AppendHash bool
	// Secret cert path
	CertPath string
	// Secret key path
	KeyPath string
	// Secret from Literal input
	FromLiteral []string
	// Secret from environment file
	FromEnvFile string
	// Number of iteration for the same question for fromFile
	NoOfFromFileIteration int
	// Number of iteration for the same question for fromLiteral
	NoOfFromLiteralIteration int
}

type DeploymentOpts struct {
	// Deployment
	RequireObjectMeta       bool
	RequireDeploymentSpec   bool
	RequireDeploymentStatus bool
	// DeploymentSpec
	DeploymentSpecOpts
	// DeploymentStatus
	DeploymentStatusOpts
}

type DeploymentSpecOpts struct {
	RequireMoreThanOneReplica   bool
	RequireDeploymentStrategy   bool
	RequireMinReadySeconds      bool
	RequireRevisionHistoryLimit bool
	RequirePaused               bool
}

type DeploymentStatusOpts struct {
	RequireObservedGeneration  bool
	RequireReplicas            bool
	RequireUpdatedReplicas     bool
	RequireReadyReplicas       bool
	RequireAvailableReplicas   bool
	RequireUnavailableReplicas bool
	RequireCollisionCount      bool
}

type k8s struct {
	ObjectMeta       metav1.ObjectMeta
	TypeMeta         metav1.TypeMeta
	Deployment       appsv1.Deployment
	DeploymentSpec   appsv1.DeploymentSpec
	DeploymentStatus appsv1.DeploymentStatus
}
