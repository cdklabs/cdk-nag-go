package cdknag


// Experimental.
type NagReportSchema struct {
	// Experimental.
	Lines *[]*NagReportLine `field:"required" json:"lines" yaml:"lines"`
}

