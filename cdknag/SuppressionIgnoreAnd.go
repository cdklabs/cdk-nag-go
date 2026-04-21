package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"
)

// Ignore the suppression if all of the given INagSuppressionIgnore return a non-empty message.
// Experimental.
type SuppressionIgnoreAnd interface {
	INagSuppressionIgnore
	// Experimental.
	CreateMessage(input *SuppressionIgnoreInput) *string
}

// The jsii proxy struct for SuppressionIgnoreAnd
type jsiiProxy_SuppressionIgnoreAnd struct {
	jsiiProxy_INagSuppressionIgnore
}

// Experimental.
func NewSuppressionIgnoreAnd(SuppressionIgnoreAnds ...INagSuppressionIgnore) SuppressionIgnoreAnd {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range SuppressionIgnoreAnds {
		args = append(args, a)
	}

	j := jsiiProxy_SuppressionIgnoreAnd{}

	_jsii_.Create(
		"cdk-nag.SuppressionIgnoreAnd",
		args,
		&j,
	)

	return &j
}

// Experimental.
func NewSuppressionIgnoreAnd_Override(s SuppressionIgnoreAnd, SuppressionIgnoreAnds ...INagSuppressionIgnore) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range SuppressionIgnoreAnds {
		args = append(args, a)
	}

	_jsii_.Create(
		"cdk-nag.SuppressionIgnoreAnd",
		args,
		s,
	)
}

func (s *jsiiProxy_SuppressionIgnoreAnd) CreateMessage(input *SuppressionIgnoreInput) *string {
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

