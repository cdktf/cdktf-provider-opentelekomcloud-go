//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cesalarmrule

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CesAlarmruleMetricDimensionsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CesAlarmruleMetricDimensionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CesAlarmruleMetricDimensionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CesAlarmruleMetricDimensionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CesAlarmruleMetricDimensionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CesAlarmruleMetricDimensionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCesAlarmruleMetricDimensionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
