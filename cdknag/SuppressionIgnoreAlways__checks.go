//go:build !no_runtime_type_checking

package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_SuppressionIgnoreAlways) validateCreateMessageParameters(_input *SuppressionIgnoreInput) error {
	if _input == nil {
		return fmt.Errorf("parameter _input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_input, func() string { return "parameter _input" }); err != nil {
		return err
	}

	return nil
}

func validateNewSuppressionIgnoreAlwaysParameters(triggerMessage *string) error {
	if triggerMessage == nil {
		return fmt.Errorf("parameter triggerMessage is required, but nil was provided")
	}

	return nil
}

