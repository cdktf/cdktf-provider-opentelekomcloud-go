// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lbpoolv2

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbPoolV2PersistenceList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LbPoolV2PersistenceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LbPoolV2PersistenceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbPoolV2PersistenceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbPoolV2PersistenceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LbPoolV2PersistenceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLbPoolV2PersistenceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

