// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag


// Props for the AnnotationLogger.
type AnnotationLoggerProps struct {
	// Whether or not to log suppressed rule violations as informational messages (default: false).
	LogIgnores *bool `field:"optional" json:"logIgnores" yaml:"logIgnores"`
	// Whether or not to enable extended explanatory descriptions on warning, error, and logged ignore messages.
	Verbose *bool `field:"optional" json:"verbose" yaml:"verbose"`
}

