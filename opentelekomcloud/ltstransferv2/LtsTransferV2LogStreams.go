// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltstransferv2


type LtsTransferV2LogStreams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/lts_transfer_v2#log_stream_id LtsTransferV2#log_stream_id}.
	LogStreamId *string `field:"required" json:"logStreamId" yaml:"logStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/lts_transfer_v2#log_stream_name LtsTransferV2#log_stream_name}.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
}

