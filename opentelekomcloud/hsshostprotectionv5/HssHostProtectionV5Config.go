// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hsshostprotectionv5

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HssHostProtectionV5Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/hss_host_protection_v5#charging_mode HssHostProtectionV5#charging_mode}.
	ChargingMode *string `field:"required" json:"chargingMode" yaml:"chargingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/hss_host_protection_v5#host_id HssHostProtectionV5#host_id}.
	HostId *string `field:"required" json:"hostId" yaml:"hostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/hss_host_protection_v5#version HssHostProtectionV5#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/hss_host_protection_v5#id HssHostProtectionV5#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/hss_host_protection_v5#is_wait_host_available HssHostProtectionV5#is_wait_host_available}.
	IsWaitHostAvailable interface{} `field:"optional" json:"isWaitHostAvailable" yaml:"isWaitHostAvailable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/hss_host_protection_v5#resource_id HssHostProtectionV5#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/hss_host_protection_v5#timeouts HssHostProtectionV5#timeouts}
	Timeouts *HssHostProtectionV5Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

