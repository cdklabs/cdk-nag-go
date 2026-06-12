//go:build !no_runtime_type_checking

package cdknag

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (w *jsiiProxy_WriteNagSuppressionsToCloudFormationAspect) validateVisitParameters(node constructs.IConstruct) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

