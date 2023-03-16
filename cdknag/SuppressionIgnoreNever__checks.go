//go:build !no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (s *jsiiProxy_SuppressionIgnoreNever) validateCreateMessageParameters(_input *SuppressionIgnoreInput) error {
	if _input == nil {
		return fmt.Errorf("parameter _input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_input, func() string { return "parameter _input" }); err != nil {
		return err
	}

	return nil
}

