// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Check for HIPAA Security compliance.
//
// Based on the HIPAA Security AWS operational best practices: https://docs.aws.amazon.com/config/latest/developerguide/operational-best-practices-for-hipaa_security.html
type HIPAASecurityChecks interface {
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

// The jsii proxy struct for HIPAASecurityChecks
type jsiiProxy_HIPAASecurityChecks struct {
	jsiiProxy_NagPack
}

func (j *jsiiProxy_HIPAASecurityChecks) LogIgnores() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logIgnores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HIPAASecurityChecks) PackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HIPAASecurityChecks) ReadPackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readPackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HIPAASecurityChecks) ReadReportStacks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readReportStacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HIPAASecurityChecks) Reports() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"reports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HIPAASecurityChecks) ReportStacks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reportStacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HIPAASecurityChecks) Verbose() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"verbose",
		&returns,
	)
	return returns
}


func NewHIPAASecurityChecks(props *NagPackProps) HIPAASecurityChecks {
	_init_.Initialize()

	if err := validateNewHIPAASecurityChecksParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HIPAASecurityChecks{}

	_jsii_.Create(
		"cdk-nag.HIPAASecurityChecks",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewHIPAASecurityChecks_Override(h HIPAASecurityChecks, props *NagPackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.HIPAASecurityChecks",
		[]interface{}{props},
		h,
	)
}

func (j *jsiiProxy_HIPAASecurityChecks)SetLogIgnores(val *bool) {
	if err := j.validateSetLogIgnoresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logIgnores",
		val,
	)
}

func (j *jsiiProxy_HIPAASecurityChecks)SetPackName(val *string) {
	if err := j.validateSetPackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packName",
		val,
	)
}

func (j *jsiiProxy_HIPAASecurityChecks)SetReports(val *bool) {
	if err := j.validateSetReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reports",
		val,
	)
}

func (j *jsiiProxy_HIPAASecurityChecks)SetReportStacks(val *[]*string) {
	if err := j.validateSetReportStacksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportStacks",
		val,
	)
}

func (j *jsiiProxy_HIPAASecurityChecks)SetVerbose(val *bool) {
	if err := j.validateSetVerboseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verbose",
		val,
	)
}

func (h *jsiiProxy_HIPAASecurityChecks) ApplyRule(params IApplyRule) {
	if err := h.validateApplyRuleParameters(params); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"applyRule",
		[]interface{}{params},
	)
}

func (h *jsiiProxy_HIPAASecurityChecks) CreateComplianceReportLine(params IApplyRule, ruleId *string, compliance interface{}, explanation *string) *string {
	if err := h.validateCreateComplianceReportLineParameters(params, ruleId, compliance); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"createComplianceReportLine",
		[]interface{}{params, ruleId, compliance, explanation},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HIPAASecurityChecks) CreateMessage(ruleId *string, findingId *string, info *string, explanation *string) *string {
	if err := h.validateCreateMessageParameters(ruleId, findingId, info, explanation); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"createMessage",
		[]interface{}{ruleId, findingId, info, explanation},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HIPAASecurityChecks) IgnoreRule(ignores *[]*NagPackSuppression, ruleId *string, findingId *string) *string {
	if err := h.validateIgnoreRuleParameters(ignores, ruleId, findingId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"ignoreRule",
		[]interface{}{ignores, ruleId, findingId},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HIPAASecurityChecks) InitializeStackReport(params IApplyRule) {
	if err := h.validateInitializeStackReportParameters(params); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"initializeStackReport",
		[]interface{}{params},
	)
}

func (h *jsiiProxy_HIPAASecurityChecks) Visit(node constructs.IConstruct) {
	if err := h.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"visit",
		[]interface{}{node},
	)
}

func (h *jsiiProxy_HIPAASecurityChecks) WriteToStackComplianceReport(params IApplyRule, ruleId *string, compliance interface{}, explanation *string) {
	if err := h.validateWriteToStackComplianceReportParameters(params, ruleId, compliance); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"writeToStackComplianceReport",
		[]interface{}{params, ruleId, compliance, explanation},
	)
}

