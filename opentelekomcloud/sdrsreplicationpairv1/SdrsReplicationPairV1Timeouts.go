// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sdrsreplicationpairv1


type SdrsReplicationPairV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/sdrs_replication_pair_v1#create SdrsReplicationPairV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/sdrs_replication_pair_v1#delete SdrsReplicationPairV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

