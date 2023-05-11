package cdknag

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Shared data for all INagLogger methods.
type NagLoggerBaseData struct {
	NagPackName *string `field:"required" json:"nagPackName" yaml:"nagPackName"`
	Resource awscdk.CfnResource `field:"required" json:"resource" yaml:"resource"`
	RuleExplanation *string `field:"required" json:"ruleExplanation" yaml:"ruleExplanation"`
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	RuleInfo *string `field:"required" json:"ruleInfo" yaml:"ruleInfo"`
	RuleLevel NagMessageLevel `field:"required" json:"ruleLevel" yaml:"ruleLevel"`
}

