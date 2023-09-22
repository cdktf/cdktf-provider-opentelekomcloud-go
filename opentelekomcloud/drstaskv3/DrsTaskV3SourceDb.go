// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package drstaskv3


type DrsTaskV3SourceDb struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#engine_type DrsTaskV3#engine_type}.
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#ip DrsTaskV3#ip}.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#password DrsTaskV3#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#port DrsTaskV3#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#user DrsTaskV3#user}.
	User *string `field:"required" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#instance_id DrsTaskV3#instance_id}.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#name DrsTaskV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#region DrsTaskV3#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#ssl_cert_check_sum DrsTaskV3#ssl_cert_check_sum}.
	SslCertCheckSum *string `field:"optional" json:"sslCertCheckSum" yaml:"sslCertCheckSum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#ssl_cert_key DrsTaskV3#ssl_cert_key}.
	SslCertKey *string `field:"optional" json:"sslCertKey" yaml:"sslCertKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#ssl_cert_name DrsTaskV3#ssl_cert_name}.
	SslCertName *string `field:"optional" json:"sslCertName" yaml:"sslCertName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#ssl_cert_password DrsTaskV3#ssl_cert_password}.
	SslCertPassword *string `field:"optional" json:"sslCertPassword" yaml:"sslCertPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#ssl_enabled DrsTaskV3#ssl_enabled}.
	SslEnabled interface{} `field:"optional" json:"sslEnabled" yaml:"sslEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#subnet_id DrsTaskV3#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

