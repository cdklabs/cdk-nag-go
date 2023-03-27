// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag


// Additional states a rule can be in post compliance validation.
type NagRulePostValidationStates string

const (
	NagRulePostValidationStates_SUPPRESSED NagRulePostValidationStates = "SUPPRESSED"
	NagRulePostValidationStates_UNKNOWN NagRulePostValidationStates = "UNKNOWN"
)

