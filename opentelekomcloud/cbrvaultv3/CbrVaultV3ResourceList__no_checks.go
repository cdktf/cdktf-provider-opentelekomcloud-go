// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cbrvaultv3

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CbrVaultV3ResourceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CbrVaultV3ResourceList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CbrVaultV3ResourceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CbrVaultV3ResourceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CbrVaultV3ResourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CbrVaultV3ResourceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CbrVaultV3ResourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCbrVaultV3ResourceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

