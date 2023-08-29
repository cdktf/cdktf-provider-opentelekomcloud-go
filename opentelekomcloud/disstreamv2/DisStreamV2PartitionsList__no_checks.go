// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package disstreamv2

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DisStreamV2PartitionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DisStreamV2PartitionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DisStreamV2PartitionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DisStreamV2PartitionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DisStreamV2PartitionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDisStreamV2PartitionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

