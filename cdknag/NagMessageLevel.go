package cdknag


// The severity level of the rule.
type NagMessageLevel string

const (
	NagMessageLevel_WARN NagMessageLevel = "WARN"
	NagMessageLevel_ERROR NagMessageLevel = "ERROR"
	NagMessageLevel_INFO NagMessageLevel = "INFO"
)

