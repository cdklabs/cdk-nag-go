package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v3/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Check Best practices based on AWS Solutions Security Matrix.
// Experimental.
type AwsSolutionsChecks interface {
	NagPack
	// The name of the plugin that will be displayed in the validation report.
	// Experimental.
	Name() *string
	// Experimental.
	PackName() *string
	// Experimental.
	SetPackName(val *string)
	// Experimental.
	ReadPackName() *string
	// The list of rule IDs that the plugin will evaluate.
	//
	// Used for analytics
	// purposes.
	// Experimental.
	RuleIds() *[]*string
	// The version of the plugin, following the Semantic Versioning specification (see https://semver.org/). This version is used for analytics purposes, to measure the usage of different plugins and different versions. The value of this property should be kept in sync with the actual version of the software package. If the version is not provided or is not a valid semantic version, it will be reported as `0.0.0`.
	// Experimental.
	Version() *string
	// Create a rule to be used in the NagPack.
	// Experimental.
	ApplyRule(params IApplyRule)
	// Subclasses implement this to apply rules to each CfnResource.
	// Experimental.
	CheckResource(node awscdk.CfnResource)
	// Entry point called by the CDK validation framework.
	//
	// Requires `appConstruct` to be present on the context (CDK core change).
	// For testing or direct invocation, use `validateScope(scope)`.
	// Experimental.
	Validate(context awscdk.IPolicyValidationContext) *awscdk.PolicyValidationPluginReport
	// Validate a construct tree directly.
	//
	// This is the primary entry point
	// for testing and for CDK versions that do not yet provide `appConstruct` on
	// `IPolicyValidationContext`.
	// Experimental.
	ValidateScope(scope constructs.IConstruct) *awscdk.PolicyValidationPluginReport
}

// The jsii proxy struct for AwsSolutionsChecks
type jsiiProxy_AwsSolutionsChecks struct {
	jsiiProxy_NagPack
}

func (j *jsiiProxy_AwsSolutionsChecks) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
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

func (j *jsiiProxy_AwsSolutionsChecks) RuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ruleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSolutionsChecks) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSolutionsChecks(scope constructs.IConstruct, props *NagPackProps) AwsSolutionsChecks {
	_init_.Initialize()

	if err := validateNewAwsSolutionsChecksParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSolutionsChecks{}

	_jsii_.Create(
		"cdk-nag.AwsSolutionsChecks",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSolutionsChecks_Override(a AwsSolutionsChecks, scope constructs.IConstruct, props *NagPackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.AwsSolutionsChecks",
		[]interface{}{scope, props},
		a,
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

func (a *jsiiProxy_AwsSolutionsChecks) CheckResource(node awscdk.CfnResource) {
	if err := a.validateCheckResourceParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"checkResource",
		[]interface{}{node},
	)
}

func (a *jsiiProxy_AwsSolutionsChecks) Validate(context awscdk.IPolicyValidationContext) *awscdk.PolicyValidationPluginReport {
	if err := a.validateValidateParameters(context); err != nil {
		panic(err)
	}
	var returns *awscdk.PolicyValidationPluginReport

	_jsii_.Invoke(
		a,
		"validate",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSolutionsChecks) ValidateScope(scope constructs.IConstruct) *awscdk.PolicyValidationPluginReport {
	if err := a.validateValidateScopeParameters(scope); err != nil {
		panic(err)
	}
	var returns *awscdk.PolicyValidationPluginReport

	_jsii_.Invoke(
		a,
		"validateScope",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

