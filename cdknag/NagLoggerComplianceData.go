// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Data for onCompliance method of an INagLogger.
type NagLoggerComplianceData struct {
	NagPackName *string `field:"required" json:"nagPackName" yaml:"nagPackName"`
	Resource awscdk.CfnResource `field:"required" json:"resource" yaml:"resource"`
	RuleExplanation *string `field:"required" json:"ruleExplanation" yaml:"ruleExplanation"`
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	RuleInfo *string `field:"required" json:"ruleInfo" yaml:"ruleInfo"`
	RuleLevel NagMessageLevel `field:"required" json:"ruleLevel" yaml:"ruleLevel"`
}

