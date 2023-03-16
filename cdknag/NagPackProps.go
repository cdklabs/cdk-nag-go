// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag


// Interface for creating a Nag rule pack.
type NagPackProps struct {
	// Whether or not to log triggered rules that have been suppressed as informational messages (default: false).
	LogIgnores *bool `field:"optional" json:"logIgnores" yaml:"logIgnores"`
	// Whether or not to generate CSV compliance reports for applied Stacks in the App's output directory (default: true).
	Reports *bool `field:"optional" json:"reports" yaml:"reports"`
	// Conditionally prevent rules from being suppressed (default: no user provided condition).
	SuppressionIgnoreCondition INagSuppressionIgnore `field:"optional" json:"suppressionIgnoreCondition" yaml:"suppressionIgnoreCondition"`
	// Whether or not to enable extended explanatory descriptions on warning, error, and logged ignore messages (default: false).
	Verbose *bool `field:"optional" json:"verbose" yaml:"verbose"`
}

