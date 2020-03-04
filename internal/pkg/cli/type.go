package cli


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
	DeploymentCmdOpts
	SecretCmdOpts

	DeploymentOpts

	KindName string
}

type askOpts struct {
	askVars
}

type DeploymentCmdOpts struct {
	// Name of the Deployment
	deploymentName string
	// Name of the image
	imageName string
}

type SecretCmdOpts struct {
	// Type of Secret
	secretCmdName string
	// Name of Secret to-be-created
	secretName string
	// Server location for Docker registry
	dockerServer string
	// Username for Docker registry authentication
	dockerUserName string
	// Password for Docker registry authentication
	dockerUserPassword string
	// Email for Docker registry
	dockerEmail string
	// Key files can be specified using their file path, in which case a default name will be given to them
	// or optionally with a name and file path, in which case the given name will be used
	// Specifying a directory will iterate each named file in the directory that is a valid secret key
	fromFile []string
	// Append a hash of the secret to its name
	appendHash bool
	// Secret cert path
	certPath string
	// Secret key path
	keyPath string
	// Secret from Literal input
	fromLiteral []string
	// Secret from environment file
	fromEnvFile string
	// Number of iteration for the same question for fromFile
	noOfFromFileIteration int
	// Number of iteration for the same question for fromLiteral
	noOfFromLiteralIteration int
}