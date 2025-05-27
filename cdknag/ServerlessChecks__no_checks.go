//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServerlessChecks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (s *jsiiProxy_ServerlessChecks) validateIgnoreRuleParameters(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel) error {
	return nil
}

func (s *jsiiProxy_ServerlessChecks) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_ServerlessChecks) validateSetLoggersParameters(val *[]INagLogger) error {
	return nil
}

func (j *jsiiProxy_ServerlessChecks) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewServerlessChecksParameters(props *NagPackProps) error {
	return nil
}

