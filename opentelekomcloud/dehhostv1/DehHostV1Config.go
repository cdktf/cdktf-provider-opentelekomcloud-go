// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dehhostv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DehHostV1Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#availability_zone DehHostV1#availability_zone}.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#host_type DehHostV1#host_type}.
	HostType *string `field:"required" json:"hostType" yaml:"hostType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#name DehHostV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#auto_placement DehHostV1#auto_placement}.
	AutoPlacement *string `field:"optional" json:"autoPlacement" yaml:"autoPlacement"`
	// available_instance_capacities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#available_instance_capacities DehHostV1#available_instance_capacities}
	AvailableInstanceCapacities interface{} `field:"optional" json:"availableInstanceCapacities" yaml:"availableInstanceCapacities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#available_memory DehHostV1#available_memory}.
	AvailableMemory *float64 `field:"optional" json:"availableMemory" yaml:"availableMemory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#available_vcpus DehHostV1#available_vcpus}.
	AvailableVcpus *float64 `field:"optional" json:"availableVcpus" yaml:"availableVcpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#cores DehHostV1#cores}.
	Cores *float64 `field:"optional" json:"cores" yaml:"cores"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#host_type_name DehHostV1#host_type_name}.
	HostTypeName *string `field:"optional" json:"hostTypeName" yaml:"hostTypeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#id DehHostV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#instance_total DehHostV1#instance_total}.
	InstanceTotal *float64 `field:"optional" json:"instanceTotal" yaml:"instanceTotal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#instance_uuids DehHostV1#instance_uuids}.
	InstanceUuids *[]*string `field:"optional" json:"instanceUuids" yaml:"instanceUuids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#memory DehHostV1#memory}.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#region DehHostV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#sockets DehHostV1#sockets}.
	Sockets *float64 `field:"optional" json:"sockets" yaml:"sockets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#status DehHostV1#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#tags DehHostV1#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#timeouts DehHostV1#timeouts}
	Timeouts *DehHostV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/deh_host_v1#vcpus DehHostV1#vcpus}.
	Vcpus *float64 `field:"optional" json:"vcpus" yaml:"vcpus"`
}

