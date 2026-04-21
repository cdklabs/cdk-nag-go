package cdknag


// Experimental.
type NagReportLine struct {
	// Experimental.
	Compliance *string `field:"required" json:"compliance" yaml:"compliance"`
	// Experimental.
	ExceptionReason *string `field:"required" json:"exceptionReason" yaml:"exceptionReason"`
	// Experimental.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Experimental.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Experimental.
	RuleInfo *string `field:"required" json:"ruleInfo" yaml:"ruleInfo"`
	// Experimental.
	RuleLevel *string `field:"required" json:"ruleLevel" yaml:"ruleLevel"`
}

