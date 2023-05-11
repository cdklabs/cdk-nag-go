package cdknag


// Props for the NagReportLogger.
type NagReportLoggerProps struct {
	Formats *[]NagReportFormat `field:"required" json:"formats" yaml:"formats"`
}

