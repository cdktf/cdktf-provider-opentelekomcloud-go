// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltstransferv2


type LtsTransferV2LogTransferInfoLogTransferDetail struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#obs_bucket_name LtsTransferV2#obs_bucket_name}.
	ObsBucketName *string `field:"optional" json:"obsBucketName" yaml:"obsBucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#obs_dir_prefix_name LtsTransferV2#obs_dir_prefix_name}.
	ObsDirPrefixName *string `field:"optional" json:"obsDirPrefixName" yaml:"obsDirPrefixName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#obs_encrypted_enable LtsTransferV2#obs_encrypted_enable}.
	ObsEncryptedEnable interface{} `field:"optional" json:"obsEncryptedEnable" yaml:"obsEncryptedEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#obs_encrypted_id LtsTransferV2#obs_encrypted_id}.
	ObsEncryptedId *string `field:"optional" json:"obsEncryptedId" yaml:"obsEncryptedId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#obs_eps_id LtsTransferV2#obs_eps_id}.
	ObsEpsId *string `field:"optional" json:"obsEpsId" yaml:"obsEpsId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#obs_period LtsTransferV2#obs_period}.
	ObsPeriod *float64 `field:"optional" json:"obsPeriod" yaml:"obsPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#obs_period_unit LtsTransferV2#obs_period_unit}.
	ObsPeriodUnit *string `field:"optional" json:"obsPeriodUnit" yaml:"obsPeriodUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#obs_prefix_name LtsTransferV2#obs_prefix_name}.
	ObsPrefixName *string `field:"optional" json:"obsPrefixName" yaml:"obsPrefixName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#obs_time_zone LtsTransferV2#obs_time_zone}.
	ObsTimeZone *string `field:"optional" json:"obsTimeZone" yaml:"obsTimeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#obs_time_zone_id LtsTransferV2#obs_time_zone_id}.
	ObsTimeZoneId *string `field:"optional" json:"obsTimeZoneId" yaml:"obsTimeZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#obs_transfer_path LtsTransferV2#obs_transfer_path}.
	ObsTransferPath *string `field:"optional" json:"obsTransferPath" yaml:"obsTransferPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/lts_transfer_v2#tags LtsTransferV2#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

