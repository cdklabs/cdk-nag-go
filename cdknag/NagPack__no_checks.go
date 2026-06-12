//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NagPack) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (n *jsiiProxy_NagPack) validateCheckResourceParameters(node awscdk.CfnResource) error {
	return nil
}

func (n *jsiiProxy_NagPack) validateValidateParameters(context awscdk.IPolicyValidationContext) error {
	return nil
}

func (n *jsiiProxy_NagPack) validateValidateScopeParameters(scope constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_NagPack) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewNagPackParameters(props *NagPackProps) error {
	return nil
}

