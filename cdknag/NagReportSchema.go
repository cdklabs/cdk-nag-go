package cdknag


type NagReportSchema struct {
	Lines *[]*NagReportLine `field:"required" json:"lines" yaml:"lines"`
}

