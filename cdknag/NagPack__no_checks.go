//go:build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NagPack) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (n *jsiiProxy_NagPack) validateCreateComplianceReportLineParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (n *jsiiProxy_NagPack) validateCreateMessageParameters(ruleId *string, findingId *string, info *string, explanation *string) error {
	return nil
}

func (n *jsiiProxy_NagPack) validateIgnoreRuleParameters(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel) error {
	return nil
}

func (n *jsiiProxy_NagPack) validateInitializeStackReportParameters(params IApplyRule) error {
	return nil
}

func (n *jsiiProxy_NagPack) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (n *jsiiProxy_NagPack) validateWriteToStackComplianceReportParameters(params IApplyRule, ruleId *string, compliance interface{}) error {
	return nil
}

func (j *jsiiProxy_NagPack) validateSetLogIgnoresParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_NagPack) validateSetPackNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NagPack) validateSetReportsParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_NagPack) validateSetReportStacksParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_NagPack) validateSetVerboseParameters(val *bool) error {
	return nil
}

func validateNewNagPackParameters(props *NagPackProps) error {
	return nil
}

