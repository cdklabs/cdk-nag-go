//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NIST80053R4Checks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (n *jsiiProxy_NIST80053R4Checks) validateCheckResourceParameters(node awscdk.CfnResource) error {
	return nil
}

func (n *jsiiProxy_NIST80053R4Checks) validateValidateParameters(context awscdk.IPolicyValidationContext) error {
	return nil
}

func (n *jsiiProxy_NIST80053R4Checks) validateValidateScopeParameters(scope constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_NIST80053R4Checks) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewNIST80053R4ChecksParameters(props *NagPackProps) error {
	return nil
}

