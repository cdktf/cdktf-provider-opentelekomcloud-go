//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ddsinstancev3

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DdsInstanceV3NodesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DdsInstanceV3NodesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DdsInstanceV3NodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DdsInstanceV3NodesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DdsInstanceV3NodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDdsInstanceV3NodesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

