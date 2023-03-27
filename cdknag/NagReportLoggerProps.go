// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag


// Props for the NagReportLogger.
type NagReportLoggerProps struct {
	Formats *[]NagReportFormat `field:"required" json:"formats" yaml:"formats"`
}

