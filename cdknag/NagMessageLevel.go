// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag


// The severity level of the rule.
type NagMessageLevel string

const (
	NagMessageLevel_WARN NagMessageLevel = "WARN"
	NagMessageLevel_ERROR NagMessageLevel = "ERROR"
)

