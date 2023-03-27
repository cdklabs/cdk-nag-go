//go:build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PCIDSS321Checks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateIgnoreRuleParameters(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel) error {
	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_PCIDSS321Checks) validateSetLoggersParameters(val *[]INagLogger) error {
	return nil
}

func (j *jsiiProxy_PCIDSS321Checks) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewPCIDSS321ChecksParameters(props *NagPackProps) error {
	return nil
}

