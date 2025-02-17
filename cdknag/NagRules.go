package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Helper class with methods for rule creation.
type NagRules interface {
}

// The jsii proxy struct for NagRules
type jsiiProxy_NagRules struct {
	_ byte // padding
}

func NewNagRules() NagRules {
	_init_.Initialize()

	j := jsiiProxy_NagRules{}

	_jsii_.Create(
		"cdk-nag.NagRules",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewNagRules_Override(n NagRules) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.NagRules",
		nil, // no parameters
		n,
	)
}

// Use in cases where a primitive value must be known to pass a rule.
//
// https://developer.mozilla.org/en-US/docs/Glossary/Primitive
//
// Returns: Return a value if resolves to a primitive data type, otherwise throw an error.
func NagRules_ResolveIfPrimitive(node awscdk.CfnResource, parameter interface{}) interface{} {
	_init_.Initialize()

	if err := validateNagRules_ResolveIfPrimitiveParameters(node, parameter); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdk-nag.NagRules",
		"resolveIfPrimitive",
		[]interface{}{node, parameter},
		&returns,
	)

	return returns
}

// Returns: Return the Logical resource Id if resolves to a intrinsic function, otherwise the resolved provided value.
// Deprecated: Use resolveResourceFromIntrinsic instead
//
// Use in cases where a token resolves to an intrinsic function and the referenced resource must be known to pass a rule.
func NagRules_ResolveResourceFromInstrinsic(node awscdk.CfnResource, parameter interface{}) interface{} {
	_init_.Initialize()

	if err := validateNagRules_ResolveResourceFromInstrinsicParameters(node, parameter); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdk-nag.NagRules",
		"resolveResourceFromInstrinsic",
		[]interface{}{node, parameter},
		&returns,
	)

	return returns
}

// Use in cases where a token resolves to an intrinsic function and the referenced resource must be known to pass a rule.
//
// Returns: Return the Logical resource Id if resolves to a intrinsic function, otherwise the resolved provided value.
func NagRules_ResolveResourceFromIntrinsic(node awscdk.CfnResource, parameter interface{}) interface{} {
	_init_.Initialize()

	if err := validateNagRules_ResolveResourceFromIntrinsicParameters(node, parameter); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdk-nag.NagRules",
		"resolveResourceFromIntrinsic",
		[]interface{}{node, parameter},
		&returns,
	)

	return returns
}

