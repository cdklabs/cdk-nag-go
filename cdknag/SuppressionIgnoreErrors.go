// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"
)

// Ignore Suppressions for Rules with a NagMessageLevel.ERROR.
type SuppressionIgnoreErrors interface {
	INagSuppressionIgnore
	CreateMessage(input *SuppressionIgnoreInput) *string
}

// The jsii proxy struct for SuppressionIgnoreErrors
type jsiiProxy_SuppressionIgnoreErrors struct {
	jsiiProxy_INagSuppressionIgnore
}

func NewSuppressionIgnoreErrors() SuppressionIgnoreErrors {
	_init_.Initialize()

	j := jsiiProxy_SuppressionIgnoreErrors{}

	_jsii_.Create(
		"cdk-nag.SuppressionIgnoreErrors",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewSuppressionIgnoreErrors_Override(s SuppressionIgnoreErrors) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.SuppressionIgnoreErrors",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_SuppressionIgnoreErrors) CreateMessage(input *SuppressionIgnoreInput) *string {
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

