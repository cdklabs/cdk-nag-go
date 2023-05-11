//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AnnotationLogger) validateCreateMessageParameters(ruleId *string, findingId *string, ruleInfo *string, ruleExplanation *string, verbose *bool) error {
	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnComplianceParameters(_data *NagLoggerComplianceData) error {
	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnErrorParameters(data *NagLoggerErrorData) error {
	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnNonComplianceParameters(data *NagLoggerNonComplianceData) error {
	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnNotApplicableParameters(_data *NagLoggerNotApplicableData) error {
	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnSuppressedParameters(data *NagLoggerSuppressedData) error {
	return nil
}

func (a *jsiiProxy_AnnotationLogger) validateOnSuppressedErrorParameters(data *NagLoggerSuppressedErrorData) error {
	return nil
}

func (j *jsiiProxy_AnnotationLogger) validateSetSuppressionIdParameters(val *string) error {
	return nil
}

func validateNewAnnotationLoggerParameters(props *AnnotationLoggerProps) error {
	return nil
}

