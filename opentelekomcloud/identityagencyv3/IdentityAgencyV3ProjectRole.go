// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityagencyv3


type IdentityAgencyV3ProjectRole struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/identity_agency_v3#project IdentityAgencyV3#project}.
	Project *string `field:"required" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/identity_agency_v3#roles IdentityAgencyV3#roles}.
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
}

