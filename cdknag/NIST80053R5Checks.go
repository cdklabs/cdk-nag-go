package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v3/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Check for NIST 800-53 rev 5 compliance.
//
// Based on the NIST 800-53 rev 5 AWS operational best practices: https://docs.aws.amazon.com/config/latest/developerguide/operational-best-practices-for-nist-800-53_rev_5.html
// Experimental.
type NIST80053R5Checks interface {
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

// The jsii proxy struct for NIST80053R5Checks
type jsiiProxy_NIST80053R5Checks struct {
	jsiiProxy_NagPack
}

func (j *jsiiProxy_NIST80053R5Checks) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
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

func (j *jsiiProxy_NIST80053R5Checks) RuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ruleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R5Checks) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Experimental.
func NewNIST80053R5Checks(scope constructs.IConstruct, props *NagPackProps) NIST80053R5Checks {
	_init_.Initialize()

	if err := validateNewNIST80053R5ChecksParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NIST80053R5Checks{}

	_jsii_.Create(
		"cdk-nag.NIST80053R5Checks",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewNIST80053R5Checks_Override(n NIST80053R5Checks, scope constructs.IConstruct, props *NagPackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.NIST80053R5Checks",
		[]interface{}{scope, props},
		n,
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

func (n *jsiiProxy_NIST80053R5Checks) CheckResource(node awscdk.CfnResource) {
	if err := n.validateCheckResourceParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"checkResource",
		[]interface{}{node},
	)
}

func (n *jsiiProxy_NIST80053R5Checks) Validate(context awscdk.IPolicyValidationContext) *awscdk.PolicyValidationPluginReport {
	if err := n.validateValidateParameters(context); err != nil {
		panic(err)
	}
	var returns *awscdk.PolicyValidationPluginReport

	_jsii_.Invoke(
		n,
		"validate",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NIST80053R5Checks) ValidateScope(scope constructs.IConstruct) *awscdk.PolicyValidationPluginReport {
	if err := n.validateValidateScopeParameters(scope); err != nil {
		panic(err)
	}
	var returns *awscdk.PolicyValidationPluginReport

	_jsii_.Invoke(
		n,
		"validateScope",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

