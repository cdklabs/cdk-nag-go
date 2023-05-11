package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"
)

// Don't ignore the suppression.
type SuppressionIgnoreNever interface {
	INagSuppressionIgnore
	CreateMessage(_input *SuppressionIgnoreInput) *string
}

// The jsii proxy struct for SuppressionIgnoreNever
type jsiiProxy_SuppressionIgnoreNever struct {
	jsiiProxy_INagSuppressionIgnore
}

func NewSuppressionIgnoreNever() SuppressionIgnoreNever {
	_init_.Initialize()

	j := jsiiProxy_SuppressionIgnoreNever{}

	_jsii_.Create(
		"cdk-nag.SuppressionIgnoreNever",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewSuppressionIgnoreNever_Override(s SuppressionIgnoreNever) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.SuppressionIgnoreNever",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_SuppressionIgnoreNever) CreateMessage(_input *SuppressionIgnoreInput) *string {
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

