// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cesalarmrule


type CesAlarmruleOkActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/ces_alarmrule#notification_list CesAlarmrule#notification_list}.
	NotificationList *[]*string `field:"required" json:"notificationList" yaml:"notificationList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/ces_alarmrule#type CesAlarmrule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

