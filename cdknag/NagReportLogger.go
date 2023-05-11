package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"
)

// A NagLogger that creates compliance reports.
type NagReportLogger interface {
	INagLogger
	Formats() *[]NagReportFormat
	GetFormatStacks(format NagReportFormat) *[]*string
	// Initialize the report for the rule pack's compliance report for the resource's Stack if it doesn't exist.
	InitializeStackReport(data *NagLoggerBaseData)
	// Called when a CfnResource passes the compliance check for a given rule.
	OnCompliance(data *NagLoggerComplianceData)
	// Called when a rule throws an error during while validating a CfnResource for compliance.
	OnError(data *NagLoggerErrorData)
	// Called when a CfnResource does not pass the compliance check for a given rule and the the rule violation is not suppressed by the user.
	OnNonCompliance(data *NagLoggerNonComplianceData)
	// Called when a rule does not apply to the given CfnResource.
	OnNotApplicable(data *NagLoggerNotApplicableData)
	// Called when a CfnResource does not pass the compliance check for a given rule and the rule violation is suppressed by the user.
	OnSuppressed(data *NagLoggerSuppressedData)
	// Called when a rule throws an error during while validating a CfnResource for compliance and the error is suppressed.
	OnSuppressedError(data *NagLoggerSuppressedErrorData)
	WriteToStackComplianceReport(data *NagLoggerBaseData, compliance interface{})
}

// The jsii proxy struct for NagReportLogger
type jsiiProxy_NagReportLogger struct {
	jsiiProxy_INagLogger
}

func (j *jsiiProxy_NagReportLogger) Formats() *[]NagReportFormat {
	var returns *[]NagReportFormat
	_jsii_.Get(
		j,
		"formats",
		&returns,
	)
	return returns
}


func NewNagReportLogger(props *NagReportLoggerProps) NagReportLogger {
	_init_.Initialize()

	if err := validateNewNagReportLoggerParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NagReportLogger{}

	_jsii_.Create(
		"cdk-nag.NagReportLogger",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewNagReportLogger_Override(n NagReportLogger, props *NagReportLoggerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.NagReportLogger",
		[]interface{}{props},
		n,
	)
}

func (n *jsiiProxy_NagReportLogger) GetFormatStacks(format NagReportFormat) *[]*string {
	if err := n.validateGetFormatStacksParameters(format); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getFormatStacks",
		[]interface{}{format},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NagReportLogger) InitializeStackReport(data *NagLoggerBaseData) {
	if err := n.validateInitializeStackReportParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"initializeStackReport",
		[]interface{}{data},
	)
}

func (n *jsiiProxy_NagReportLogger) OnCompliance(data *NagLoggerComplianceData) {
	if err := n.validateOnComplianceParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"onCompliance",
		[]interface{}{data},
	)
}

func (n *jsiiProxy_NagReportLogger) OnError(data *NagLoggerErrorData) {
	if err := n.validateOnErrorParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"onError",
		[]interface{}{data},
	)
}

func (n *jsiiProxy_NagReportLogger) OnNonCompliance(data *NagLoggerNonComplianceData) {
	if err := n.validateOnNonComplianceParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"onNonCompliance",
		[]interface{}{data},
	)
}

func (n *jsiiProxy_NagReportLogger) OnNotApplicable(data *NagLoggerNotApplicableData) {
	if err := n.validateOnNotApplicableParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"onNotApplicable",
		[]interface{}{data},
	)
}

func (n *jsiiProxy_NagReportLogger) OnSuppressed(data *NagLoggerSuppressedData) {
	if err := n.validateOnSuppressedParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"onSuppressed",
		[]interface{}{data},
	)
}

func (n *jsiiProxy_NagReportLogger) OnSuppressedError(data *NagLoggerSuppressedErrorData) {
	if err := n.validateOnSuppressedErrorParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"onSuppressedError",
		[]interface{}{data},
	)
}

func (n *jsiiProxy_NagReportLogger) WriteToStackComplianceReport(data *NagLoggerBaseData, compliance interface{}) {
	if err := n.validateWriteToStackComplianceReportParameters(data, compliance); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"writeToStackComplianceReport",
		[]interface{}{data, compliance},
	)
}

