//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (n *jsiiProxy_NagPack) validateApplyRuleParameters(params IApplyRule) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NagPack) validateCreateComplianceReportLineParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
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

func (n *jsiiProxy_NagPack) validateCreateMessageParameters(ruleId *string, findingId *string, info *string, explanation *string) error {
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

func (n *jsiiProxy_NagPack) validateIgnoreRuleParameters(ignores *[]*NagPackSuppression, ruleId *string, findingId *string) error {
	if ignores == nil {
		return fmt.Errorf("parameter ignores is required, but nil was provided")
	}
	for idx_045ee9, v := range *ignores {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter ignores[%#v]", idx_045ee9) }); err != nil {
			return err
		}
	}

	if ruleId == nil {
		return fmt.Errorf("parameter ruleId is required, but nil was provided")
	}

	if findingId == nil {
		return fmt.Errorf("parameter findingId is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NagPack) validateInitializeStackReportParameters(params IApplyRule) error {
	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NagPack) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NagPack) validateWriteToStackComplianceReportParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
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

func (j *jsiiProxy_NagPack) validateSetLogIgnoresParameters(val *bool) error {
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

func (j *jsiiProxy_NagPack) validateSetReportsParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NagPack) validateSetReportStacksParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_NagPack) validateSetVerboseParameters(val *bool) error {
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

