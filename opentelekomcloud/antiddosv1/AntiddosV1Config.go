// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package antiddosv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AntiddosV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/antiddos_v1#app_type_id AntiddosV1#app_type_id}.
	AppTypeId *float64 `field:"required" json:"appTypeId" yaml:"appTypeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/antiddos_v1#cleaning_access_pos_id AntiddosV1#cleaning_access_pos_id}.
	CleaningAccessPosId *float64 `field:"required" json:"cleaningAccessPosId" yaml:"cleaningAccessPosId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/antiddos_v1#enable_l7 AntiddosV1#enable_l7}.
	EnableL7 interface{} `field:"required" json:"enableL7" yaml:"enableL7"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/antiddos_v1#floating_ip_id AntiddosV1#floating_ip_id}.
	FloatingIpId *string `field:"required" json:"floatingIpId" yaml:"floatingIpId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/antiddos_v1#http_request_pos_id AntiddosV1#http_request_pos_id}.
	HttpRequestPosId *float64 `field:"required" json:"httpRequestPosId" yaml:"httpRequestPosId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/antiddos_v1#traffic_pos_id AntiddosV1#traffic_pos_id}.
	TrafficPosId *float64 `field:"required" json:"trafficPosId" yaml:"trafficPosId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/antiddos_v1#id AntiddosV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/antiddos_v1#region AntiddosV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/antiddos_v1#timeouts AntiddosV1#timeouts}
	Timeouts *AntiddosV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

