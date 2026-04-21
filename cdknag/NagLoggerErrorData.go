package cdknag

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Data for onError method of an INagLogger.
// Experimental.
type NagLoggerErrorData struct {
	// Experimental.
	NagPackName *string `field:"required" json:"nagPackName" yaml:"nagPackName"`
	// Experimental.
	Resource awscdk.CfnResource `field:"required" json:"resource" yaml:"resource"`
	// Experimental.
	RuleExplanation *string `field:"required" json:"ruleExplanation" yaml:"ruleExplanation"`
	// Experimental.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// Experimental.
	RuleInfo *string `field:"required" json:"ruleInfo" yaml:"ruleInfo"`
	// Experimental.
	RuleLevel NagMessageLevel `field:"required" json:"ruleLevel" yaml:"ruleLevel"`
	// Experimental.
	RuleOriginalName *string `field:"required" json:"ruleOriginalName" yaml:"ruleOriginalName"`
	// Experimental.
	ErrorMessage *string `field:"required" json:"errorMessage" yaml:"errorMessage"`
}

