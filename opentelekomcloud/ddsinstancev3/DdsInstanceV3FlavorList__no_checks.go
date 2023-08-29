// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ddsinstancev3

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DdsInstanceV3FlavorList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DdsInstanceV3FlavorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DdsInstanceV3FlavorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DdsInstanceV3FlavorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DdsInstanceV3FlavorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DdsInstanceV3FlavorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDdsInstanceV3FlavorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

