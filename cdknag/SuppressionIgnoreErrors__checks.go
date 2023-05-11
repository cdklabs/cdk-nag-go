//go:build !no_runtime_type_checking

package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_SuppressionIgnoreErrors) validateCreateMessageParameters(input *SuppressionIgnoreInput) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

