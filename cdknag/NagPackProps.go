package cdknag


// Interface for creating a NagPack.
// Experimental.
type NagPackProps struct {
	// Whether or not to enable extended explanatory descriptions on warning, error, and logged ignore messages (default: false).
	// Experimental.
	Verbose *bool `field:"optional" json:"verbose" yaml:"verbose"`
	// Whether to write acknowledged rules into CfnResource CloudFormation Metadata as `cdk_nag: { rules_to_suppress: [...] }` for backwards compatibility with v2 audit trail tooling (default: false).
	// Experimental.
	WriteSuppressionsToCloudFormation *bool `field:"optional" json:"writeSuppressionsToCloudFormation" yaml:"writeSuppressionsToCloudFormation"`
}

