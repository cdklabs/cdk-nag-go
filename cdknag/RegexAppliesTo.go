package cdknag


// A regular expression to apply to matching findings.
type RegexAppliesTo struct {
	// An ECMA-262 regex string.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}

