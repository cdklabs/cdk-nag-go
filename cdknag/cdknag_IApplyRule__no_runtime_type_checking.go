//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IApplyRule) validateRuleParameters(node awscdk.CfnResource) error {
	return nil
}

func (j *jsiiProxy_IApplyRule) validateSetExplanationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IApplyRule) validateSetInfoParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IApplyRule) validateSetLevelParameters(val NagMessageLevel) error {
	return nil
}

func (j *jsiiProxy_IApplyRule) validateSetNodeParameters(val awscdk.CfnResource) error {
	return nil
}

