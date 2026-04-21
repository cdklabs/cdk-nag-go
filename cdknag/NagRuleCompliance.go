package cdknag


// The compliance level of a resource in relation to a rule.
// Experimental.
type NagRuleCompliance string

const (
	// Experimental.
	NagRuleCompliance_COMPLIANT NagRuleCompliance = "COMPLIANT"
	// Experimental.
	NagRuleCompliance_NON_COMPLIANT NagRuleCompliance = "NON_COMPLIANT"
	// Experimental.
	NagRuleCompliance_NOT_APPLICABLE NagRuleCompliance = "NOT_APPLICABLE"
)

