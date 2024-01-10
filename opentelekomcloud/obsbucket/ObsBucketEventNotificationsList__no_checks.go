// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package obsbucket

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObsBucketEventNotificationsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_ObsBucketEventNotificationsList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_ObsBucketEventNotificationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ObsBucketEventNotificationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObsBucketEventNotificationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObsBucketEventNotificationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ObsBucketEventNotificationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewObsBucketEventNotificationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

