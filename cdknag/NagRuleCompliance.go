package cdknag


// The compliance level of a resource in relation to a rule.
type NagRuleCompliance string

const (
	NagRuleCompliance_COMPLIANT NagRuleCompliance = "COMPLIANT"
	NagRuleCompliance_NON_COMPLIANT NagRuleCompliance = "NON_COMPLIANT"
	NagRuleCompliance_NOT_APPLICABLE NagRuleCompliance = "NOT_APPLICABLE"
)

