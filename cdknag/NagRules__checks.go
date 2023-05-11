//go:build !no_runtime_type_checking

package cdknag

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func validateNagRules_ResolveIfPrimitiveParameters(node awscdk.CfnResource, parameter interface{}) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	return nil
}

func validateNagRules_ResolveResourceFromInstrinsicParameters(node awscdk.CfnResource, parameter interface{}) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	if parameter == nil {
		return fmt.Errorf("parameter parameter is required, but nil was provided")
	}

	return nil
}

