//go:build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_INagLogger) validateOnComplianceParameters(data *NagLoggerComplianceData) error {
	return nil
}

func (i *jsiiProxy_INagLogger) validateOnErrorParameters(data *NagLoggerErrorData) error {
	return nil
}

func (i *jsiiProxy_INagLogger) validateOnNonComplianceParameters(data *NagLoggerNonComplianceData) error {
	return nil
}

func (i *jsiiProxy_INagLogger) validateOnNotApplicableParameters(data *NagLoggerNotApplicableData) error {
	return nil
}

func (i *jsiiProxy_INagLogger) validateOnSuppressedParameters(data *NagLoggerSuppressedData) error {
	return nil
}

func (i *jsiiProxy_INagLogger) validateOnSuppressedErrorParameters(data *NagLoggerSuppressedErrorData) error {
	return nil
}

