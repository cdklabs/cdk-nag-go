// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for creating NagSuppression Ignores.
type INagLogger interface {
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
}

// The jsii proxy for INagLogger
type jsiiProxy_INagLogger struct {
	_ byte // padding
}

func (i *jsiiProxy_INagLogger) OnCompliance(data *NagLoggerComplianceData) {
	if err := i.validateOnComplianceParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onCompliance",
		[]interface{}{data},
	)
}

func (i *jsiiProxy_INagLogger) OnError(data *NagLoggerErrorData) {
	if err := i.validateOnErrorParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onError",
		[]interface{}{data},
	)
}

func (i *jsiiProxy_INagLogger) OnNonCompliance(data *NagLoggerNonComplianceData) {
	if err := i.validateOnNonComplianceParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onNonCompliance",
		[]interface{}{data},
	)
}

func (i *jsiiProxy_INagLogger) OnNotApplicable(data *NagLoggerNotApplicableData) {
	if err := i.validateOnNotApplicableParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onNotApplicable",
		[]interface{}{data},
	)
}

func (i *jsiiProxy_INagLogger) OnSuppressed(data *NagLoggerSuppressedData) {
	if err := i.validateOnSuppressedParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onSuppressed",
		[]interface{}{data},
	)
}

func (i *jsiiProxy_INagLogger) OnSuppressedError(data *NagLoggerSuppressedErrorData) {
	if err := i.validateOnSuppressedErrorParameters(data); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"onSuppressedError",
		[]interface{}{data},
	)
}

