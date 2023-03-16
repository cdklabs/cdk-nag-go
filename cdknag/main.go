package cdknag

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-nag.AwsSolutionsChecks",
		reflect.TypeOf((*AwsSolutionsChecks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRule", GoMethod: "ApplyRule"},
			_jsii_.MemberMethod{JsiiMethod: "createComplianceReportLine", GoMethod: "CreateComplianceReportLine"},
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
			_jsii_.MemberMethod{JsiiMethod: "ignoreRule", GoMethod: "IgnoreRule"},
			_jsii_.MemberMethod{JsiiMethod: "initializeStackReport", GoMethod: "InitializeStackReport"},
			_jsii_.MemberProperty{JsiiProperty: "logIgnores", GoGetter: "LogIgnores"},
			_jsii_.MemberProperty{JsiiProperty: "packGlobalSuppressionIgnore", GoGetter: "PackGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "readReportStacks", GoGetter: "ReadReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "reports", GoGetter: "Reports"},
			_jsii_.MemberProperty{JsiiProperty: "reportStacks", GoGetter: "ReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "userGlobalSuppressionIgnore", GoGetter: "UserGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "verbose", GoGetter: "Verbose"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
			_jsii_.MemberMethod{JsiiMethod: "writeToStackComplianceReport", GoMethod: "WriteToStackComplianceReport"},
		},
		func() interface{} {
			j := jsiiProxy_AwsSolutionsChecks{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NagPack)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.HIPAASecurityChecks",
		reflect.TypeOf((*HIPAASecurityChecks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRule", GoMethod: "ApplyRule"},
			_jsii_.MemberMethod{JsiiMethod: "createComplianceReportLine", GoMethod: "CreateComplianceReportLine"},
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
			_jsii_.MemberMethod{JsiiMethod: "ignoreRule", GoMethod: "IgnoreRule"},
			_jsii_.MemberMethod{JsiiMethod: "initializeStackReport", GoMethod: "InitializeStackReport"},
			_jsii_.MemberProperty{JsiiProperty: "logIgnores", GoGetter: "LogIgnores"},
			_jsii_.MemberProperty{JsiiProperty: "packGlobalSuppressionIgnore", GoGetter: "PackGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "readReportStacks", GoGetter: "ReadReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "reports", GoGetter: "Reports"},
			_jsii_.MemberProperty{JsiiProperty: "reportStacks", GoGetter: "ReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "userGlobalSuppressionIgnore", GoGetter: "UserGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "verbose", GoGetter: "Verbose"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
			_jsii_.MemberMethod{JsiiMethod: "writeToStackComplianceReport", GoMethod: "WriteToStackComplianceReport"},
		},
		func() interface{} {
			j := jsiiProxy_HIPAASecurityChecks{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NagPack)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"cdk-nag.IApplyRule",
		reflect.TypeOf((*IApplyRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "explanation", GoGetter: "Explanation"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreSuppressionCondition", GoGetter: "IgnoreSuppressionCondition"},
			_jsii_.MemberProperty{JsiiProperty: "info", GoGetter: "Info"},
			_jsii_.MemberProperty{JsiiProperty: "level", GoGetter: "Level"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "rule", GoMethod: "Rule"},
			_jsii_.MemberProperty{JsiiProperty: "ruleSuffixOverride", GoGetter: "RuleSuffixOverride"},
		},
		func() interface{} {
			return &jsiiProxy_IApplyRule{}
		},
	)
	_jsii_.RegisterInterface(
		"cdk-nag.INagSuppressionIgnore",
		reflect.TypeOf((*INagSuppressionIgnore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
		},
		func() interface{} {
			return &jsiiProxy_INagSuppressionIgnore{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.NIST80053R4Checks",
		reflect.TypeOf((*NIST80053R4Checks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRule", GoMethod: "ApplyRule"},
			_jsii_.MemberMethod{JsiiMethod: "createComplianceReportLine", GoMethod: "CreateComplianceReportLine"},
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
			_jsii_.MemberMethod{JsiiMethod: "ignoreRule", GoMethod: "IgnoreRule"},
			_jsii_.MemberMethod{JsiiMethod: "initializeStackReport", GoMethod: "InitializeStackReport"},
			_jsii_.MemberProperty{JsiiProperty: "logIgnores", GoGetter: "LogIgnores"},
			_jsii_.MemberProperty{JsiiProperty: "packGlobalSuppressionIgnore", GoGetter: "PackGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "readReportStacks", GoGetter: "ReadReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "reports", GoGetter: "Reports"},
			_jsii_.MemberProperty{JsiiProperty: "reportStacks", GoGetter: "ReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "userGlobalSuppressionIgnore", GoGetter: "UserGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "verbose", GoGetter: "Verbose"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
			_jsii_.MemberMethod{JsiiMethod: "writeToStackComplianceReport", GoMethod: "WriteToStackComplianceReport"},
		},
		func() interface{} {
			j := jsiiProxy_NIST80053R4Checks{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NagPack)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.NIST80053R5Checks",
		reflect.TypeOf((*NIST80053R5Checks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRule", GoMethod: "ApplyRule"},
			_jsii_.MemberMethod{JsiiMethod: "createComplianceReportLine", GoMethod: "CreateComplianceReportLine"},
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
			_jsii_.MemberMethod{JsiiMethod: "ignoreRule", GoMethod: "IgnoreRule"},
			_jsii_.MemberMethod{JsiiMethod: "initializeStackReport", GoMethod: "InitializeStackReport"},
			_jsii_.MemberProperty{JsiiProperty: "logIgnores", GoGetter: "LogIgnores"},
			_jsii_.MemberProperty{JsiiProperty: "packGlobalSuppressionIgnore", GoGetter: "PackGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "readReportStacks", GoGetter: "ReadReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "reports", GoGetter: "Reports"},
			_jsii_.MemberProperty{JsiiProperty: "reportStacks", GoGetter: "ReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "userGlobalSuppressionIgnore", GoGetter: "UserGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "verbose", GoGetter: "Verbose"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
			_jsii_.MemberMethod{JsiiMethod: "writeToStackComplianceReport", GoMethod: "WriteToStackComplianceReport"},
		},
		func() interface{} {
			j := jsiiProxy_NIST80053R5Checks{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NagPack)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"cdk-nag.NagMessageLevel",
		reflect.TypeOf((*NagMessageLevel)(nil)).Elem(),
		map[string]interface{}{
			"WARN": NagMessageLevel_WARN,
			"ERROR": NagMessageLevel_ERROR,
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.NagPack",
		reflect.TypeOf((*NagPack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRule", GoMethod: "ApplyRule"},
			_jsii_.MemberMethod{JsiiMethod: "createComplianceReportLine", GoMethod: "CreateComplianceReportLine"},
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
			_jsii_.MemberMethod{JsiiMethod: "ignoreRule", GoMethod: "IgnoreRule"},
			_jsii_.MemberMethod{JsiiMethod: "initializeStackReport", GoMethod: "InitializeStackReport"},
			_jsii_.MemberProperty{JsiiProperty: "logIgnores", GoGetter: "LogIgnores"},
			_jsii_.MemberProperty{JsiiProperty: "packGlobalSuppressionIgnore", GoGetter: "PackGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "readReportStacks", GoGetter: "ReadReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "reports", GoGetter: "Reports"},
			_jsii_.MemberProperty{JsiiProperty: "reportStacks", GoGetter: "ReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "userGlobalSuppressionIgnore", GoGetter: "UserGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "verbose", GoGetter: "Verbose"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
			_jsii_.MemberMethod{JsiiMethod: "writeToStackComplianceReport", GoMethod: "WriteToStackComplianceReport"},
		},
		func() interface{} {
			j := jsiiProxy_NagPack{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-nag.NagPackProps",
		reflect.TypeOf((*NagPackProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-nag.NagPackSuppression",
		reflect.TypeOf((*NagPackSuppression)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-nag.NagRuleCompliance",
		reflect.TypeOf((*NagRuleCompliance)(nil)).Elem(),
		map[string]interface{}{
			"COMPLIANT": NagRuleCompliance_COMPLIANT,
			"NON_COMPLIANT": NagRuleCompliance_NON_COMPLIANT,
			"NOT_APPLICABLE": NagRuleCompliance_NOT_APPLICABLE,
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.NagRules",
		reflect.TypeOf((*NagRules)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_NagRules{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.NagSuppressions",
		reflect.TypeOf((*NagSuppressions)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_NagSuppressions{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.PCIDSS321Checks",
		reflect.TypeOf((*PCIDSS321Checks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRule", GoMethod: "ApplyRule"},
			_jsii_.MemberMethod{JsiiMethod: "createComplianceReportLine", GoMethod: "CreateComplianceReportLine"},
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
			_jsii_.MemberMethod{JsiiMethod: "ignoreRule", GoMethod: "IgnoreRule"},
			_jsii_.MemberMethod{JsiiMethod: "initializeStackReport", GoMethod: "InitializeStackReport"},
			_jsii_.MemberProperty{JsiiProperty: "logIgnores", GoGetter: "LogIgnores"},
			_jsii_.MemberProperty{JsiiProperty: "packGlobalSuppressionIgnore", GoGetter: "PackGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "readReportStacks", GoGetter: "ReadReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "reports", GoGetter: "Reports"},
			_jsii_.MemberProperty{JsiiProperty: "reportStacks", GoGetter: "ReportStacks"},
			_jsii_.MemberProperty{JsiiProperty: "userGlobalSuppressionIgnore", GoGetter: "UserGlobalSuppressionIgnore"},
			_jsii_.MemberProperty{JsiiProperty: "verbose", GoGetter: "Verbose"},
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
			_jsii_.MemberMethod{JsiiMethod: "writeToStackComplianceReport", GoMethod: "WriteToStackComplianceReport"},
		},
		func() interface{} {
			j := jsiiProxy_PCIDSS321Checks{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NagPack)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-nag.RegexAppliesTo",
		reflect.TypeOf((*RegexAppliesTo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-nag.SuppressionIgnoreAlways",
		reflect.TypeOf((*SuppressionIgnoreAlways)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
		},
		func() interface{} {
			j := jsiiProxy_SuppressionIgnoreAlways{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INagSuppressionIgnore)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.SuppressionIgnoreAnd",
		reflect.TypeOf((*SuppressionIgnoreAnd)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
		},
		func() interface{} {
			j := jsiiProxy_SuppressionIgnoreAnd{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INagSuppressionIgnore)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.SuppressionIgnoreErrors",
		reflect.TypeOf((*SuppressionIgnoreErrors)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
		},
		func() interface{} {
			j := jsiiProxy_SuppressionIgnoreErrors{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INagSuppressionIgnore)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-nag.SuppressionIgnoreInput",
		reflect.TypeOf((*SuppressionIgnoreInput)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-nag.SuppressionIgnoreNever",
		reflect.TypeOf((*SuppressionIgnoreNever)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
		},
		func() interface{} {
			j := jsiiProxy_SuppressionIgnoreNever{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INagSuppressionIgnore)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.SuppressionIgnoreOr",
		reflect.TypeOf((*SuppressionIgnoreOr)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createMessage", GoMethod: "CreateMessage"},
		},
		func() interface{} {
			j := jsiiProxy_SuppressionIgnoreOr{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INagSuppressionIgnore)
			return &j
		},
	)
}
