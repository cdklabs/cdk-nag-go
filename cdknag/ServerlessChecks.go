package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Serverless Checks are a compilation of rules to validate infrastructure-as-code template against recommended practices.
type ServerlessChecks interface {
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

// The jsii proxy struct for ServerlessChecks
type jsiiProxy_ServerlessChecks struct {
	jsiiProxy_NagPack
}

func (j *jsiiProxy_ServerlessChecks) Loggers() *[]INagLogger {
	var returns *[]INagLogger
	_jsii_.Get(
		j,
		"loggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessChecks) PackGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"packGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessChecks) PackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessChecks) ReadPackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readPackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerlessChecks) UserGlobalSuppressionIgnore() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"userGlobalSuppressionIgnore",
		&returns,
	)
	return returns
}


func NewServerlessChecks(props *NagPackProps) ServerlessChecks {
	_init_.Initialize()

	if err := validateNewServerlessChecksParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServerlessChecks{}

	_jsii_.Create(
		"cdk-nag.ServerlessChecks",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewServerlessChecks_Override(s ServerlessChecks, props *NagPackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.ServerlessChecks",
		[]interface{}{props},
		s,
	)
}

func (j *jsiiProxy_ServerlessChecks)SetLoggers(val *[]INagLogger) {
	if err := j.validateSetLoggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggers",
		val,
	)
}

func (j *jsiiProxy_ServerlessChecks)SetPackGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"packGlobalSuppressionIgnore",
		val,
	)
}

func (j *jsiiProxy_ServerlessChecks)SetPackName(val *string) {
	if err := j.validateSetPackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packName",
		val,
	)
}

func (j *jsiiProxy_ServerlessChecks)SetUserGlobalSuppressionIgnore(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"userGlobalSuppressionIgnore",
		val,
	)
}

func (s *jsiiProxy_ServerlessChecks) ApplyRule(params IApplyRule) {
	if err := s.validateApplyRuleParameters(params); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRule",
		[]interface{}{params},
	)
}

func (s *jsiiProxy_ServerlessChecks) IgnoreRule(suppressions *[]*NagPackSuppression, ruleId *string, findingId *string, resource awscdk.CfnResource, level NagMessageLevel, ignoreSuppressionCondition INagSuppressionIgnore, validationFailure *bool) *string {
	if err := s.validateIgnoreRuleParameters(suppressions, ruleId, findingId, resource, level); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"ignoreRule",
		[]interface{}{suppressions, ruleId, findingId, resource, level, ignoreSuppressionCondition, validationFailure},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerlessChecks) Visit(node constructs.IConstruct) {
	if err := s.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"visit",
		[]interface{}{node},
	)
}

