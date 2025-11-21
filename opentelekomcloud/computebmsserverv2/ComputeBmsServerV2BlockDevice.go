// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computebmsserverv2


type ComputeBmsServerV2BlockDevice struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/compute_bms_server_v2#source_type ComputeBmsServerV2#source_type}.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/compute_bms_server_v2#boot_index ComputeBmsServerV2#boot_index}.
	BootIndex *float64 `field:"optional" json:"bootIndex" yaml:"bootIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/compute_bms_server_v2#delete_on_termination ComputeBmsServerV2#delete_on_termination}.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/compute_bms_server_v2#destination_type ComputeBmsServerV2#destination_type}.
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/compute_bms_server_v2#device_name ComputeBmsServerV2#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/compute_bms_server_v2#guest_format ComputeBmsServerV2#guest_format}.
	GuestFormat *string `field:"optional" json:"guestFormat" yaml:"guestFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/compute_bms_server_v2#uuid ComputeBmsServerV2#uuid}.
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/compute_bms_server_v2#volume_size ComputeBmsServerV2#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/compute_bms_server_v2#volume_type ComputeBmsServerV2#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

