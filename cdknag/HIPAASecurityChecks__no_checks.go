//go:build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HIPAASecurityChecks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (h *jsiiProxy_HIPAASecurityChecks) validateIgnoreRuleParameters(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel) error {
	return nil
}

func (h *jsiiProxy_HIPAASecurityChecks) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_HIPAASecurityChecks) validateSetLoggersParameters(val *[]INagLogger) error {
	return nil
}

func (j *jsiiProxy_HIPAASecurityChecks) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewHIPAASecurityChecksParameters(props *NagPackProps) error {
	return nil
}

