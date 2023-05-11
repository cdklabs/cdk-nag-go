package cdknag


// Additional states a rule can be in post compliance validation.
type NagRulePostValidationStates string

const (
	NagRulePostValidationStates_SUPPRESSED NagRulePostValidationStates = "SUPPRESSED"
	NagRulePostValidationStates_UNKNOWN NagRulePostValidationStates = "UNKNOWN"
)

