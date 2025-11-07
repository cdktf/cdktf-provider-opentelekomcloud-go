// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ddsbackupv3

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DdsBackupV3DatastoreList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DdsBackupV3DatastoreList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DdsBackupV3DatastoreList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DdsBackupV3DatastoreList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DdsBackupV3DatastoreList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DdsBackupV3DatastoreList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDdsBackupV3DatastoreListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

