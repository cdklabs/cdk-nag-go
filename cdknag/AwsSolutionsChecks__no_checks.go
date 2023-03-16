//go:build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsSolutionsChecks) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateCreateComplianceReportLineParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateCreateMessageParameters(ruleId *string, findingId *string, info *string, explanation *string) error {
	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateIgnoreRuleParameters(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel) error {
	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateInitializeStackReportParameters(params IApplyRule) error {
	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (a *jsiiProxy_AwsSolutionsChecks) validateWriteToStackComplianceReportParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetLogIgnoresParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetPackNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetReportsParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetReportStacksParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_AwsSolutionsChecks) validateSetVerboseParameters(val *bool) error {
	return nil
}

func validateNewAwsSolutionsChecksParameters(props *NagPackProps) error {
	return nil
}

