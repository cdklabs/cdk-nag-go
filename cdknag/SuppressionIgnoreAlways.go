package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"
)

// Always ignore the suppression.
// Experimental.
type SuppressionIgnoreAlways interface {
	INagSuppressionIgnore
	// Experimental.
	CreateMessage(input *SuppressionIgnoreInput) *string
}

// The jsii proxy struct for SuppressionIgnoreAlways
type jsiiProxy_SuppressionIgnoreAlways struct {
	jsiiProxy_INagSuppressionIgnore
}

// Experimental.
func NewSuppressionIgnoreAlways(triggerMessage *string) SuppressionIgnoreAlways {
	_init_.Initialize()

	if err := validateNewSuppressionIgnoreAlwaysParameters(triggerMessage); err != nil {
		panic(err)
	}
	j := jsiiProxy_SuppressionIgnoreAlways{}

	_jsii_.Create(
		"cdk-nag.SuppressionIgnoreAlways",
		[]interface{}{triggerMessage},
		&j,
	)

	return &j
}

// Experimental.
func NewSuppressionIgnoreAlways_Override(s SuppressionIgnoreAlways, triggerMessage *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.SuppressionIgnoreAlways",
		[]interface{}{triggerMessage},
		s,
	)
}

func (s *jsiiProxy_SuppressionIgnoreAlways) CreateMessage(input *SuppressionIgnoreInput) *string {
	if err := s.validateCreateMessageParameters(input); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"createMessage",
		[]interface{}{input},
		&returns,
	)

	return returns
}

