// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package identityprovider

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IdentityProviderConversionRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IdentityProviderConversionRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IdentityProviderConversionRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IdentityProviderConversionRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IdentityProviderConversionRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIdentityProviderConversionRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

