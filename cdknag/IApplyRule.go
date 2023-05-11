package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Interface for JSII interoperability for passing parameters and the Rule Callback to @applyRule method.
type IApplyRule interface {
	// The callback to the rule.
	Rule(node awscdk.CfnResource) interface{}
	// Why the rule exists.
	Explanation() *string
	SetExplanation(e *string)
	// A condition in which a suppression should be ignored.
	IgnoreSuppressionCondition() INagSuppressionIgnore
	SetIgnoreSuppressionCondition(i INagSuppressionIgnore)
	// Why the rule was triggered.
	Info() *string
	SetInfo(i *string)
	// The annotations message level to apply to the rule if triggered.
	Level() NagMessageLevel
	SetLevel(l NagMessageLevel)
	// The CfnResource to check.
	Node() awscdk.CfnResource
	SetNode(n awscdk.CfnResource)
	// Override for the suffix of the Rule ID for this rule.
	RuleSuffixOverride() *string
	SetRuleSuffixOverride(r *string)
}

// The jsii proxy for IApplyRule
type jsiiProxy_IApplyRule struct {
	_ byte // padding
}

func (i *jsiiProxy_IApplyRule) Rule(node awscdk.CfnResource) interface{} {
	if err := i.validateRuleParameters(node); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"rule",
		[]interface{}{node},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IApplyRule) Explanation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"explanation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplyRule)SetExplanation(val *string) {
	if err := j.validateSetExplanationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"explanation",
		val,
	)
}

func (j *jsiiProxy_IApplyRule) IgnoreSuppressionCondition() INagSuppressionIgnore {
	var returns INagSuppressionIgnore
	_jsii_.Get(
		j,
		"ignoreSuppressionCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplyRule)SetIgnoreSuppressionCondition(val INagSuppressionIgnore) {
	_jsii_.Set(
		j,
		"ignoreSuppressionCondition",
		val,
	)
}

func (j *jsiiProxy_IApplyRule) Info() *string {
	var returns *string
	_jsii_.Get(
		j,
		"info",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplyRule)SetInfo(val *string) {
	if err := j.validateSetInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"info",
		val,
	)
}

func (j *jsiiProxy_IApplyRule) Level() NagMessageLevel {
	var returns NagMessageLevel
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplyRule)SetLevel(val NagMessageLevel) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_IApplyRule) Node() awscdk.CfnResource {
	var returns awscdk.CfnResource
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplyRule)SetNode(val awscdk.CfnResource) {
	if err := j.validateSetNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"node",
		val,
	)
}

func (j *jsiiProxy_IApplyRule) RuleSuffixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleSuffixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplyRule)SetRuleSuffixOverride(val *string) {
	_jsii_.Set(
		j,
		"ruleSuffixOverride",
		val,
	)
}

