// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancev2


type ComputeInstanceV2BlockDevice struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/compute_instance_v2#source_type ComputeInstanceV2#source_type}.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/compute_instance_v2#boot_index ComputeInstanceV2#boot_index}.
	BootIndex *float64 `field:"optional" json:"bootIndex" yaml:"bootIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/compute_instance_v2#delete_on_termination ComputeInstanceV2#delete_on_termination}.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/compute_instance_v2#destination_type ComputeInstanceV2#destination_type}.
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/compute_instance_v2#device_name ComputeInstanceV2#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/compute_instance_v2#guest_format ComputeInstanceV2#guest_format}.
	GuestFormat *string `field:"optional" json:"guestFormat" yaml:"guestFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/compute_instance_v2#uuid ComputeInstanceV2#uuid}.
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/compute_instance_v2#volume_size ComputeInstanceV2#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/compute_instance_v2#volume_type ComputeInstanceV2#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

