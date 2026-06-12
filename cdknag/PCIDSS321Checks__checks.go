//go:build !no_runtime_type_checking

package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_PCIDSS321Checks) validateApplyRuleParameters(params IApplyRule) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateCheckResourceParameters(node awscdk.CfnResource) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateValidateParameters(context awscdk.IPolicyValidationContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateValidateScopeParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PCIDSS321Checks) validateSetPackNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPCIDSS321ChecksParameters(props *NagPackProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

