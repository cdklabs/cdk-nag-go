package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Check for NIST 800-53 rev 4 compliance.
//
// Based on the NIST 800-53 rev 4 AWS operational best practices: https://docs.aws.amazon.com/config/latest/developerguide/operational-best-practices-for-nist-800-53_rev_4.html
type NIST80053R4Checks interface {
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
	IgnoreRule(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel, ignoreSuppressionCondition INagSuppressionIgnore, validationFailure *bool) *string
	// All aspects can visit an IConstruct.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for NIST80053R4Checks
type jsiiProxy_NIST80053R4Checks struct {
	jsiiProxy_NagPack
}

func (j *jsiiProxy_NIST80053R4Checks) Loggers() *[]INagLogger {
	var returns *[]INagLogger
	_jsii_.Get(
		j,
		"loggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R4Checks) PackGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"packGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R4Checks) PackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R4Checks) ReadPackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readPackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R4Checks) UserGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"userGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}


func NewNIST80053R4Checks(props *NagPackProps) NIST80053R4Checks {
	_init_.Initialize()

	if err := validateNewNIST80053R4ChecksParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NIST80053R4Checks{}

	_jsii_.Create(
		"cdk-nag.NIST80053R4Checks",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewNIST80053R4Checks_Override(n NIST80053R4Checks, props *NagPackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.NIST80053R4Checks",
		[]interface{}{props},
		n,
	)
}

func (j *jsiiProxy_NIST80053R4Checks)SetLoggers(val *[]INagLogger) {
	if err := j.validateSetLoggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggers",
		val,
	)
}

func (j *jsiiProxy_NIST80053R4Checks)SetPackGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"packGlobalSuppressionIgnore",
		val,
	)
}

func (j *jsiiProxy_NIST80053R4Checks)SetPackName(val *string) {
	if err := j.validateSetPackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packName",
		val,
	)
}

func (j *jsiiProxy_NIST80053R4Checks)SetUserGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"userGlobalSuppressionIgnore",
		val,
	)
}

func (n *jsiiProxy_NIST80053R4Checks) ApplyRule(params IApplyRule) {
	if err := n.validateApplyRuleParameters(params); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"applyRule",
		[]interface{}{params},
	)
}

func (n *jsiiProxy_NIST80053R4Checks) IgnoreRule(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel, ignoreSuppressionCondition INagSuppressionIgnore, validationFailure *bool) *string {
	if err := n.validateIgnoreRuleParameters(suppressions, ruleId, findingId, resource, level); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"ignoreRule",
		[]interface{}{suppressions, ruleId, findingId, resource, level, ignoreSuppressionCondition, validationFailure},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NIST80053R4Checks) Visit(node constructs.IConstruct) {
	if err := n.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"visit",
		[]interface{}{node},
	)
}

