// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cbrvaultv3


type CbrVaultV3Policy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cbr_vault_v3#id CbrVaultV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cbr_vault_v3#destination_vault_id CbrVaultV3#destination_vault_id}.
	DestinationVaultId *string `field:"optional" json:"destinationVaultId" yaml:"destinationVaultId"`
}

