package cdknag

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Data for onSuppressed method of an INagLogger.
type NagLoggerSuppressedData struct {
	NagPackName *string `field:"required" json:"nagPackName" yaml:"nagPackName"`
	Resource awscdk.CfnResource `field:"required" json:"resource" yaml:"resource"`
	RuleExplanation *string `field:"required" json:"ruleExplanation" yaml:"ruleExplanation"`
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	RuleInfo *string `field:"required" json:"ruleInfo" yaml:"ruleInfo"`
	RuleLevel NagMessageLevel `field:"required" json:"ruleLevel" yaml:"ruleLevel"`
	FindingId *string `field:"required" json:"findingId" yaml:"findingId"`
	SuppressionReason *string `field:"required" json:"suppressionReason" yaml:"suppressionReason"`
}

