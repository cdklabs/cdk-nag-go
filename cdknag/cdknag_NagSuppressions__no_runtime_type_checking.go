//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func validateNagSuppressions_AddResourceSuppressionsParameters(construct constructs.IConstruct, suppressions *[]*NagPackSuppression) error {
	return nil
}

func validateNagSuppressions_AddResourceSuppressionsByPathParameters(stack awscdk.Stack, path *string, suppressions *[]*NagPackSuppression) error {
	return nil
}

func validateNagSuppressions_AddStackSuppressionsParameters(stack awscdk.Stack, suppressions *[]*NagPackSuppression) error {
	return nil
}

