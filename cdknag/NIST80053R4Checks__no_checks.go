//go:build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NIST80053R4Checks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (n *jsiiProxy_NIST80053R4Checks) validateCreateComplianceReportLineParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (n *jsiiProxy_NIST80053R4Checks) validateCreateMessageParameters(ruleId *string, findingId *string, info *string, explanation *string) error {
	return nil
}

func (n *jsiiProxy_NIST80053R4Checks) validateIgnoreRuleParameters(ignores *[]*NagPackSuppression, ruleId *string, findingId *string) error {
	return nil
}

func (n *jsiiProxy_NIST80053R4Checks) validateInitializeStackReportParameters(params IApplyRule) error {
	return nil
}

func (n *jsiiProxy_NIST80053R4Checks) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (n *jsiiProxy_NIST80053R4Checks) validateWriteToStackComplianceReportParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (j *jsiiProxy_NIST80053R4Checks) validateSetLogIgnoresParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_NIST80053R4Checks) validateSetPackNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NIST80053R4Checks) validateSetReportsParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_NIST80053R4Checks) validateSetReportStacksParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_NIST80053R4Checks) validateSetVerboseParameters(val *bool) error {
	return nil
}

func validateNewNIST80053R4ChecksParameters(props *NagPackProps) error {
	return nil
}

