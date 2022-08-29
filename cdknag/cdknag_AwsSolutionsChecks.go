// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Check Best practices based on AWS Solutions Security Matrix.
type AwsSolutionsChecks interface {
	NagPack
	LogIgnores() *bool
	SetLogIgnores(val *bool)
	PackName() *string
	SetPackName(val *string)
	ReadPackName() *string
	ReadReportStacks() *[]*string
	Reports() *bool
	SetReports(val *bool)
	ReportStacks() *[]*string
	SetReportStacks(val *[]*string)
	Verbose() *bool
	SetVerbose(val *bool)
	// Create a rule to be used in the NagPack.
	ApplyRule(params IApplyRule)
	// Helper function to create a line for the compliance report.
	CreateComplianceReportLine(params IApplyRule, ruleId *string, compliance interface{}, explanation *string) *string
	// The message to output to the console when a rule is triggered.
	//
	// Returns: The formatted message string.
	CreateMessage(ruleId *string, findingId *string, info *string, explanation *string) *string
	// Check whether a specific rule should be ignored.
	//
	// Returns: The reason the rule was ignored, or an empty string.
	IgnoreRule(ignores *[]*NagPackSuppression, ruleId *string, findingId *string) *string
	// Initialize the report for the rule pack's compliance report for the resource's Stack if it doesn't exist.
	InitializeStackReport(params IApplyRule)
	// All aspects can visit an IConstruct.
	Visit(node constructs.IConstruct)
	// Write a line to the rule pack's compliance report for the resource's Stack.
	WriteToStackComplianceReport(params IApplyRule, ruleId *string, compliance interface{}, explanation *string)
}

// The jsii proxy struct for AwsSolutionsChecks
type jsiiProxy_AwsSolutionsChecks struct {
	jsiiProxy_NagPack
}

func (j *jsiiProxy_AwsSolutionsChecks) LogIgnores() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logIgnores",
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

func (j *jsiiProxy_AwsSolutionsChecks) ReadReportStacks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readReportStacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSolutionsChecks) Reports() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"reports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSolutionsChecks) ReportStacks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reportStacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSolutionsChecks) Verbose() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"verbose",
		&returns,
	)
	return returns
}


func NewAwsSolutionsChecks(props *NagPackProps) AwsSolutionsChecks {
	_init_.Initialize()

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

func (j *jsiiProxy_AwsSolutionsChecks) SetLogIgnores(val *bool) {
	_jsii_.Set(
		j,
		"logIgnores",
		val,
	)
}

func (j *jsiiProxy_AwsSolutionsChecks) SetPackName(val *string) {
	_jsii_.Set(
		j,
		"packName",
		val,
	)
}

func (j *jsiiProxy_AwsSolutionsChecks) SetReports(val *bool) {
	_jsii_.Set(
		j,
		"reports",
		val,
	)
}

func (j *jsiiProxy_AwsSolutionsChecks) SetReportStacks(val *[]*string) {
	_jsii_.Set(
		j,
		"reportStacks",
		val,
	)
}

func (j *jsiiProxy_AwsSolutionsChecks) SetVerbose(val *bool) {
	_jsii_.Set(
		j,
		"verbose",
		val,
	)
}

func (a *jsiiProxy_AwsSolutionsChecks) ApplyRule(params IApplyRule) {
	_jsii_.InvokeVoid(
		a,
		"applyRule",
		[]interface{}{params},
	)
}

func (a *jsiiProxy_AwsSolutionsChecks) CreateComplianceReportLine(params IApplyRule, ruleId *string, compliance interface{}, explanation *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"createComplianceReportLine",
		[]interface{}{params, ruleId, compliance, explanation},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSolutionsChecks) CreateMessage(ruleId *string, findingId *string, info *string, explanation *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"createMessage",
		[]interface{}{ruleId, findingId, info, explanation},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSolutionsChecks) IgnoreRule(ignores *[]*NagPackSuppression, ruleId *string, findingId *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"ignoreRule",
		[]interface{}{ignores, ruleId, findingId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSolutionsChecks) InitializeStackReport(params IApplyRule) {
	_jsii_.InvokeVoid(
		a,
		"initializeStackReport",
		[]interface{}{params},
	)
}

func (a *jsiiProxy_AwsSolutionsChecks) Visit(node constructs.IConstruct) {
	_jsii_.InvokeVoid(
		a,
		"visit",
		[]interface{}{node},
	)
}

func (a *jsiiProxy_AwsSolutionsChecks) WriteToStackComplianceReport(params IApplyRule, ruleId *string, compliance interface{}, explanation *string) {
	_jsii_.InvokeVoid(
		a,
		"writeToStackComplianceReport",
		[]interface{}{params, ruleId, compliance, explanation},
	)
}

