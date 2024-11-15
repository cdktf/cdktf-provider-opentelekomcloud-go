// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ddmschemav1

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DdmSchemaV1DatabasesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DdmSchemaV1DatabasesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DdmSchemaV1DatabasesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DdmSchemaV1DatabasesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DdmSchemaV1DatabasesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DdmSchemaV1DatabasesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDdmSchemaV1DatabasesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

