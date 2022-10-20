//go:build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PCIDSS321Checks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateCreateComplianceReportLineParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateCreateMessageParameters(ruleId *string, findingId *string, info *string, explanation *string) error {
	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateIgnoreRuleParameters(ignores *[]*NagPackSuppression, ruleId *string, findingId *string) error {
	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateInitializeStackReportParameters(params IApplyRule) error {
	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (p *jsiiProxy_PCIDSS321Checks) validateWriteToStackComplianceReportParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (j *jsiiProxy_PCIDSS321Checks) validateSetLogIgnoresParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_PCIDSS321Checks) validateSetPackNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PCIDSS321Checks) validateSetReportsParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_PCIDSS321Checks) validateSetReportStacksParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_PCIDSS321Checks) validateSetVerboseParameters(val *bool) error {
	return nil
}

func validateNewPCIDSS321ChecksParameters(props *NagPackProps) error {
	return nil
}

