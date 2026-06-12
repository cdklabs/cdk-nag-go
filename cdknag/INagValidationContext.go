package cdknag

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-nag-go/cdknag/v3/internal"
)

// Extended validation context that includes the construct tree.
//
// Requires CDK core change to populate `appConstruct` during plugin validation.
// Experimental.
type INagValidationContext interface {
	awscdk.IPolicyValidationContext
	// Experimental.
	AppConstruct() constructs.IConstruct
}

// The jsii proxy for INagValidationContext
type jsiiProxy_INagValidationContext struct {
	internal.Type__awscdkIPolicyValidationContext
}

func (j *jsiiProxy_INagValidationContext) AppConstruct() constructs.IConstruct {
	var returns constructs.IConstruct
	_jsii_.Get(
		j,
		"appConstruct",
		&returns,
	)
	return returns
}

