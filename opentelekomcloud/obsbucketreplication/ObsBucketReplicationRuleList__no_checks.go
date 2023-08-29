// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package obsbucketreplication

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObsBucketReplicationRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_ObsBucketReplicationRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ObsBucketReplicationRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObsBucketReplicationRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObsBucketReplicationRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ObsBucketReplicationRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewObsBucketReplicationRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

