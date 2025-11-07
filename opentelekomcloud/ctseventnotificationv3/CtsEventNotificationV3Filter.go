// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ctseventnotificationv3


type CtsEventNotificationV3Filter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cts_event_notification_v3#condition CtsEventNotificationV3#condition}.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cts_event_notification_v3#rule CtsEventNotificationV3#rule}.
	Rule *[]*string `field:"required" json:"rule" yaml:"rule"`
}

