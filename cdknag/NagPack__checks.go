//go:build !no_runtime_type_checking

package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (n *jsiiProxy_NagPack) validateApplyRuleParameters(params IApplyRule) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NagPack) validateIgnoreRuleParameters(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel) error {
	if suppressions == nil {
		return fmt.Errorf("parameter suppressions is required, but nil was provided")
	}
	for idx_1cb3ae, v := range *suppressions {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter suppressions[%#v]", idx_1cb3ae) }); err != nil {
			return err
		}
	}

	if ruleId == nil {
		return fmt.Errorf("parameter ruleId is required, but nil was provided")
	}

	if findingId == nil {
		return fmt.Errorf("parameter findingId is required, but nil was provided")
	}

	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	if level == "" {
		return fmt.Errorf("parameter level is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NagPack) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NagPack) validateSetLoggersParameters(val *[]INagLogger) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NagPack) validateSetPackNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewNagPackParameters(props *NagPackProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

