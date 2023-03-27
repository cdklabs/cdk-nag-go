// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag


// Possible output formats of the NagReport.
type NagReportFormat string

const (
	NagReportFormat_CSV NagReportFormat = "CSV"
	NagReportFormat_JSON NagReportFormat = "JSON"
)

