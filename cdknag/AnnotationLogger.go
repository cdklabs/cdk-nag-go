// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"
)

// A NagLogger that outputs to the CDK Annotations system.
type AnnotationLogger interface {
	INagLogger
	LogIgnores() *bool
	SuppressionId() *string
	SetSuppressionId(val *string)
	Verbose() *bool
	CreateMessage(ruleId *string, findingId *string, ruleInfo *string, ruleExplanation *string, verbose *bool) *string
	// Called when a CfnResource passes the compliance check for a given rule.
	OnCompliance(_data *NagLoggerComplianceData)
	// Called when a rule throws an error during while validating a CfnResource for compliance.
	OnError(data *NagLoggerErrorData)
	// Called when a CfnResource does not pass the compliance check for a given rule and the the rule violation is not suppressed by the user.
	OnNonCompliance(data *NagLoggerNonComplianceData)
	// Called when a rule does not apply to the given CfnResource.
	OnNotApplicable(_data *NagLoggerNotApplicableData)
	// Called when a CfnResource does not pass the compliance check for a given rule and the rule violation is suppressed by the user.
	OnSuppressed(data *NagLoggerSuppressedData)
	// Called when a rule throws an error during while validating a CfnResource for compliance and the error is suppressed.
	OnSuppressedError(data *NagLoggerSuppressedErrorData)
}

// The jsii proxy struct for AnnotationLogger
type jsiiProxy_AnnotationLogger struct {
	jsiiProxy_INagLogger
}

func (j *jsiiProxy_AnnotationLogger) LogIgnores() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logIgnores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnnotationLogger) SuppressionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suppressionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnnotationLogger) Verbose() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"verbose",
		&returns,
	)
	return returns
}


func NewAnnotationLogger(props *AnnotationLoggerProps) AnnotationLogger {
	_init_.Initialize()

	if err := validateNewAnnotationLoggerParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AnnotationLogger{}

	_jsii_.Create(
		"cdk-nag.AnnotationLogger",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAnnotationLogger_Override(a AnnotationLogger, props *AnnotationLoggerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.AnnotationLogger",
		[]interface{}{props},
		a,
	)
}

func (j *jsiiProxy_AnnotationLogger)SetSuppressionId(val *string) {
	if err := j.validateSetSuppressionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suppressionId",
		val,
	)
}

func (a *jsiiProxy_AnnotationLogger) CreateMessage(ruleId *string, findingId *string, ruleInfo *string, ruleExplanation *string, verbose *bool) *string {
	if err := a.validateCreateMessageParameters(ruleId, findingId, ruleInfo, ruleExplanation, verbose); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"createMessage",
		[]interface{}{ruleId, findingId, ruleInfo, ruleExplanation, verbose},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AnnotationLogger) OnCompliance(_data *NagLoggerComplianceData) {
	if err := a.validateOnComplianceParameters(_data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onCompliance",
		[]interface{}{_data},
	)
}

func (a *jsiiProxy_AnnotationLogger) OnError(data *NagLoggerErrorData) {
	if err := a.validateOnErrorParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onError",
		[]interface{}{data},
	)
}

func (a *jsiiProxy_AnnotationLogger) OnNonCompliance(data *NagLoggerNonComplianceData) {
	if err := a.validateOnNonComplianceParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onNonCompliance",
		[]interface{}{data},
	)
}

func (a *jsiiProxy_AnnotationLogger) OnNotApplicable(_data *NagLoggerNotApplicableData) {
	if err := a.validateOnNotApplicableParameters(_data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onNotApplicable",
		[]interface{}{_data},
	)
}

func (a *jsiiProxy_AnnotationLogger) OnSuppressed(data *NagLoggerSuppressedData) {
	if err := a.validateOnSuppressedParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSuppressed",
		[]interface{}{data},
	)
}

func (a *jsiiProxy_AnnotationLogger) OnSuppressedError(data *NagLoggerSuppressedErrorData) {
	if err := a.validateOnSuppressedErrorParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSuppressedError",
		[]interface{}{data},
	)
}

