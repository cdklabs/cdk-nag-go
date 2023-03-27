//go:build no_runtime_type_checking

// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NagReportLogger) validateGetFormatStacksParameters(format NagReportFormat) error {
	return nil
}

func (n *jsiiProxy_NagReportLogger) validateInitializeStackReportParameters(data *NagLoggerBaseData) error {
	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnComplianceParameters(data *NagLoggerComplianceData) error {
	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnErrorParameters(data *NagLoggerErrorData) error {
	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnNonComplianceParameters(data *NagLoggerNonComplianceData) error {
	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnNotApplicableParameters(data *NagLoggerNotApplicableData) error {
	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnSuppressedParameters(data *NagLoggerSuppressedData) error {
	return nil
}

func (n *jsiiProxy_NagReportLogger) validateOnSuppressedErrorParameters(data *NagLoggerSuppressedErrorData) error {
	return nil
}

func (n *jsiiProxy_NagReportLogger) validateWriteToStackComplianceReportParameters(data *NagLoggerBaseData, compliance interface{}) error {
	return nil
}

func validateNewNagReportLoggerParameters(props *NagReportLoggerProps) error {
	return nil
}

