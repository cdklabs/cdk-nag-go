//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HIPAASecurityChecks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (h *jsiiProxy_HIPAASecurityChecks) validateCheckResourceParameters(node awscdk.CfnResource) error {
	return nil
}

func (h *jsiiProxy_HIPAASecurityChecks) validateValidateParameters(context awscdk.IPolicyValidationContext) error {
	return nil
}

func (h *jsiiProxy_HIPAASecurityChecks) validateValidateScopeParameters(scope constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_HIPAASecurityChecks) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewHIPAASecurityChecksParameters(props *NagPackProps) error {
	return nil
}

