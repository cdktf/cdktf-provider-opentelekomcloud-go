// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package identityprovider

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IdentityProviderConversionRulesRemoteList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IdentityProviderConversionRulesRemoteList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IdentityProviderConversionRulesRemoteList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IdentityProviderConversionRulesRemoteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IdentityProviderConversionRulesRemoteList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IdentityProviderConversionRulesRemoteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIdentityProviderConversionRulesRemoteListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

