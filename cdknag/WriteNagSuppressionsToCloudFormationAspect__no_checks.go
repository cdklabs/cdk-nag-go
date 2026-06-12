//go:build no_runtime_type_checking

package cdknag

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WriteNagSuppressionsToCloudFormationAspect) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

