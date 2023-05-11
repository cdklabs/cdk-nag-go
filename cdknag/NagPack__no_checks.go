//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NagPack) validateApplyRuleParameters(params IApplyRule) error {
	return nil
}

func (n *jsiiProxy_NagPack) validateIgnoreRuleParameters(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel) error {
	return nil
}

func (n *jsiiProxy_NagPack) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_NagPack) validateSetLoggersParameters(val *[]INagLogger) error {
	return nil
}

func (j *jsiiProxy_NagPack) validateSetPackNameParameters(val *string) error {
	return nil
}

func validateNewNagPackParameters(props *NagPackProps) error {
	return nil
}

