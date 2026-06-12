//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsSolutionsChecks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateCheckResourceParameters(node awscdk.CfnResource) error {
	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateValidateParameters(context awscdk.IPolicyValidationContext) error {
	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateValidateScopeParameters(scope constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewAwsSolutionsChecksParameters(props *NagPackProps) error {
	return nil
}

