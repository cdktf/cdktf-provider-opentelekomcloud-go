//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsInstanceV3NodesList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsInstanceV3NodesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsInstanceV3NodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsInstanceV3NodesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsInstanceV3NodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsInstanceV3NodesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
