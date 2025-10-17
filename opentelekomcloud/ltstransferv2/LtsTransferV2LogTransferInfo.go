// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltstransferv2


type LtsTransferV2LogTransferInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_transfer_v2#log_storage_format LtsTransferV2#log_storage_format}.
	LogStorageFormat *string `field:"required" json:"logStorageFormat" yaml:"logStorageFormat"`
	// log_transfer_detail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_transfer_v2#log_transfer_detail LtsTransferV2#log_transfer_detail}
	LogTransferDetail *LtsTransferV2LogTransferInfoLogTransferDetail `field:"required" json:"logTransferDetail" yaml:"logTransferDetail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_transfer_v2#log_transfer_mode LtsTransferV2#log_transfer_mode}.
	LogTransferMode *string `field:"required" json:"logTransferMode" yaml:"logTransferMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_transfer_v2#log_transfer_status LtsTransferV2#log_transfer_status}.
	LogTransferStatus *string `field:"required" json:"logTransferStatus" yaml:"logTransferStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_transfer_v2#log_transfer_type LtsTransferV2#log_transfer_type}.
	LogTransferType *string `field:"required" json:"logTransferType" yaml:"logTransferType"`
	// log_agency_transfer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_transfer_v2#log_agency_transfer LtsTransferV2#log_agency_transfer}
	LogAgencyTransfer *LtsTransferV2LogTransferInfoLogAgencyTransfer `field:"optional" json:"logAgencyTransfer" yaml:"logAgencyTransfer"`
}

