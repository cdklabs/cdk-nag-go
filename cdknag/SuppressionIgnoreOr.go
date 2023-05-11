package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"
)

// Ignore the suppression if any of the given INagSuppressionIgnore return a non-empty message.
type SuppressionIgnoreOr interface {
	INagSuppressionIgnore
	CreateMessage(input *SuppressionIgnoreInput) *string
}

// The jsii proxy struct for SuppressionIgnoreOr
type jsiiProxy_SuppressionIgnoreOr struct {
	jsiiProxy_INagSuppressionIgnore
}

func NewSuppressionIgnoreOr(orSuppressionIgnores ...INagSuppressionIgnore) SuppressionIgnoreOr {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range orSuppressionIgnores {
		args = append(args, a)
	}

	j := jsiiProxy_SuppressionIgnoreOr{}

	_jsii_.Create(
		"cdk-nag.SuppressionIgnoreOr",
		args,
		&j,
	)

	return &j
}

func NewSuppressionIgnoreOr_Override(s SuppressionIgnoreOr, orSuppressionIgnores ...INagSuppressionIgnore) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range orSuppressionIgnores {
		args = append(args, a)
	}

	_jsii_.Create(
		"cdk-nag.SuppressionIgnoreOr",
		args,
		s,
	)
}

func (s *jsiiProxy_SuppressionIgnoreOr) CreateMessage(input *SuppressionIgnoreInput) *string {
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

