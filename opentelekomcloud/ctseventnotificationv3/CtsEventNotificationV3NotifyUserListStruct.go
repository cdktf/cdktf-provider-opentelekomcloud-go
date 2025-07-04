// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ctseventnotificationv3


type CtsEventNotificationV3NotifyUserListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/cts_event_notification_v3#user_group CtsEventNotificationV3#user_group}.
	UserGroup *string `field:"required" json:"userGroup" yaml:"userGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/cts_event_notification_v3#user_list CtsEventNotificationV3#user_list}.
	UserList *[]*string `field:"required" json:"userList" yaml:"userList"`
}

