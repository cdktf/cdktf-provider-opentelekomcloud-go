// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ddmschemav1


type DdmSchemaV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/ddm_schema_v1#create DdmSchemaV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/ddm_schema_v1#delete DdmSchemaV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/ddm_schema_v1#update DdmSchemaV1#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

