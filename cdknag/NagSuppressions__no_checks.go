//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func validateNagSuppressions_AddResourceSuppressionsParameters(construct interface{}, suppressions *[]*NagPackSuppression) error {
	return nil
}

func validateNagSuppressions_AddResourceSuppressionsByPathParameters(stack awscdk.Stack, path interface{}, suppressions *[]*NagPackSuppression) error {
	return nil
}

func validateNagSuppressions_AddStackSuppressionsParameters(stack awscdk.Stack, suppressions *[]*NagPackSuppression) error {
	return nil
}

