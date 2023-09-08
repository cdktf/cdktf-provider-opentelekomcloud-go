// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lblistenerv3

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbListenerV3IpGroupList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LbListenerV3IpGroupList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LbListenerV3IpGroupList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbListenerV3IpGroupList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbListenerV3IpGroupList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LbListenerV3IpGroupList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLbListenerV3IpGroupListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

