package cdknag


// Interface for creating a NagPack.
type NagPackProps struct {
	// Additional NagLoggers for logging rule validation outputs.
	AdditionalLoggers *[]INagLogger `field:"optional" json:"additionalLoggers" yaml:"additionalLoggers"`
	// Whether or not to log suppressed rule violations as informational messages (default: false).
	LogIgnores *bool `field:"optional" json:"logIgnores" yaml:"logIgnores"`
	// If reports are enabled, the output formats of compliance reports in the App's output directory (default: only CSV).
	ReportFormats *[]NagReportFormat `field:"optional" json:"reportFormats" yaml:"reportFormats"`
	// Whether or not to generate compliance reports for applied Stacks in the App's output directory (default: true).
	Reports *bool `field:"optional" json:"reports" yaml:"reports"`
	// Conditionally prevent rules from being suppressed (default: no user provided condition).
	SuppressionIgnoreCondition INagSuppressionIgnore `field:"optional" json:"suppressionIgnoreCondition" yaml:"suppressionIgnoreCondition"`
	// Whether or not to enable extended explanatory descriptions on warning, error, and logged ignore messages (default: false).
	Verbose *bool `field:"optional" json:"verbose" yaml:"verbose"`
}

