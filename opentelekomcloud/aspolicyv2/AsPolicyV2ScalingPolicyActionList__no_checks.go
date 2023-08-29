// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package aspolicyv2

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AsPolicyV2ScalingPolicyActionList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AsPolicyV2ScalingPolicyActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AsPolicyV2ScalingPolicyActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AsPolicyV2ScalingPolicyActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AsPolicyV2ScalingPolicyActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AsPolicyV2ScalingPolicyActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAsPolicyV2ScalingPolicyActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

