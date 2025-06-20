// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcepservicev1


type VpcepServiceV1Port struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/vpcep_service_v1#client_port VpcepServiceV1#client_port}.
	ClientPort *float64 `field:"required" json:"clientPort" yaml:"clientPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/vpcep_service_v1#server_port VpcepServiceV1#server_port}.
	ServerPort *float64 `field:"required" json:"serverPort" yaml:"serverPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/vpcep_service_v1#protocol VpcepServiceV1#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

