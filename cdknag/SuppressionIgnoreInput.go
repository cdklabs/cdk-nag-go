// Check CDK v2 applications for best practices using a combination on available rule packs.
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

