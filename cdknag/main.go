// Check CDK v2 applications for best practices using a combination on available rule packs.
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
			_jsii_.MemberMethod{JsiiMethod: "checkResource", GoMethod: "CheckResource"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleIds", GoGetter: "RuleIds"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateScope", GoMethod: "ValidateScope"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
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
			_jsii_.MemberMethod{JsiiMethod: "checkResource", GoMethod: "CheckResource"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleIds", GoGetter: "RuleIds"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateScope", GoMethod: "ValidateScope"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
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
		"cdk-nag.INagValidationContext",
		reflect.TypeOf((*INagValidationContext)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "appConstruct", GoGetter: "AppConstruct"},
			_jsii_.MemberProperty{JsiiProperty: "templatePaths", GoGetter: "TemplatePaths"},
		},
		func() interface{} {
			j := jsiiProxy_INagValidationContext{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIPolicyValidationContext)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.NIST80053R4Checks",
		reflect.TypeOf((*NIST80053R4Checks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRule", GoMethod: "ApplyRule"},
			_jsii_.MemberMethod{JsiiMethod: "checkResource", GoMethod: "CheckResource"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleIds", GoGetter: "RuleIds"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateScope", GoMethod: "ValidateScope"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
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
			_jsii_.MemberMethod{JsiiMethod: "checkResource", GoMethod: "CheckResource"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleIds", GoGetter: "RuleIds"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateScope", GoMethod: "ValidateScope"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
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
			"INFO": NagMessageLevel_INFO,
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.NagPack",
		reflect.TypeOf((*NagPack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRule", GoMethod: "ApplyRule"},
			_jsii_.MemberMethod{JsiiMethod: "checkResource", GoMethod: "CheckResource"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleIds", GoGetter: "RuleIds"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateScope", GoMethod: "ValidateScope"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_NagPack{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIPolicyValidationPlugin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-nag.NagPackProps",
		reflect.TypeOf((*NagPackProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-nag.NagReportFormat",
		reflect.TypeOf((*NagReportFormat)(nil)).Elem(),
		map[string]interface{}{
			"CSV": NagReportFormat_CSV,
			"JSON": NagReportFormat_JSON,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-nag.NagReportLine",
		reflect.TypeOf((*NagReportLine)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-nag.NagReportSchema",
		reflect.TypeOf((*NagReportSchema)(nil)).Elem(),
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
	_jsii_.RegisterEnum(
		"cdk-nag.NagRulePostValidationStates",
		reflect.TypeOf((*NagRulePostValidationStates)(nil)).Elem(),
		map[string]interface{}{
			"UNKNOWN": NagRulePostValidationStates_UNKNOWN,
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
		"cdk-nag.PCIDSS321Checks",
		reflect.TypeOf((*PCIDSS321Checks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRule", GoMethod: "ApplyRule"},
			_jsii_.MemberMethod{JsiiMethod: "checkResource", GoMethod: "CheckResource"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleIds", GoGetter: "RuleIds"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateScope", GoMethod: "ValidateScope"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_PCIDSS321Checks{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NagPack)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.ServerlessChecks",
		reflect.TypeOf((*ServerlessChecks)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRule", GoMethod: "ApplyRule"},
			_jsii_.MemberMethod{JsiiMethod: "checkResource", GoMethod: "CheckResource"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "packName", GoGetter: "PackName"},
			_jsii_.MemberProperty{JsiiProperty: "readPackName", GoGetter: "ReadPackName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleIds", GoGetter: "RuleIds"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateScope", GoMethod: "ValidateScope"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_ServerlessChecks{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NagPack)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-nag.WriteNagSuppressionsToCloudFormationAspect",
		reflect.TypeOf((*WriteNagSuppressionsToCloudFormationAspect)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "visit", GoMethod: "Visit"},
		},
		func() interface{} {
			j := jsiiProxy_WriteNagSuppressionsToCloudFormationAspect{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIAspect)
			return &j
		},
	)
}
