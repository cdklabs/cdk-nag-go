//go:build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NIST80053R5Checks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (n *jsiiProxy_NIST80053R5Checks) validateCreateComplianceReportLineParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (n *jsiiProxy_NIST80053R5Checks) validateCreateMessageParameters(ruleId *string, findingId *string, info *string, explanation *string) error {
	return nil
}

func (n *jsiiProxy_NIST80053R5Checks) validateIgnoreRuleParameters(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel) error {
	return nil
}

func (n *jsiiProxy_NIST80053R5Checks) validateInitializeStackReportParameters(params IApplyRule) error {
	return nil
}

func (n *jsiiProxy_NIST80053R5Checks) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (n *jsiiProxy_NIST80053R5Checks) validateWriteToStackComplianceReportParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (j *jsiiProxy_NIST80053R5Checks) validateSetLogIgnoresParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_NIST80053R5Checks) validateSetPackNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NIST80053R5Checks) validateSetReportsParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_NIST80053R5Checks) validateSetReportStacksParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_NIST80053R5Checks) validateSetVerboseParameters(val *bool) error {
	return nil
}

func validateNewNIST80053R5ChecksParameters(props *NagPackProps) error {
	return nil
}

