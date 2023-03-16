//go:build !no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AwsSolutionsChecks) validateApplyRuleParameters(params IApplyRule) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateCreateComplianceReportLineParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}

	if ruleId == nil {
		return fmt.Errorf("parameter ruleId is required, but nil was provided")
	}

	if compliance == nil {
		return fmt.Errorf("parameter compliance is required, but nil was provided")
	}
	switch compliance.(type) {
	case NagRuleCompliance:
		// ok
	case *string:
		// ok
	case string:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(compliance) {
			return fmt.Errorf("parameter compliance must be one of the allowed types: NagRuleCompliance, *string; received %#v (a %T)", compliance, compliance)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateCreateMessageParameters(ruleId *string, findingId *string, info *string, explanation *string) error {
	if ruleId == nil {
		return fmt.Errorf("parameter ruleId is required, but nil was provided")
	}

	if findingId == nil {
		return fmt.Errorf("parameter findingId is required, but nil was provided")
	}

	if info == nil {
		return fmt.Errorf("parameter info is required, but nil was provided")
	}

	if explanation == nil {
		return fmt.Errorf("parameter explanation is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateIgnoreRuleParameters(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel) error {
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

func (a *jsiiProxy_AwsSolutionsChecks) validateInitializeStackReportParameters(params IApplyRule) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateWriteToStackComplianceReportParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}

	if ruleId == nil {
		return fmt.Errorf("parameter ruleId is required, but nil was provided")
	}

	if compliance == nil {
		return fmt.Errorf("parameter compliance is required, but nil was provided")
	}
	switch compliance.(type) {
	case NagRuleCompliance:
		// ok
	case *string:
		// ok
	case string:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(compliance) {
			return fmt.Errorf("parameter compliance must be one of the allowed types: NagRuleCompliance, *string; received %#v (a %T)", compliance, compliance)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetLogIgnoresParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetPackNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetReportsParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetReportStacksParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetVerboseParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsSolutionsChecksParameters(props *NagPackProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

