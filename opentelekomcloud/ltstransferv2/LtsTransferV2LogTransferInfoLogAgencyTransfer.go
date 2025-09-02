// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltstransferv2


type LtsTransferV2LogTransferInfoLogAgencyTransfer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/lts_transfer_v2#agency_domain_id LtsTransferV2#agency_domain_id}.
	AgencyDomainId *string `field:"required" json:"agencyDomainId" yaml:"agencyDomainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/lts_transfer_v2#agency_domain_name LtsTransferV2#agency_domain_name}.
	AgencyDomainName *string `field:"required" json:"agencyDomainName" yaml:"agencyDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/lts_transfer_v2#agency_name LtsTransferV2#agency_name}.
	AgencyName *string `field:"required" json:"agencyName" yaml:"agencyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/lts_transfer_v2#agency_project_id LtsTransferV2#agency_project_id}.
	AgencyProjectId *string `field:"required" json:"agencyProjectId" yaml:"agencyProjectId"`
}

