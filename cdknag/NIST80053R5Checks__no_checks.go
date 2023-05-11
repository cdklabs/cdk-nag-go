//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NIST80053R5Checks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (n *jsiiProxy_NIST80053R5Checks) validateIgnoreRuleParameters(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel) error {
	return nil
}

func (n *jsiiProxy_NIST80053R5Checks) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_NIST80053R5Checks) validateSetLoggersParameters(val *[]INagLogger) error {
	return nil
}

func (j *jsiiProxy_NIST80053R5Checks) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewNIST80053R5ChecksParameters(props *NagPackProps) error {
	return nil
}

