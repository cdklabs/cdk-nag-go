//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func validateNagRules_ResolveIfPrimitiveParameters(node awscdk.CfnResource, parameter interface{}) error {
	return nil
}

func validateNagRules_ResolveResourceFromInstrinsicParameters(node awscdk.CfnResource, parameter interface{}) error {
	return nil
}

func validateNagRules_ResolveResourceFromIntrinsicParameters(node awscdk.CfnResource, parameter interface{}) error {
	return nil
}

