// Check CDK v2 applications for best practices using a combination on available rule packs.
package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Helper class with methods to add cdk-nag suppressions to cdk resources.
type NagSuppressions interface {
}

// The jsii proxy struct for NagSuppressions
type jsiiProxy_NagSuppressions struct {
	_ byte // padding
}

func NewNagSuppressions() NagSuppressions {
	_init_.Initialize()

	j := jsiiProxy_NagSuppressions{}

	_jsii_.Create(
		"cdk-nag.NagSuppressions",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewNagSuppressions_Override(n NagSuppressions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.NagSuppressions",
		nil, // no parameters
		n,
	)
}

// Add cdk-nag suppressions to a CfnResource and optionally its children.
func NagSuppressions_AddResourceSuppressions(construct interface{}, suppressions *[]*NagPackSuppression, applyToChildren *bool) {
	_init_.Initialize()

	if err := validateNagSuppressions_AddResourceSuppressionsParameters(construct, suppressions); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"cdk-nag.NagSuppressions",
		"addResourceSuppressions",
		[]interface{}{construct, suppressions, applyToChildren},
	)
}

// Add cdk-nag suppressions to a CfnResource and optionally its children via its path.
func NagSuppressions_AddResourceSuppressionsByPath(stack awscdk.Stack, path interface{}, suppressions *[]*NagPackSuppression, applyToChildren *bool) {
	_init_.Initialize()

	if err := validateNagSuppressions_AddResourceSuppressionsByPathParameters(stack, path, suppressions); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"cdk-nag.NagSuppressions",
		"addResourceSuppressionsByPath",
		[]interface{}{stack, path, suppressions, applyToChildren},
	)
}

// Apply cdk-nag suppressions to a Stack and optionally nested stacks.
func NagSuppressions_AddStackSuppressions(stack awscdk.Stack, suppressions *[]*NagPackSuppression, applyToNestedStacks *bool) {
	_init_.Initialize()

	if err := validateNagSuppressions_AddStackSuppressionsParameters(stack, suppressions); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"cdk-nag.NagSuppressions",
		"addStackSuppressions",
		[]interface{}{stack, suppressions, applyToNestedStacks},
	)
}

