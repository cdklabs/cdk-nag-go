//go:build !no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_INagLogger) validateOnComplianceParameters(data *NagLoggerComplianceData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_INagLogger) validateOnErrorParameters(data *NagLoggerErrorData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_INagLogger) validateOnNonComplianceParameters(data *NagLoggerNonComplianceData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_INagLogger) validateOnNotApplicableParameters(data *NagLoggerNotApplicableData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_INagLogger) validateOnSuppressedParameters(data *NagLoggerSuppressedData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_INagLogger) validateOnSuppressedErrorParameters(data *NagLoggerSuppressedErrorData) error {
	if data == nil {
		return fmt.Errorf("parameter data is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(data, func() string { return "parameter data" }); err != nil {
		return err
	}

	return nil
}

