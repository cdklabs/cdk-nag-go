package cdknag


// A regular expression to apply to matching findings.
// Experimental.
type RegexAppliesTo struct {
	// An ECMA-262 regex string.
	// Experimental.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

