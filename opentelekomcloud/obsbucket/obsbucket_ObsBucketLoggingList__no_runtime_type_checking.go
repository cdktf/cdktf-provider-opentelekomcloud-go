//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package obsbucket

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObsBucketLoggingList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_ObsBucketLoggingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLoggingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLoggingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLoggingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ObsBucketLoggingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewObsBucketLoggingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

