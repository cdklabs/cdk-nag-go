package cdknag

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Information about the NagRule and the relevant NagSuppression for the INagSuppressionIgnore.
type SuppressionIgnoreInput struct {
	FindingId *string `field:"required" json:"findingId" yaml:"findingId"`
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	Resource awscdk.CfnResource `field:"required" json:"resource" yaml:"resource"`
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	RuleLevel NagMessageLevel `field:"required" json:"ruleLevel" yaml:"ruleLevel"`
}

