//go:build no_runtime_type_checking

package identityprovider

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IdentityProviderConversionRulesLocalList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IdentityProviderConversionRulesLocalList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IdentityProviderConversionRulesLocalList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IdentityProviderConversionRulesLocalList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IdentityProviderConversionRulesLocalList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIdentityProviderConversionRulesLocalListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

