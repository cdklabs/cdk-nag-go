//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func validateNagSuppressions_AddResourceSuppressionsParameters(construct constructs.IConstruct, suppressions *[]*NagPackSuppression) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	if suppressions == nil {
		return fmt.Errorf("parameter suppressions is required, but nil was provided")
	}
	for idx_1cb3ae, v := range *suppressions {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter suppressions[%#v]", idx_1cb3ae) }); err != nil {
			return err
		}
	}

	return nil
}

func validateNagSuppressions_AddResourceSuppressionsByPathParameters(stack awscdk.Stack, path *string, suppressions *[]*NagPackSuppression) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if suppressions == nil {
		return fmt.Errorf("parameter suppressions is required, but nil was provided")
	}
	for idx_1cb3ae, v := range *suppressions {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter suppressions[%#v]", idx_1cb3ae) }); err != nil {
			return err
		}
	}

	return nil
}

func validateNagSuppressions_AddStackSuppressionsParameters(stack awscdk.Stack, suppressions *[]*NagPackSuppression) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if suppressions == nil {
		return fmt.Errorf("parameter suppressions is required, but nil was provided")
	}
	for idx_1cb3ae, v := range *suppressions {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter suppressions[%#v]", idx_1cb3ae) }); err != nil {
			return err
		}
	}

	return nil
}
