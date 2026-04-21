package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Check for NIST 800-53 rev 5 compliance.
//
// Based on the NIST 800-53 rev 5 AWS operational best practices: https://docs.aws.amazon.com/config/latest/developerguide/operational-best-practices-for-nist-800-53_rev_5.html
// Experimental.
type NIST80053R5Checks interface {
	NagPack
	// Experimental.
	Loggers() *[]INagLogger
	// Experimental.
	SetLoggers(val *[]INagLogger)
	// Experimental.
	PackGlobalSuppressionIgnore() INagSuppressionIgnore
	// Experimental.
	SetPackGlobalSuppressionIgnore(val INagSuppressionIgnore)
	// Experimental.
	PackName() *string
	// Experimental.
	SetPackName(val *string)
	// Experimental.
	ReadPackName() *string
	// Experimental.
	UserGlobalSuppressionIgnore() INagSuppressionIgnore
	// Experimental.
	SetUserGlobalSuppressionIgnore(val INagSuppressionIgnore)
	// Create a rule to be used in the NagPack.
	// Experimental.
	ApplyRule(params IApplyRule)
	// Check whether a specific rule should be ignored.
	//
	// Returns: The reason the rule was ignored, or an empty string.
	// Experimental.
	IgnoreRule(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel, ignoreSuppressionCondition INagSuppressionIgnore, validationFailure *bool) *string
	// All aspects can visit an IConstruct.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for NIST80053R5Checks
type jsiiProxy_NIST80053R5Checks struct {
	jsiiProxy_NagPack
}

func (j *jsiiProxy_NIST80053R5Checks) Loggers() *[]INagLogger {
	var returns *[]INagLogger
	_jsii_.Get(
		j,
		"loggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R5Checks) PackGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"packGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R5Checks) PackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R5Checks) ReadPackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readPackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R5Checks) UserGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"userGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}


// Experimental.
func NewNIST80053R5Checks(props *NagPackProps) NIST80053R5Checks {
	_init_.Initialize()

	if err := validateNewNIST80053R5ChecksParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NIST80053R5Checks{}

	_jsii_.Create(
		"cdk-nag.NIST80053R5Checks",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewNIST80053R5Checks_Override(n NIST80053R5Checks, props *NagPackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.NIST80053R5Checks",
		[]interface{}{props},
		n,
	)
}

func (j *jsiiProxy_NIST80053R5Checks)SetLoggers(val *[]INagLogger) {
	if err := j.validateSetLoggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggers",
		val,
	)
}

func (j *jsiiProxy_NIST80053R5Checks)SetPackGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"packGlobalSuppressionIgnore",
		val,
	)
}

func (j *jsiiProxy_NIST80053R5Checks)SetPackName(val *string) {
	if err := j.validateSetPackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packName",
		val,
	)
}

func (j *jsiiProxy_NIST80053R5Checks)SetUserGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"userGlobalSuppressionIgnore",
		val,
	)
}

func (n *jsiiProxy_NIST80053R5Checks) ApplyRule(params IApplyRule) {
	if err := n.validateApplyRuleParameters(params); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"applyRule",
		[]interface{}{params},
	)
}

func (n *jsiiProxy_NIST80053R5Checks) IgnoreRule(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel, ignoreSuppressionCondition INagSuppressionIgnore, validationFailure *bool) *string {
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

func (n *jsiiProxy_NIST80053R5Checks) Visit(node constructs.IConstruct) {
	if err := n.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"visit",
		[]interface{}{node},
	)
}

