//go:build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SuppressionIgnoreAlways) validateCreateMessageParameters(_input *SuppressionIgnoreInput) error {
	return nil
}

func validateNewSuppressionIgnoreAlwaysParameters(triggerMessage *string) error {
	return nil
}

