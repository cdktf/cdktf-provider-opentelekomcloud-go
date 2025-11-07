// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package obsbucket

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObsBucketLifecycleRuleTagList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_ObsBucketLifecycleRuleTagList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_ObsBucketLifecycleRuleTagList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLifecycleRuleTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLifecycleRuleTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLifecycleRuleTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLifecycleRuleTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewObsBucketLifecycleRuleTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

