// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3bucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionExpirationList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionExpirationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionExpirationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionExpirationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionExpirationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleNoncurrentVersionExpirationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketLifecycleRuleNoncurrentVersionExpirationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

