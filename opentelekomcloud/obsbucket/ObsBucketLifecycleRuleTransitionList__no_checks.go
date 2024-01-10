// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package obsbucket

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObsBucketLifecycleRuleTransitionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_ObsBucketLifecycleRuleTransitionList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_ObsBucketLifecycleRuleTransitionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLifecycleRuleTransitionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLifecycleRuleTransitionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLifecycleRuleTransitionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLifecycleRuleTransitionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewObsBucketLifecycleRuleTransitionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

