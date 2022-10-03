//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package s3bucket

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketLifecycleRuleExpirationList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3BucketLifecycleRuleExpirationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleExpirationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleExpirationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleExpirationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3BucketLifecycleRuleExpirationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3BucketLifecycleRuleExpirationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

