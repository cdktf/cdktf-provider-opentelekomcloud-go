// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package tmstagsv1

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TmsTagsV1TagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TmsTagsV1TagsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TmsTagsV1TagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TmsTagsV1TagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TmsTagsV1TagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TmsTagsV1TagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TmsTagsV1TagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTmsTagsV1TagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

