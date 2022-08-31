//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func validateNagRules_ResolveIfPrimitiveParameters(node awscdk.CfnResource, parameter interface{}) error {
	return nil
}

func validateNagRules_ResolveResourceFromInstrinsicParameters(node awscdk.CfnResource, parameter interface{}) error {
	return nil
}

