// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Check for PCI DSS 3.2.1 compliance. Based on the PCI DSS 3.2.1 AWS operational best practices: https://docs.aws.amazon.com/config/latest/developerguide/operational-best-practices-for-pci-dss.html.
type PCIDSS321Checks interface {
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

// The jsii proxy struct for PCIDSS321Checks
type jsiiProxy_PCIDSS321Checks struct {
	jsiiProxy_NagPack
}

func (j *jsiiProxy_PCIDSS321Checks) Loggers() *[]INagLogger {
	var returns *[]INagLogger
	_jsii_.Get(
		j,
		"loggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PCIDSS321Checks) PackGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"packGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PCIDSS321Checks) PackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PCIDSS321Checks) ReadPackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readPackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PCIDSS321Checks) UserGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"userGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}


func NewPCIDSS321Checks(props *NagPackProps) PCIDSS321Checks {
	_init_.Initialize()

	if err := validateNewPCIDSS321ChecksParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PCIDSS321Checks{}

	_jsii_.Create(
		"cdk-nag.PCIDSS321Checks",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPCIDSS321Checks_Override(p PCIDSS321Checks, props *NagPackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.PCIDSS321Checks",
		[]interface{}{props},
		p,
	)
}

func (j *jsiiProxy_PCIDSS321Checks)SetLoggers(val *[]INagLogger) {
	if err := j.validateSetLoggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggers",
		val,
	)
}

func (j *jsiiProxy_PCIDSS321Checks)SetPackGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"packGlobalSuppressionIgnore",
		val,
	)
}

func (j *jsiiProxy_PCIDSS321Checks)SetPackName(val *string) {
	if err := j.validateSetPackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packName",
		val,
	)
}

func (j *jsiiProxy_PCIDSS321Checks)SetUserGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"userGlobalSuppressionIgnore",
		val,
	)
}

func (p *jsiiProxy_PCIDSS321Checks) ApplyRule(params IApplyRule) {
	if err := p.validateApplyRuleParameters(params); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRule",
		[]interface{}{params},
	)
}

func (p *jsiiProxy_PCIDSS321Checks) IgnoreRule(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel, ignoreSuppressionCondition INagSuppressionIgnore) *string {
	if err := p.validateIgnoreRuleParameters(suppressions, ruleId, findingId, resource, level); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"ignoreRule",
		[]interface{}{suppressions, ruleId, findingId, resource, level, ignoreSuppressionCondition},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PCIDSS321Checks) Visit(node constructs.IConstruct) {
	if err := p.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"visit",
		[]interface{}{node},
	)
}

