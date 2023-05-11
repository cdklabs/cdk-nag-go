//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsSolutionsChecks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateIgnoreRuleParameters(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel) error {
	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetLoggersParameters(val *[]INagLogger) error {
	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewAwsSolutionsChecksParameters(props *NagPackProps) error {
	return nil
}

