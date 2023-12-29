// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcsinstancev1


type DcsInstanceV1Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/dcs_instance_v1#parameter_id DcsInstanceV1#parameter_id}.
	ParameterId *string `field:"required" json:"parameterId" yaml:"parameterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/dcs_instance_v1#parameter_name DcsInstanceV1#parameter_name}.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/dcs_instance_v1#parameter_value DcsInstanceV1#parameter_value}.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

