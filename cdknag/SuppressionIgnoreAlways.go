package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"
)

// Always ignore the suppression.
type SuppressionIgnoreAlways interface {
	INagSuppressionIgnore
	CreateMessage(_input *SuppressionIgnoreInput) *string
}

// The jsii proxy struct for SuppressionIgnoreAlways
type jsiiProxy_SuppressionIgnoreAlways struct {
	jsiiProxy_INagSuppressionIgnore
}

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

func NewSuppressionIgnoreAlways_Override(s SuppressionIgnoreAlways, triggerMessage *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.SuppressionIgnoreAlways",
		[]interface{}{triggerMessage},
		s,
	)
}

func (s *jsiiProxy_SuppressionIgnoreAlways) CreateMessage(_input *SuppressionIgnoreInput) *string {
	if err := s.validateCreateMessageParameters(_input); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"createMessage",
		[]interface{}{_input},
		&returns,
	)

	return returns
}

