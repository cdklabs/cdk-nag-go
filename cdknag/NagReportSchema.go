// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag


type NagReportSchema struct {
	Lines *[]*NagReportLine `field:"required" json:"lines" yaml:"lines"`
}

