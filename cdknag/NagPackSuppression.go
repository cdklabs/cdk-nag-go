package cdknag


// Interface for creating a rule suppression.
type NagPackSuppression struct {
	// The id of the rule to ignore.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The reason to ignore the rule (minimum 10 characters).
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// Rule specific granular suppressions.
	AppliesTo *[]interface{} `field:"optional" json:"appliesTo" yaml:"appliesTo"`
}

