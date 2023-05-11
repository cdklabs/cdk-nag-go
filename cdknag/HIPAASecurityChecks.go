package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Check for HIPAA Security compliance.
//
// Based on the HIPAA Security AWS operational best practices: https://docs.aws.amazon.com/config/latest/developerguide/operational-best-practices-for-hipaa_security.html
type HIPAASecurityChecks interface {
	NagPack
	Loggers() *[]INagLogger
	SetLoggers(val *[]INagLogger)
	PackGlobalSuppressionIgnore() INagSuppressionIgnore
	SetPackGlobalSuppressionIgnore(val INagSuppressionIgnore)
	PackName() *string
	SetPackName(val *string)
	ReadPackName() *string
	UserGlobalSuppressionIgnore() INagSuppressionIgnore
	SetUserGlobalSuppressionIgnore(val INagSuppressionIgnore)
	// Create a rule to be used in the NagPack.
	ApplyRule(params IApplyRule)
	// Check whether a specific rule should be ignored.
	//
	// Returns: The reason the rule was ignored, or an empty string.
	IgnoreRule(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel, ignoreSuppressionCondition INagSuppressionIgnore) *string
	// All aspects can visit an IConstruct.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for HIPAASecurityChecks
type jsiiProxy_HIPAASecurityChecks struct {
	jsiiProxy_NagPack
}

func (j *jsiiProxy_HIPAASecurityChecks) Loggers() *[]INagLogger {
	var returns *[]INagLogger
	_jsii_.Get(
		j,
		"loggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HIPAASecurityChecks) PackGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"packGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HIPAASecurityChecks) PackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HIPAASecurityChecks) ReadPackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readPackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HIPAASecurityChecks) UserGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"userGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}


func NewHIPAASecurityChecks(props *NagPackProps) HIPAASecurityChecks {
	_init_.Initialize()

	if err := validateNewHIPAASecurityChecksParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HIPAASecurityChecks{}

	_jsii_.Create(
		"cdk-nag.HIPAASecurityChecks",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewHIPAASecurityChecks_Override(h HIPAASecurityChecks, props *NagPackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.HIPAASecurityChecks",
		[]interface{}{props},
		h,
	)
}

func (j *jsiiProxy_HIPAASecurityChecks)SetLoggers(val *[]INagLogger) {
	if err := j.validateSetLoggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggers",
		val,
	)
}

func (j *jsiiProxy_HIPAASecurityChecks)SetPackGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"packGlobalSuppressionIgnore",
		val,
	)
}

func (j *jsiiProxy_HIPAASecurityChecks)SetPackName(val *string) {
	if err := j.validateSetPackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packName",
		val,
	)
}

func (j *jsiiProxy_HIPAASecurityChecks)SetUserGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"userGlobalSuppressionIgnore",
		val,
	)
}

func (h *jsiiProxy_HIPAASecurityChecks) ApplyRule(params IApplyRule) {
	if err := h.validateApplyRuleParameters(params); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"applyRule",
		[]interface{}{params},
	)
}

func (h *jsiiProxy_HIPAASecurityChecks) IgnoreRule(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel, ignoreSuppressionCondition INagSuppressionIgnore) *string {
	if err := h.validateIgnoreRuleParameters(suppressions, ruleId, findingId, resource, level); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"ignoreRule",
		[]interface{}{suppressions, ruleId, findingId, resource, level, ignoreSuppressionCondition},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HIPAASecurityChecks) Visit(node constructs.IConstruct) {
	if err := h.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"visit",
		[]interface{}{node},
	)
}

