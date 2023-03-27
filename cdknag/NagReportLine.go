// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag


type NagReportLine struct {
	Compliance *string `field:"required" json:"compliance" yaml:"compliance"`
	ExceptionReason *string `field:"required" json:"exceptionReason" yaml:"exceptionReason"`
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	RuleInfo *string `field:"required" json:"ruleInfo" yaml:"ruleInfo"`
	RuleLevel *string `field:"required" json:"ruleLevel" yaml:"ruleLevel"`
}

