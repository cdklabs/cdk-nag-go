package cdknag


// Props for the AnnotationLogger.
// Experimental.
type AnnotationLoggerProps struct {
	// Whether or not to log suppressed rule violations as informational messages (default: false).
	// Experimental.
	LogIgnores *bool `field:"optional" json:"logIgnores" yaml:"logIgnores"`
	// Whether or not to enable extended explanatory descriptions on warning, error, and logged ignore messages.
	// Experimental.
	Verbose *bool `field:"optional" json:"verbose" yaml:"verbose"`
}

