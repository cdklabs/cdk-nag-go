//go:build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HIPAASecurityChecks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (h *jsiiProxy_HIPAASecurityChecks) validateCreateComplianceReportLineParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (h *jsiiProxy_HIPAASecurityChecks) validateCreateMessageParameters(ruleId *string, findingId *string, info *string, explanation *string) error {
	return nil
}

func (h *jsiiProxy_HIPAASecurityChecks) validateIgnoreRuleParameters(ignores *[]*NagPackSuppression, ruleId *string, findingId *string) error {
	return nil
}

func (h *jsiiProxy_HIPAASecurityChecks) validateInitializeStackReportParameters(params IApplyRule) error {
	return nil
}

func (h *jsiiProxy_HIPAASecurityChecks) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (h *jsiiProxy_HIPAASecurityChecks) validateWriteToStackComplianceReportParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (j *jsiiProxy_HIPAASecurityChecks) validateSetLogIgnoresParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_HIPAASecurityChecks) validateSetPackNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HIPAASecurityChecks) validateSetReportsParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_HIPAASecurityChecks) validateSetReportStacksParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_HIPAASecurityChecks) validateSetVerboseParameters(val *bool) error {
	return nil
}

func validateNewHIPAASecurityChecksParameters(props *NagPackProps) error {
	return nil
}

