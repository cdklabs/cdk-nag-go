// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Check Best practices based on AWS Solutions Security Matrix.
type AwsSolutionsChecks interface {
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

// The jsii proxy struct for AwsSolutionsChecks
type jsiiProxy_AwsSolutionsChecks struct {
	jsiiProxy_NagPack
}

func (j *jsiiProxy_AwsSolutionsChecks) Loggers() *[]INagLogger {
	var returns *[]INagLogger
	_jsii_.Get(
		j,
		"loggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSolutionsChecks) PackGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"packGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSolutionsChecks) PackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSolutionsChecks) ReadPackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readPackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSolutionsChecks) UserGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"userGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}


func NewAwsSolutionsChecks(props *NagPackProps) AwsSolutionsChecks {
	_init_.Initialize()

	if err := validateNewAwsSolutionsChecksParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSolutionsChecks{}

	_jsii_.Create(
		"cdk-nag.AwsSolutionsChecks",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAwsSolutionsChecks_Override(a AwsSolutionsChecks, props *NagPackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.AwsSolutionsChecks",
		[]interface{}{props},
		a,
	)
}

func (j *jsiiProxy_AwsSolutionsChecks)SetLoggers(val *[]INagLogger) {
	if err := j.validateSetLoggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggers",
		val,
	)
}

func (j *jsiiProxy_AwsSolutionsChecks)SetPackGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"packGlobalSuppressionIgnore",
		val,
	)
}

func (j *jsiiProxy_AwsSolutionsChecks)SetPackName(val *string) {
	if err := j.validateSetPackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packName",
		val,
	)
}

func (j *jsiiProxy_AwsSolutionsChecks)SetUserGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"userGlobalSuppressionIgnore",
		val,
	)
}

func (a *jsiiProxy_AwsSolutionsChecks) ApplyRule(params IApplyRule) {
	if err := a.validateApplyRuleParameters(params); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRule",
		[]interface{}{params},
	)
}

func (a *jsiiProxy_AwsSolutionsChecks) IgnoreRule(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel, ignoreSuppressionCondition INagSuppressionIgnore) *string {
	if err := a.validateIgnoreRuleParameters(suppressions, ruleId, findingId, resource, level); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"ignoreRule",
		[]interface{}{suppressions, ruleId, findingId, resource, level, ignoreSuppressionCondition},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSolutionsChecks) Visit(node constructs.IConstruct) {
	if err := a.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"visit",
		[]interface{}{node},
	)
}

