// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asconfigurationv1


type AsConfigurationV1InstanceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_configuration_v1#key_name AsConfigurationV1#key_name}.
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_configuration_v1#disk AsConfigurationV1#disk}
	Disk interface{} `field:"optional" json:"disk" yaml:"disk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_configuration_v1#flavor AsConfigurationV1#flavor}.
	Flavor *string `field:"optional" json:"flavor" yaml:"flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_configuration_v1#image AsConfigurationV1#image}.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_configuration_v1#instance_id AsConfigurationV1#instance_id}.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_configuration_v1#metadata AsConfigurationV1#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// personality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_configuration_v1#personality AsConfigurationV1#personality}
	Personality interface{} `field:"optional" json:"personality" yaml:"personality"`
	// public_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_configuration_v1#public_ip AsConfigurationV1#public_ip}
	PublicIp *AsConfigurationV1InstanceConfigPublicIp `field:"optional" json:"publicIp" yaml:"publicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_configuration_v1#security_groups AsConfigurationV1#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/as_configuration_v1#user_data AsConfigurationV1#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

