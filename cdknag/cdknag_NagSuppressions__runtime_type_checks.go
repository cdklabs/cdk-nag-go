//go:build !no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func validateNagSuppressions_AddResourceSuppressionsParameters(construct interface{}, suppressions *[]*NagPackSuppression) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}
	switch construct.(type) {
	case constructs.IConstruct:
		// ok
	case *[]constructs.IConstruct:
		// ok
	case []constructs.IConstruct:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(construct) {
			return fmt.Errorf("parameter construct must be one of the allowed types: constructs.IConstruct, *[]constructs.IConstruct; received %#v (a %T)", construct, construct)
		}
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

func validateNagSuppressions_AddResourceSuppressionsByPathParameters(stack awscdk.Stack, path interface{}, suppressions *[]*NagPackSuppression) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}
	switch path.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *[]*string:
		// ok
	case []*string:
		// ok
	default:
		return fmt.Errorf("parameter path must be one of the allowed types: *string, *[]*string; received %#v (a %T)", path, path)
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

