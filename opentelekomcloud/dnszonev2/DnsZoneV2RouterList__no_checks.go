// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dnszonev2

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsZoneV2RouterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DnsZoneV2RouterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DnsZoneV2RouterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsZoneV2RouterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsZoneV2RouterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DnsZoneV2RouterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDnsZoneV2RouterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

