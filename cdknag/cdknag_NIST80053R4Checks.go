// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Check for NIST 800-53 rev 4 compliance.
//
// Based on the NIST 800-53 rev 4 AWS operational best practices: https://docs.aws.amazon.com/config/latest/developerguide/operational-best-practices-for-nist-800-53_rev_4.html
type NIST80053R4Checks interface {
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

// The jsii proxy struct for NIST80053R4Checks
type jsiiProxy_NIST80053R4Checks struct {
	jsiiProxy_NagPack
}

func (j *jsiiProxy_NIST80053R4Checks) LogIgnores() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logIgnores",
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

func (j *jsiiProxy_NIST80053R4Checks) ReadReportStacks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readReportStacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R4Checks) Reports() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"reports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R4Checks) ReportStacks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reportStacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NIST80053R4Checks) Verbose() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"verbose",
		&returns,
	)
	return returns
}


func NewNIST80053R4Checks(props *NagPackProps) NIST80053R4Checks {
	_init_.Initialize()

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

func (j *jsiiProxy_NIST80053R4Checks) SetLogIgnores(val *bool) {
	_jsii_.Set(
		j,
		"logIgnores",
		val,
	)
}

func (j *jsiiProxy_NIST80053R4Checks) SetPackName(val *string) {
	_jsii_.Set(
		j,
		"packName",
		val,
	)
}

func (j *jsiiProxy_NIST80053R4Checks) SetReports(val *bool) {
	_jsii_.Set(
		j,
		"reports",
		val,
	)
}

func (j *jsiiProxy_NIST80053R4Checks) SetReportStacks(val *[]*string) {
	_jsii_.Set(
		j,
		"reportStacks",
		val,
	)
}

func (j *jsiiProxy_NIST80053R4Checks) SetVerbose(val *bool) {
	_jsii_.Set(
		j,
		"verbose",
		val,
	)
}

func (n *jsiiProxy_NIST80053R4Checks) ApplyRule(params IApplyRule) {
	_jsii_.InvokeVoid(
		n,
		"applyRule",
		[]interface{}{params},
	)
}

func (n *jsiiProxy_NIST80053R4Checks) CreateComplianceReportLine(params IApplyRule, ruleId *string, compliance interface{}, explanation *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"createComplianceReportLine",
		[]interface{}{params, ruleId, compliance, explanation},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NIST80053R4Checks) CreateMessage(ruleId *string, findingId *string, info *string, explanation *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"createMessage",
		[]interface{}{ruleId, findingId, info, explanation},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NIST80053R4Checks) IgnoreRule(ignores *[]*NagPackSuppression, ruleId *string, findingId *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"ignoreRule",
		[]interface{}{ignores, ruleId, findingId},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NIST80053R4Checks) InitializeStackReport(params IApplyRule) {
	_jsii_.InvokeVoid(
		n,
		"initializeStackReport",
		[]interface{}{params},
	)
}

func (n *jsiiProxy_NIST80053R4Checks) Visit(node constructs.IConstruct) {
	_jsii_.InvokeVoid(
		n,
		"visit",
		[]interface{}{node},
	)
}

func (n *jsiiProxy_NIST80053R4Checks) WriteToStackComplianceReport(params IApplyRule, ruleId *string, compliance interface{}, explanation *string) {
	_jsii_.InvokeVoid(
		n,
		"writeToStackComplianceReport",
		[]interface{}{params, ruleId, compliance, explanation},
	)
}

