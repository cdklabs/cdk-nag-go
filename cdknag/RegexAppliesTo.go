// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag


// A regular expression to apply to matching findings.
type RegexAppliesTo struct {
	// An ECMA-262 regex string.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

