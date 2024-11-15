// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ddminstancev1

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DdmInstanceV1NodesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DdmInstanceV1NodesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DdmInstanceV1NodesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DdmInstanceV1NodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DdmInstanceV1NodesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DdmInstanceV1NodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDdmInstanceV1NodesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

