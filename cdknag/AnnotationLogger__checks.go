//go:build !no_runtime_type_checking

package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (a *jsiiProxy_AnnotationLogger) validateCreateMessageParameters(ruleId *string, findingId *string, ruleInfo *string, ruleExplanation *string, verbose *bool) error {
	if ruleId == nil {
		return fmt.Errorf("parameter ruleId is required, but nil was provided")
	}

	if findingId == nil {
		return fmt.Errorf("parameter findingId is required, but nil was provided")
	}

	if ruleInfo == nil {
		return fmt.Errorf("parameter ruleInfo is required, but nil was provided")
	}

	if ruleExplanation == nil {
		return fmt.Errorf("parameter ruleExplanation is required, but nil was provided")
	}

	if verbose == nil {
		return fmt.Errorf("parameter verbose is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnComplianceParameters(_data *NagLoggerComplianceData) error {
	if _data == nil {
		return fmt.Errorf("parameter _data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_data, func() string { return "parameter _data" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnErrorParameters(data *NagLoggerErrorData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnNonComplianceParameters(data *NagLoggerNonComplianceData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnNotApplicableParameters(_data *NagLoggerNotApplicableData) error {
	if _data == nil {
		return fmt.Errorf("parameter _data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_data, func() string { return "parameter _data" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnSuppressedParameters(data *NagLoggerSuppressedData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnSuppressedErrorParameters(data *NagLoggerSuppressedErrorData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_AnnotationLogger) validateSetSuppressionIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAnnotationLoggerParameters(props *AnnotationLoggerProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

