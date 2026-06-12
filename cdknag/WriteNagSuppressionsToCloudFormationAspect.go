package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-nag-go/cdknag/v3/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-nag-go/cdknag/v3/internal"
)

// An IAspect that reads acknowledged rules from construct metadata and writes them into the CfnResource's CloudFormation Metadata for audit trail persistence in the synthesized template.
//
// Preserves the v2 `cdk_nag`
// metadata format.
// Experimental.
type WriteNagSuppressionsToCloudFormationAspect interface {
	awscdk.IAspect
	// All aspects can visit an IConstruct.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for WriteNagSuppressionsToCloudFormationAspect
type jsiiProxy_WriteNagSuppressionsToCloudFormationAspect struct {
	internal.Type__awscdkIAspect
}

// Experimental.
func NewWriteNagSuppressionsToCloudFormationAspect() WriteNagSuppressionsToCloudFormationAspect {
	_init_.Initialize()

	j := jsiiProxy_WriteNagSuppressionsToCloudFormationAspect{}

	_jsii_.Create(
		"cdk-nag.WriteNagSuppressionsToCloudFormationAspect",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWriteNagSuppressionsToCloudFormationAspect_Override(w WriteNagSuppressionsToCloudFormationAspect) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-nag.WriteNagSuppressionsToCloudFormationAspect",
		nil, // no parameters
		w,
	)
}

func (w *jsiiProxy_WriteNagSuppressionsToCloudFormationAspect) Visit(node constructs.IConstruct) {
	if err := w.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"visit",
		[]interface{}{node},
	)
}

