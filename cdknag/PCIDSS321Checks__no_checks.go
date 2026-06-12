//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PCIDSS321Checks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateCheckResourceParameters(node awscdk.CfnResource) error {
	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateValidateParameters(context awscdk.IPolicyValidationContext) error {
	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateValidateScopeParameters(scope constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_PCIDSS321Checks) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewPCIDSS321ChecksParameters(props *NagPackProps) error {
	return nil
}

