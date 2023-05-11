//go:build !no_runtime_type_checking

package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (n *jsiiProxy_NagReportLogger) validateGetFormatStacksParameters(format NagReportFormat) error {
	if format == "" {
		return fmt.Errorf("parameter format is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NagReportLogger) validateInitializeStackReportParameters(data *NagLoggerBaseData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnComplianceParameters(data *NagLoggerComplianceData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnErrorParameters(data *NagLoggerErrorData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnNonComplianceParameters(data *NagLoggerNonComplianceData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnNotApplicableParameters(data *NagLoggerNotApplicableData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnSuppressedParameters(data *NagLoggerSuppressedData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnSuppressedErrorParameters(data *NagLoggerSuppressedErrorData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NagReportLogger) validateWriteToStackComplianceReportParameters(data *NagLoggerBaseData, compliance interface{}) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	if compliance == nil {
		return fmt.Errorf("parameter compliance is required, but nil was provided")
	}
	switch compliance.(type) {
	case NagRuleCompliance:
		// ok
	case NagRulePostValidationStates:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(compliance) {
			return fmt.Errorf("parameter compliance must be one of the allowed types: NagRuleCompliance, NagRulePostValidationStates; received %#v (a %T)", compliance, compliance)
		}
	}

	return nil
}

func validateNewNagReportLoggerParameters(props *NagReportLoggerProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

