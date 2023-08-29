// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkingportv2

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkingPortV2FixedIpList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkingPortV2FixedIpList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkingPortV2FixedIpList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkingPortV2FixedIpList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkingPortV2FixedIpList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkingPortV2FixedIpList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkingPortV2FixedIpListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

