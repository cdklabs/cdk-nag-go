//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServerlessChecks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (s *jsiiProxy_ServerlessChecks) validateCheckResourceParameters(node awscdk.CfnResource) error {
	return nil
}

func (s *jsiiProxy_ServerlessChecks) validateValidateParameters(context awscdk.IPolicyValidationContext) error {
	return nil
}

func (s *jsiiProxy_ServerlessChecks) validateValidateScopeParameters(scope constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_ServerlessChecks) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewServerlessChecksParameters(props *NagPackProps) error {
	return nil
}

