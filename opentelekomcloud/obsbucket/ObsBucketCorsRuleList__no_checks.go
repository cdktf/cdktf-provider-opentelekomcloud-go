// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package obsbucket

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObsBucketCorsRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_ObsBucketCorsRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_ObsBucketCorsRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ObsBucketCorsRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObsBucketCorsRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObsBucketCorsRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ObsBucketCorsRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewObsBucketCorsRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

