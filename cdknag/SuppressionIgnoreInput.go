package cdknag

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Information about the NagRule and the relevant NagSuppression for the INagSuppressionIgnore.
// Experimental.
type SuppressionIgnoreInput struct {
	// Experimental.
	FindingId *string `field:"required" json:"findingId" yaml:"findingId"`
	// Experimental.
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// Experimental.
	Resource awscdk.CfnResource `field:"required" json:"resource" yaml:"resource"`
	// Experimental.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Experimental.
	RuleLevel NagMessageLevel `field:"required" json:"ruleLevel" yaml:"ruleLevel"`
}

