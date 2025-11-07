// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package obsbucketobjectacl

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObsBucketObjectAclOwnerPermissionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_ObsBucketObjectAclOwnerPermissionList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_ObsBucketObjectAclOwnerPermissionList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ObsBucketObjectAclOwnerPermissionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObsBucketObjectAclOwnerPermissionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ObsBucketObjectAclOwnerPermissionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewObsBucketObjectAclOwnerPermissionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

