// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lbrulev3

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbRuleV3ConditionsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LbRuleV3ConditionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LbRuleV3ConditionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbRuleV3ConditionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbRuleV3ConditionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LbRuleV3ConditionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLbRuleV3ConditionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

