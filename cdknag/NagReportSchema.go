package cdknag


// Schema for the NagReport output.
// Experimental.
type NagReportSchema struct {
	// Experimental.
	Lines *[]*NagReportLine `field:"required" json:"lines" yaml:"lines"`
}

