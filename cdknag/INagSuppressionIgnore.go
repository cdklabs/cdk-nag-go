package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for creating NagSuppression Ignores.
type INagSuppressionIgnore interface {
	CreateMessage(input *SuppressionIgnoreInput) *string
}

// The jsii proxy for INagSuppressionIgnore
type jsiiProxy_INagSuppressionIgnore struct {
	_ byte // padding
}

func (i *jsiiProxy_INagSuppressionIgnore) CreateMessage(input *SuppressionIgnoreInput) *string {
	if err := i.validateCreateMessageParameters(input); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"createMessage",
		[]interface{}{input},
		&returns,
	)

	return returns
}

